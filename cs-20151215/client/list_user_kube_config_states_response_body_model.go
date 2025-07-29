// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserKubeConfigStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v *ListUserKubeConfigStatesResponseBodyPage) *ListUserKubeConfigStatesResponseBody
	GetPage() *ListUserKubeConfigStatesResponseBodyPage
	SetStates(v []*ListUserKubeConfigStatesResponseBodyStates) *ListUserKubeConfigStatesResponseBody
	GetStates() []*ListUserKubeConfigStatesResponseBodyStates
}

type ListUserKubeConfigStatesResponseBody struct {
	// The pagination information.
	Page *ListUserKubeConfigStatesResponseBodyPage `json:"page,omitempty" xml:"page,omitempty" type:"Struct"`
	// The status of the kubeconfig files.
	States []*ListUserKubeConfigStatesResponseBodyStates `json:"states,omitempty" xml:"states,omitempty" type:"Repeated"`
}

func (s ListUserKubeConfigStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserKubeConfigStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserKubeConfigStatesResponseBody) GetPage() *ListUserKubeConfigStatesResponseBodyPage {
	return s.Page
}

func (s *ListUserKubeConfigStatesResponseBody) GetStates() []*ListUserKubeConfigStatesResponseBodyStates {
	return s.States
}

func (s *ListUserKubeConfigStatesResponseBody) SetPage(v *ListUserKubeConfigStatesResponseBodyPage) *ListUserKubeConfigStatesResponseBody {
	s.Page = v
	return s
}

func (s *ListUserKubeConfigStatesResponseBody) SetStates(v []*ListUserKubeConfigStatesResponseBodyStates) *ListUserKubeConfigStatesResponseBody {
	s.States = v
	return s
}

func (s *ListUserKubeConfigStatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserKubeConfigStatesResponseBodyPage struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s ListUserKubeConfigStatesResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListUserKubeConfigStatesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListUserKubeConfigStatesResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserKubeConfigStatesResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserKubeConfigStatesResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserKubeConfigStatesResponseBodyPage) SetPageNumber(v int32) *ListUserKubeConfigStatesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListUserKubeConfigStatesResponseBodyPage) SetPageSize(v int32) *ListUserKubeConfigStatesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListUserKubeConfigStatesResponseBodyPage) SetTotalCount(v int32) *ListUserKubeConfigStatesResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *ListUserKubeConfigStatesResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type ListUserKubeConfigStatesResponseBodyStates struct {
	// The expiration date of the certificate used in a kubeconfig file. Format: the UTC time in the RFC3339 format.
	//
	// example:
	//
	// 2026-11-30T07:41:50Z
	CertExpireTime *string `json:"cert_expire_time,omitempty" xml:"cert_expire_time,omitempty"`
	// The current status of the certificate used in a kubeconfig file. Valid values:
	//
	// 	- Expired: The certificate is expired.
	//
	// 	- Unexpired: The certificate is not expired.
	//
	// 	- Unissued: The certificate is not issued.
	//
	// 	- Unknown: The status of the certificate is unknown.
	//
	// 	- Removed: The certificate is removed. An issue record is found for the certificate.
	//
	// example:
	//
	// Unissued
	CertState *string `json:"cert_state,omitempty" xml:"cert_state,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c5b5e80b0b64a4bf6939d2d8fbbc5****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The name of the cluster.
	//
	// The name must be 1 to 63 characters in length, and can contain digits, underscores (_), and hyphens (-). The name must start with a letter or number.
	//
	// example:
	//
	// cluster-demo
	ClusterName *string `json:"cluster_name,omitempty" xml:"cluster_name,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- `initial`: The cluster is being created.
	//
	// 	- `failed`: The cluster failed to be created.
	//
	// 	- `running`: The cluster is running.
	//
	// 	- `updating`: The cluster is being upgraded.
	//
	// 	- `updating_failed`: The cluster failed to be updated.
	//
	// 	- `scaling`: The cluster is being scaled.
	//
	// 	- `stopped`: The cluster is stopped.
	//
	// 	- `deleting`: The cluster is being deleted.
	//
	// 	- `deleted`: The cluster is deleted.
	//
	// 	- `delete_failed`: The cluster failed to be deleted.
	//
	// example:
	//
	// running
	ClusterState *string `json:"cluster_state,omitempty" xml:"cluster_state,omitempty"`
}

func (s ListUserKubeConfigStatesResponseBodyStates) String() string {
	return dara.Prettify(s)
}

func (s ListUserKubeConfigStatesResponseBodyStates) GoString() string {
	return s.String()
}

func (s *ListUserKubeConfigStatesResponseBodyStates) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *ListUserKubeConfigStatesResponseBodyStates) GetCertState() *string {
	return s.CertState
}

func (s *ListUserKubeConfigStatesResponseBodyStates) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListUserKubeConfigStatesResponseBodyStates) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListUserKubeConfigStatesResponseBodyStates) GetClusterState() *string {
	return s.ClusterState
}

func (s *ListUserKubeConfigStatesResponseBodyStates) SetCertExpireTime(v string) *ListUserKubeConfigStatesResponseBodyStates {
	s.CertExpireTime = &v
	return s
}

func (s *ListUserKubeConfigStatesResponseBodyStates) SetCertState(v string) *ListUserKubeConfigStatesResponseBodyStates {
	s.CertState = &v
	return s
}

func (s *ListUserKubeConfigStatesResponseBodyStates) SetClusterId(v string) *ListUserKubeConfigStatesResponseBodyStates {
	s.ClusterId = &v
	return s
}

func (s *ListUserKubeConfigStatesResponseBodyStates) SetClusterName(v string) *ListUserKubeConfigStatesResponseBodyStates {
	s.ClusterName = &v
	return s
}

func (s *ListUserKubeConfigStatesResponseBodyStates) SetClusterState(v string) *ListUserKubeConfigStatesResponseBodyStates {
	s.ClusterState = &v
	return s
}

func (s *ListUserKubeConfigStatesResponseBodyStates) Validate() error {
	return dara.Validate(s)
}
