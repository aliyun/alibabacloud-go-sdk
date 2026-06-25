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
	// The ID of the metadata export task.
	//
	// >To obtain the export task ID, call the [metadata export](https://help.aliyun.com/document_detail/377466.html) operation first and retrieve the ID from the response.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. The instance name must be unique within the same region. Make sure you distinguish between the two.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number of the returned results. Default value: 1. Maximum value: 30.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return. Default value: 10. Maximum value: 30.
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
