// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAddonInstanceResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHelmRelease(v *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) *ListClusterAddonInstanceResourcesResponseBody
	GetHelmRelease() *ListClusterAddonInstanceResourcesResponseBodyHelmRelease
	SetKubernetesObjects(v []*ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) *ListClusterAddonInstanceResourcesResponseBody
	GetKubernetesObjects() []*ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects
}

type ListClusterAddonInstanceResourcesResponseBody struct {
	// Information about the Helm release instance corresponding to the add-on.
	HelmRelease *ListClusterAddonInstanceResourcesResponseBodyHelmRelease `json:"helm_release,omitempty" xml:"helm_release,omitempty" type:"Struct"`
	// A list of Kubernetes objects associated with the add-on.
	KubernetesObjects []*ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects `json:"kubernetes_objects,omitempty" xml:"kubernetes_objects,omitempty" type:"Repeated"`
}

func (s ListClusterAddonInstanceResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstanceResourcesResponseBody) GetHelmRelease() *ListClusterAddonInstanceResourcesResponseBodyHelmRelease {
	return s.HelmRelease
}

func (s *ListClusterAddonInstanceResourcesResponseBody) GetKubernetesObjects() []*ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects {
	return s.KubernetesObjects
}

func (s *ListClusterAddonInstanceResourcesResponseBody) SetHelmRelease(v *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) *ListClusterAddonInstanceResourcesResponseBody {
	s.HelmRelease = v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBody) SetKubernetesObjects(v []*ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) *ListClusterAddonInstanceResourcesResponseBody {
	s.KubernetesObjects = v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBody) Validate() error {
	if s.HelmRelease != nil {
		if err := s.HelmRelease.Validate(); err != nil {
			return err
		}
	}
	if s.KubernetesObjects != nil {
		for _, item := range s.KubernetesObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterAddonInstanceResourcesResponseBodyHelmRelease struct {
	// The name of the Helm chart.
	//
	// example:
	//
	// ack-node-problem-detector
	ChartName *string `json:"chart_name,omitempty" xml:"chart_name,omitempty"`
	// The version of the Helm chart.
	//
	// example:
	//
	// 1.2.28
	ChartVersion *string `json:"chart_version,omitempty" xml:"chart_version,omitempty"`
	// The namespace where the Helm release is located.
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the Helm release instance.
	//
	// example:
	//
	// ack-node-problem-detector
	ReleaseName *string `json:"release_name,omitempty" xml:"release_name,omitempty"`
}

func (s ListClusterAddonInstanceResourcesResponseBodyHelmRelease) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstanceResourcesResponseBodyHelmRelease) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) GetChartName() *string {
	return s.ChartName
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) GetChartVersion() *string {
	return s.ChartVersion
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) GetNamespace() *string {
	return s.Namespace
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) SetChartName(v string) *ListClusterAddonInstanceResourcesResponseBodyHelmRelease {
	s.ChartName = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) SetChartVersion(v string) *ListClusterAddonInstanceResourcesResponseBodyHelmRelease {
	s.ChartVersion = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) SetNamespace(v string) *ListClusterAddonInstanceResourcesResponseBodyHelmRelease {
	s.Namespace = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) SetReleaseName(v string) *ListClusterAddonInstanceResourcesResponseBodyHelmRelease {
	s.ReleaseName = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyHelmRelease) Validate() error {
	return dara.Validate(s)
}

type ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects struct {
	// The Kubernetes API group to which the object belongs.
	//
	// example:
	//
	// rbac.authorization.k8s.io
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The Kubernetes API type to which the object belongs.
	//
	// example:
	//
	// ClusterRole
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The name of the Kubernetes object.
	//
	// example:
	//
	// terway-pod-reader
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The namespace to which the object belongs.
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The Kubernetes API version to which the object belongs.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) GetGroup() *string {
	return s.Group
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) GetKind() *string {
	return s.Kind
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) GetName() *string {
	return s.Name
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) GetNamespace() *string {
	return s.Namespace
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) GetVersion() *string {
	return s.Version
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) SetGroup(v string) *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects {
	s.Group = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) SetKind(v string) *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects {
	s.Kind = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) SetName(v string) *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects {
	s.Name = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) SetNamespace(v string) *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects {
	s.Namespace = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) SetVersion(v string) *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects {
	s.Version = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponseBodyKubernetesObjects) Validate() error {
	return dara.Validate(s)
}
