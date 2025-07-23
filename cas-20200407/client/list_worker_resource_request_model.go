// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkerResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudProduct(v string) *ListWorkerResourceRequest
	GetCloudProduct() *string
	SetCurrentPage(v int32) *ListWorkerResourceRequest
	GetCurrentPage() *int32
	SetJobId(v int64) *ListWorkerResourceRequest
	GetJobId() *int64
	SetShowSize(v int32) *ListWorkerResourceRequest
	GetShowSize() *int32
	SetStatus(v string) *ListWorkerResourceRequest
	GetStatus() *string
}

type ListWorkerResourceRequest struct {
	// The cloud service in the deployment task.
	//
	// example:
	//
	// NLB
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the **ID*	- parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The number of entries per page. Default value: 50.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The status of the worker task.
	//
	// Valid values:
	//
	// 	- rollback
	//
	// 	- rollback_error
	//
	// 	- success
	//
	// 	- rollback_success
	//
	// 	- pending
	//
	// 	- scheduling
	//
	// 	- processing
	//
	// 	- error
	//
	// 	- editing
	//
	// example:
	//
	// editing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWorkerResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkerResourceRequest) GoString() string {
	return s.String()
}

func (s *ListWorkerResourceRequest) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *ListWorkerResourceRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListWorkerResourceRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListWorkerResourceRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListWorkerResourceRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWorkerResourceRequest) SetCloudProduct(v string) *ListWorkerResourceRequest {
	s.CloudProduct = &v
	return s
}

func (s *ListWorkerResourceRequest) SetCurrentPage(v int32) *ListWorkerResourceRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListWorkerResourceRequest) SetJobId(v int64) *ListWorkerResourceRequest {
	s.JobId = &v
	return s
}

func (s *ListWorkerResourceRequest) SetShowSize(v int32) *ListWorkerResourceRequest {
	s.ShowSize = &v
	return s
}

func (s *ListWorkerResourceRequest) SetStatus(v string) *ListWorkerResourceRequest {
	s.Status = &v
	return s
}

func (s *ListWorkerResourceRequest) Validate() error {
	return dara.Validate(s)
}
