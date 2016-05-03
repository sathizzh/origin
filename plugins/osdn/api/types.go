package api

import (
	knetwork "k8s.io/kubernetes/pkg/kubelet/network"
	pconfig "k8s.io/kubernetes/pkg/proxy/config"
)

type OsdnMasterPlugin interface {
	Start(clusterNetworkCIDR string, clusterBitsPerSubnet uint, serviceNetworkCIDR string) error
}

type OsdnNodePlugin interface {
	knetwork.NetworkPlugin

	Start(mtu uint) error
}

type FilteringEndpointsConfigHandler interface {
	pconfig.EndpointsConfigHandler
	Start(baseHandler pconfig.EndpointsConfigHandler) error
}
