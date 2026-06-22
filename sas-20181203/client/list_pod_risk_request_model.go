// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPodRiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListPodRiskRequest
	GetAppName() *string
	SetClusterId(v string) *ListPodRiskRequest
	GetClusterId() *string
	SetCurrentPage(v int64) *ListPodRiskRequest
	GetCurrentPage() *int64
	SetNamespace(v string) *ListPodRiskRequest
	GetNamespace() *string
	SetPageSize(v int64) *ListPodRiskRequest
	GetPageSize() *int64
	SetPodName(v string) *ListPodRiskRequest
	GetPodName() *string
}

type ListPodRiskRequest struct {
	// The application name.
	//
	// example:
	//
	// nginx1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the container cluster to query.
	//
	// > You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to obtain this parameter.
	//
	// example:
	//
	// c314aa5b2f208461dad821cdfed82****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the current page when paging is used.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The Kubernetes cluster namespace.
	//
	// example:
	//
	// taas
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of entries per page when paging is used. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pod name.
	//
	// example:
	//
	// abcd-84898334227-p****
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s ListPodRiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPodRiskRequest) GoString() string {
	return s.String()
}

func (s *ListPodRiskRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListPodRiskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPodRiskRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListPodRiskRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListPodRiskRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPodRiskRequest) GetPodName() *string {
	return s.PodName
}

func (s *ListPodRiskRequest) SetAppName(v string) *ListPodRiskRequest {
	s.AppName = &v
	return s
}

func (s *ListPodRiskRequest) SetClusterId(v string) *ListPodRiskRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPodRiskRequest) SetCurrentPage(v int64) *ListPodRiskRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPodRiskRequest) SetNamespace(v string) *ListPodRiskRequest {
	s.Namespace = &v
	return s
}

func (s *ListPodRiskRequest) SetPageSize(v int64) *ListPodRiskRequest {
	s.PageSize = &v
	return s
}

func (s *ListPodRiskRequest) SetPodName(v string) *ListPodRiskRequest {
	s.PodName = &v
	return s
}

func (s *ListPodRiskRequest) Validate() error {
	return dara.Validate(s)
}
