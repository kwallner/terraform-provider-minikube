---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "minikube_cluster Resource - terraform-provider-minikube"
subcategory: ""
description: |-
  Used to create a minikube cluster on the current host
---

# minikube_cluster (Resource)

Used to create a minikube cluster on the current host

## Example Usage

```terraform
provider "minikube" {
  kubernetes_version = "v1.26.1"
}

resource "minikube_cluster" "docker" {
  driver       = "docker"
  cluster_name = "terraform-provider-minikube-acc-docker"
  addons = [
    "default-storageclass",
  ]
}

resource "minikube_cluster" "hyperkit" {
  vm           = true
  driver       = "hyperkit"
  cluster_name = "terraform-provider-minikube-acc-hyperkit"
  nodes        = 3
  addons = [
    "dashboard",
    "default-storageclass",
    "ingress"
  ]
}

provider "kubernetes" {
  host = minikube_cluster.docker.host

  client_certificate     = minikube_cluster.docker.client_certificate
  client_key             = minikube_cluster.docker.client_key
  cluster_ca_certificate = minikube_cluster.docker.cluster_ca_certificate
}


resource "kubernetes_deployment" "deployment" {
  metadata {
    name = "nginx-example"
    labels = {
      App = "NginxExample"
    }
  }

  spec {
    replicas = 2
    selector {
      match_labels = {
        App = "NginxExample"
      }
    }
    template {
      metadata {
        labels = {
          App = "NginxExample"
        }
      }
      spec {
        container {
          image = "nginx:1.7.8"
          name  = "example"

          port {
            container_port = 80
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `addons` (List of String) Enable addons. see `minikube addons list` for a list of valid addon names.
- `apiserver_ips` (List of String) A set of apiserver IP Addresses which are used in the generated certificate for kubernetes.  This can be used if you want to make the apiserver available from outside the machine
- `apiserver_name` (String) The authoritative apiserver hostname for apiserver certificates and connectivity. This can be used if you want to make the apiserver available from outside the machine
- `apiserver_names` (List of String) A set of apiserver names which are used in the generated certificate for kubernetes.  This can be used if you want to make the apiserver available from outside the machine
- `apiserver_port` (Number) The apiserver listening port
- `auto_update_drivers` (Boolean) If set, automatically updates drivers to the latest version. Defaults to true.
- `base_image` (String) The base image to use for docker/podman drivers. Intended for local development.
- `binary_mirror` (String) Location to fetch kubectl, kubelet, & kubeadm binaries from.
- `cache_images` (Boolean) If true, cache docker images for the current bootstrapper and load them into the machine. Always false with --driver=none.
- `cert_expiration` (Number) Duration until minikube certificate expiration, defaults to three years (26280h). (Configured in minutes)
- `cluster_name` (String) The name of the minikube cluster
- `cni` (String) CNI plug-in to use. Valid options: auto, bridge, calico, cilium, flannel, kindnet, or path to a CNI manifest (default: auto)
- `container_runtime` (String) The container runtime to be used. Valid options: docker, cri-o, containerd (default: auto)
- `cpus` (Number) Amount of CPUs to allocate to Kubernetes
- `cri_socket` (String) The cri socket path to be used.
- `delete_on_failure` (Boolean) If set, delete the current cluster if start fails and try again. Defaults to false.
- `disable_driver_mounts` (Boolean) Disables the filesystem mounts provided by the hypervisors
- `disable_metrics` (Boolean) If set, disables metrics reporting (CPU and memory usage), this can improve CPU usage. Defaults to false.
- `disable_optimizations` (Boolean) If set, disables optimizations that are set for local Kubernetes. Including decreasing CoreDNS replicas from 2 to 1. Defaults to false.
- `disk_size` (String) Disk size allocated to the minikube VM (format: <number>[<unit>], where unit = b, k, m or g).
- `dns_domain` (String) The cluster dns domain name used in the Kubernetes cluster
- `dns_proxy` (Boolean) Enable proxy for NAT DNS requests (virtualbox driver only)
- `docker_env` (List of String) Environment variables to pass to the Docker daemon. (format: key=value)
- `docker_opt` (List of String) Specify arbitrary flags to pass to the Docker daemon. (format: key=value)
- `download_only` (Boolean) If true, only download and cache files for later use - don't install or start anything.
- `driver` (String) Driver is one of the following - Windows: (hyperv, docker, virtualbox, vmware, qemu2, ssh) - OSX: (virtualbox, parallels, vmwarefusion, hyperkit, vmware, qemu2, docker, podman, ssh) - Linux: (docker, kvm2, virtualbox, qemu2, none, podman, ssh)
- `dry_run` (Boolean) dry-run mode. Validates configuration, but does not mutate system state
- `embed_certs` (Boolean) if true, will embed the certs in kubeconfig.
- `enable_default_cni` (Boolean) DEPRECATED: Replaced by --cni=bridge
- `extra_config` (String) A set of key=value pairs that describe configuration that may be passed to different components. 		The key should be '.' separated, and the first part before the dot is the component to apply the configuration to. 		Valid components are: kubelet, kubeadm, apiserver, controller-manager, etcd, proxy, scheduler 		Valid kubeadm parameters: ignore-preflight-errors, dry-run, kubeconfig, kubeconfig-dir, node-name, cri-socket, experimental-upload-certs, certificate-key, rootfs, skip-phases, pod-network-cidr
- `extra_disks` (Number) Number of extra disks created and attached to the minikube VM (currently only implemented for hyperkit and kvm2 drivers)
- `feature_gates` (String) A set of key=value pairs that describe feature gates for alpha/experimental features.
- `force` (Boolean) Force minikube to perform possibly dangerous operations
- `force_systemd` (Boolean) If set, force the container runtime to use systemd as cgroup manager. Defaults to false.
- `host_dns_resolver` (Boolean) Enable host resolver for NAT DNS requests (virtualbox driver only)
- `host_only_cidr` (String) The CIDR to be used for the minikube VM (virtualbox driver only)
- `host_only_nic_type` (String) NIC Type used for host only network. One of Am79C970A, Am79C973, 82540EM, 82543GC, 82545EM, or virtio (virtualbox driver only)
- `hyperkit_vpnkit_sock` (String) Location of the VPNKit socket used for networking. If empty, disables Hyperkit VPNKitSock, if 'auto' uses Docker for Mac VPNKit connection, otherwise uses the specified VSock (hyperkit driver only)
- `hyperkit_vsock_ports` (List of String) List of guest VSock ports that should be exposed as sockets on the host (hyperkit driver only)
- `hyperv_external_adapter` (String) External Adapter on which external switch will be created if no external switch is found. (hyperv driver only)
- `hyperv_use_external_switch` (Boolean) Whether to use external switch over Default Switch if virtual switch not explicitly specified. (hyperv driver only)
- `hyperv_virtual_switch` (String) The hyperv virtual switch name. Defaults to first found. (hyperv driver only)
- `image_mirror_country` (String) Country code of the image mirror to be used. Leave empty to use the global one. For Chinese mainland users, set it to cn.
- `image_repository` (String) Alternative image repository to pull docker images from. This can be used when you have limited access to gcr.io. Set it to "auto" to let minikube decide one for you. For Chinese mainland users, you may use local gcr.io mirrors such as registry.cn-hangzhou.aliyuncs.com/google_containers
- `insecure_registry` (List of String) Insecure Docker registries to pass to the Docker daemon.  The default service CIDR range will automatically be added.
- `install_addons` (Boolean) If set, install addons. Defaults to true.
- `interactive` (Boolean) Allow user prompts for more information
- `iso_url` (List of String) Locations to fetch the minikube ISO from.
- `keep_context` (Boolean) This will keep the existing kubectl context and will create a minikube context.
- `kubernetes_version` (String) The Kubernetes version that the minikube VM will use (ex: v1.2.3, 'stable' for v1.26.1, 'latest' for v1.26.1). Defaults to 'stable'.
- `kvm_gpu` (Boolean) Enable experimental NVIDIA GPU support in minikube
- `kvm_hidden` (Boolean) Hide the hypervisor signature from the guest in minikube (kvm2 driver only)
- `kvm_network` (String) The KVM default network name. (kvm2 driver only)
- `kvm_numa_count` (Number) Simulate numa node count in minikube, supported numa node count range is 1-8 (kvm2 driver only)
- `kvm_qemu_uri` (String) The KVM QEMU connection URI. (kvm2 driver only)
- `listen_address` (String) IP Address to use to expose ports (docker and podman driver only)
- `memory` (String) Amount of RAM to allocate to Kubernetes (format: <number>[<unit>], where unit = b, k, m or g)
- `mount` (Boolean) This will start the mount daemon and automatically mount files into minikube.
- `mount_9p_version` (String) Specify the 9p version that the mount should use
- `mount_gid` (String) Default group id used for the mount
- `mount_ip` (String) Specify the ip that the mount should be setup on
- `mount_msize` (Number) The number of bytes to use for 9p packet payload
- `mount_options` (List of String) Additional mount options, such as cache=fscache
- `mount_port` (Number) Specify the port that the mount should be setup on, where 0 means any free port.
- `mount_string` (String) The argument to pass the minikube mount command on start.
- `mount_type` (String) Specify the mount filesystem type (supported types: 9p)
- `mount_uid` (String) Default user id used for the mount
- `namespace` (String) The named space to activate after start
- `nat_nic_type` (String) NIC Type used for nat network. One of Am79C970A, Am79C973, 82540EM, 82543GC, 82545EM, or virtio (virtualbox driver only)
- `native_ssh` (Boolean) Use native Golang SSH client (default true). Set to 'false' to use the command line 'ssh' command when accessing the docker machine. Useful for the machine drivers when they will not start with 'Waiting for SSH'.
- `network` (String) network to run minikube with. Now it is used by docker/podman and KVM drivers. If left empty, minikube will create a new network.
- `network_plugin` (String) DEPRECATED: Replaced by --cni
- `nfs_share` (List of String) Local folders to share with Guest via NFS mounts (hyperkit driver only)
- `nfs_shares_root` (String) Where to root the NFS Shares, defaults to /nfsshares (hyperkit driver only)
- `no_kubernetes` (Boolean) If set, minikube VM/container will start without starting or configuring Kubernetes. (only works on new clusters)
- `no_vtx_check` (Boolean) Disable checking for the availability of hardware virtualization before the vm is started (virtualbox driver only)
- `nodes` (Number) Amount of nodes in the cluster
- `ports` (List of String) List of ports that should be exposed (docker and podman driver only)
- `preload` (Boolean) If set, download tarball of preloaded images if available to improve start time. Defaults to true.
- `qemu_firmware_path` (String) Path to the qemu firmware file. Defaults: For Linux, the default firmware location. For macOS, the brew installation location. For Windows, C:\Program Files\qemu\share
- `registry_mirror` (List of String) Registry mirrors to pass to the Docker daemon
- `service_cluster_ip_range` (String) The CIDR to be used for service cluster IPs.
- `socket_vmnet_client_path` (String) Path to the socket vmnet client binary (QEMU driver only)
- `socket_vmnet_path` (String) Path to socket vmnet binary (QEMU driver only)
- `ssh_ip_address` (String) IP address (ssh driver only)
- `ssh_key` (String) SSH key (ssh driver only)
- `ssh_port` (Number) SSH port (ssh driver only)
- `ssh_user` (String) SSH user (ssh driver only)
- `static_ip` (String) Set a static IP for the minikube cluster, the IP must be: private, IPv4, and the last octet must be between 2 and 254, for example 192.168.200.200 (Docker and Podman drivers only)
- `subnet` (String) Subnet to be used on kic cluster. If left empty, minikube will choose subnet address, beginning from 192.168.49.0. (docker and podman driver only)
- `trace` (String) Send trace events. Options include: [gcp]
- `uuid` (String) Provide VM UUID to restore MAC address (hyperkit driver only)
- `vm` (Boolean) Filter to use only VM Drivers
- `vm_driver` (String) DEPRECATED, use `driver` instead.
- `wait` (List of String) comma separated list of Kubernetes components to verify and wait for after starting a cluster. defaults to "apiserver,system_pods", available options: "apiserver,system_pods,default_sa,apps_running,node_ready,kubelet" . other acceptable values are 'all' or 'none', 'true' and 'false'
- `wait_timeout` (Number) max time to wait per Kubernetes or host to be healthy. (Configured in minutes)

### Read-Only

- `client_certificate` (String, Sensitive) client certificate used in cluster
- `client_key` (String, Sensitive) client key for cluster
- `cluster_ca_certificate` (String, Sensitive) certificate authority for cluster
- `host` (String) the host name for the cluster
- `id` (String) The ID of this resource.


