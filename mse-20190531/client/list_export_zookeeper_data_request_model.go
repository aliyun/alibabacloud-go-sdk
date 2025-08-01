// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExportZookeeperDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListExportZookeeperDataRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *ListExportZookeeperDataRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListExportZookeeperDataRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExportZookeeperDataRequest
	GetPageSize() *int32
}

type ListExportZookeeperDataRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-7pp2d1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 0
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExportZookeeperDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExportZookeeperDataRequest) GoString() string {
	return s.String()
}

func (s *ListExportZookeeperDataRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListExportZookeeperDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExportZookeeperDataRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExportZookeeperDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExportZookeeperDataRequest) SetAcceptLanguage(v string) *ListExportZookeeperDataRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListExportZookeeperDataRequest) SetInstanceId(v string) *ListExportZookeeperDataRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExportZookeeperDataRequest) SetPageNumber(v int32) *ListExportZookeeperDataRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExportZookeeperDataRequest) SetPageSize(v int32) *ListExportZookeeperDataRequest {
	s.PageSize = &v
	return s
}

func (s *ListExportZookeeperDataRequest) Validate() error {
	return dara.Validate(s)
}
