// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *QueryPictureSearchDevicesRequest
	GetAppInstanceId() *string
	SetCurrentPage(v int32) *QueryPictureSearchDevicesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QueryPictureSearchDevicesRequest
	GetPageSize() *int32
}

type QueryPictureSearchDevicesRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](https://help.aliyun.com/document_detail/206608.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a502d3fbab8410e8fd4be9037c7****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The page number of the page to return. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Minimum value: 1. Maximum value: 50. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchDevicesRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchDevicesRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *QueryPictureSearchDevicesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureSearchDevicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureSearchDevicesRequest) SetAppInstanceId(v string) *QueryPictureSearchDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) SetCurrentPage(v int32) *QueryPictureSearchDevicesRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) SetPageSize(v int32) *QueryPictureSearchDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchDevicesRequest) Validate() error {
	return dara.Validate(s)
}
