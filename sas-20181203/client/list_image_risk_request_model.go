// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListImageRiskRequest
	GetAppName() *string
	SetClusterId(v string) *ListImageRiskRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *ListImageRiskRequest
	GetCurrentPage() *int32
	SetImageName(v string) *ListImageRiskRequest
	GetImageName() *string
	SetNamespace(v string) *ListImageRiskRequest
	GetNamespace() *string
	SetPageSize(v int32) *ListImageRiskRequest
	GetPageSize() *int32
}

type ListImageRiskRequest struct {
	// The name of the application.
	//
	// example:
	//
	// e****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the ID of the container cluster.
	//
	// example:
	//
	// c80f79959fd724a888e1187779b13****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// container-***:****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// kube-sy****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImageRiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageRiskRequest) GoString() string {
	return s.String()
}

func (s *ListImageRiskRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListImageRiskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListImageRiskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListImageRiskRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ListImageRiskRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListImageRiskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageRiskRequest) SetAppName(v string) *ListImageRiskRequest {
	s.AppName = &v
	return s
}

func (s *ListImageRiskRequest) SetClusterId(v string) *ListImageRiskRequest {
	s.ClusterId = &v
	return s
}

func (s *ListImageRiskRequest) SetCurrentPage(v int32) *ListImageRiskRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListImageRiskRequest) SetImageName(v string) *ListImageRiskRequest {
	s.ImageName = &v
	return s
}

func (s *ListImageRiskRequest) SetNamespace(v string) *ListImageRiskRequest {
	s.Namespace = &v
	return s
}

func (s *ListImageRiskRequest) SetPageSize(v int32) *ListImageRiskRequest {
	s.PageSize = &v
	return s
}

func (s *ListImageRiskRequest) Validate() error {
	return dara.Validate(s)
}
