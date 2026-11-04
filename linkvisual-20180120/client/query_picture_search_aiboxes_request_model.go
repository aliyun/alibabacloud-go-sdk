// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchAiboxesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *QueryPictureSearchAiboxesRequest
	GetAppInstanceId() *string
	SetCurrentPage(v int32) *QueryPictureSearchAiboxesRequest
	GetCurrentPage() *int32
	SetIotInstanceId(v string) *QueryPictureSearchAiboxesRequest
	GetIotInstanceId() *string
	SetPageSize(v int32) *QueryPictureSearchAiboxesRequest
	GetPageSize() *int32
}

type QueryPictureSearchAiboxesRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](https://help.aliyun.com/document_detail/206608.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a502d3fbab8410e8fd4be9037c7****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The page number. Minimum value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance.
	//
	// You can view the instance ID on the **Instance Overview*	- page in the IoT Platform console.
	//
	// > If a public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot_instc_pu****_c*-v64********
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 30. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchAiboxesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchAiboxesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAiboxesRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *QueryPictureSearchAiboxesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureSearchAiboxesRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryPictureSearchAiboxesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureSearchAiboxesRequest) SetAppInstanceId(v string) *QueryPictureSearchAiboxesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetCurrentPage(v int32) *QueryPictureSearchAiboxesRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetIotInstanceId(v string) *QueryPictureSearchAiboxesRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) SetPageSize(v int32) *QueryPictureSearchAiboxesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchAiboxesRequest) Validate() error {
	return dara.Validate(s)
}
