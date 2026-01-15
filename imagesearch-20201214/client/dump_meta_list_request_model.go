// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDumpMetaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DumpMetaListRequest
	GetId() *int64
	SetInstanceName(v string) *DumpMetaListRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *DumpMetaListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DumpMetaListRequest
	GetPageSize() *int32
}

type DumpMetaListRequest struct {
	// The ID of the export task.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of images to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DumpMetaListRequest) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaListRequest) GoString() string {
	return s.String()
}

func (s *DumpMetaListRequest) GetId() *int64 {
	return s.Id
}

func (s *DumpMetaListRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DumpMetaListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DumpMetaListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DumpMetaListRequest) SetId(v int64) *DumpMetaListRequest {
	s.Id = &v
	return s
}

func (s *DumpMetaListRequest) SetInstanceName(v string) *DumpMetaListRequest {
	s.InstanceName = &v
	return s
}

func (s *DumpMetaListRequest) SetPageNumber(v int32) *DumpMetaListRequest {
	s.PageNumber = &v
	return s
}

func (s *DumpMetaListRequest) SetPageSize(v int32) *DumpMetaListRequest {
	s.PageSize = &v
	return s
}

func (s *DumpMetaListRequest) Validate() error {
	return dara.Validate(s)
}
