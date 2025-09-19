// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContainerScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateContainerScanTaskRequest
	GetClusterId() *string
	SetContainerIds(v string) *CreateContainerScanTaskRequest
	GetContainerIds() *string
	SetLang(v string) *CreateContainerScanTaskRequest
	GetLang() *string
}

type CreateContainerScanTaskRequest struct {
	// The ID of the cluster to which the container belongs.
	//
	// > You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to query the IDs of clusters.
	//
	// example:
	//
	// c22143730ab6e40b09ec7c1c51d4d****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// c927cf190e886696db53cda0efb57145394ccf0bf9f525353fa5c22a26e4****
	ContainerIds *string `json:"ContainerIds,omitempty" xml:"ContainerIds,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CreateContainerScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContainerScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateContainerScanTaskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateContainerScanTaskRequest) GetContainerIds() *string {
	return s.ContainerIds
}

func (s *CreateContainerScanTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateContainerScanTaskRequest) SetClusterId(v string) *CreateContainerScanTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateContainerScanTaskRequest) SetContainerIds(v string) *CreateContainerScanTaskRequest {
	s.ContainerIds = &v
	return s
}

func (s *CreateContainerScanTaskRequest) SetLang(v string) *CreateContainerScanTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateContainerScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
