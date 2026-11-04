// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryPictureSearchAppsRequest
	GetCurrentPage() *int32
	SetIotInstanceId(v string) *QueryPictureSearchAppsRequest
	GetIotInstanceId() *string
	SetPageSize(v int32) *QueryPictureSearchAppsRequest
	GetPageSize() *int32
}

type QueryPictureSearchAppsRequest struct {
	// The number of the page to return. Minimum value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// > If a public instance does not have an ID, refer to [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot_instc_pu****_c*-v64********
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchAppsRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchAppsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureSearchAppsRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryPictureSearchAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureSearchAppsRequest) SetCurrentPage(v int32) *QueryPictureSearchAppsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchAppsRequest) SetIotInstanceId(v string) *QueryPictureSearchAppsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryPictureSearchAppsRequest) SetPageSize(v int32) *QueryPictureSearchAppsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchAppsRequest) Validate() error {
	return dara.Validate(s)
}
