// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccessList(v []*ListCloudAccessResponseBodyCloudAccessList) *ListCloudAccessResponseBody
	GetCloudAccessList() []*ListCloudAccessResponseBodyCloudAccessList
	SetCurrentPage(v int32) *ListCloudAccessResponseBody
	GetCurrentPage() *int32
	SetRequestId(v string) *ListCloudAccessResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListCloudAccessResponseBody
	GetShowSize() *int32
	SetTotalCount(v int32) *ListCloudAccessResponseBody
	GetTotalCount() *int32
}

type ListCloudAccessResponseBody struct {
	// The query results.
	CloudAccessList []*ListCloudAccessResponseBodyCloudAccessList `json:"CloudAccessList,omitempty" xml:"CloudAccessList,omitempty" type:"Repeated"`
	// The default value is the current page. If CurrentPage is not specified, this parameter is not returned.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D3F1FA43-1C26-50A2-8F0F-7A03851DBB46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries per page. If ShowSize is not specified, this parameter is not returned.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccessResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAccessResponseBody) GetCloudAccessList() []*ListCloudAccessResponseBodyCloudAccessList {
	return s.CloudAccessList
}

func (s *ListCloudAccessResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAccessResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCloudAccessResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCloudAccessResponseBody) SetCloudAccessList(v []*ListCloudAccessResponseBodyCloudAccessList) *ListCloudAccessResponseBody {
	s.CloudAccessList = v
	return s
}

func (s *ListCloudAccessResponseBody) SetCurrentPage(v int32) *ListCloudAccessResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudAccessResponseBody) SetRequestId(v string) *ListCloudAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAccessResponseBody) SetShowSize(v int32) *ListCloudAccessResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCloudAccessResponseBody) SetTotalCount(v int32) *ListCloudAccessResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAccessResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCloudAccessResponseBodyCloudAccessList struct {
	// The ID of the primary key.
	//
	// example:
	//
	// 888
	AccessId *int64 `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The cloud service provider.
	//
	// example:
	//
	// Tencent
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The AccessKey ID that is used to access cloud resources.
	//
	// example:
	//
	// AAAqdwPBA****
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The service status. The value normal indicates that the service runs as expected.
	//
	// example:
	//
	// normal
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s ListCloudAccessResponseBodyCloudAccessList) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccessResponseBodyCloudAccessList) GoString() string {
	return s.String()
}

func (s *ListCloudAccessResponseBodyCloudAccessList) GetAccessId() *int64 {
	return s.AccessId
}

func (s *ListCloudAccessResponseBodyCloudAccessList) GetCloudName() *string {
	return s.CloudName
}

func (s *ListCloudAccessResponseBodyCloudAccessList) GetSecretId() *string {
	return s.SecretId
}

func (s *ListCloudAccessResponseBodyCloudAccessList) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListCloudAccessResponseBodyCloudAccessList) SetAccessId(v int64) *ListCloudAccessResponseBodyCloudAccessList {
	s.AccessId = &v
	return s
}

func (s *ListCloudAccessResponseBodyCloudAccessList) SetCloudName(v string) *ListCloudAccessResponseBodyCloudAccessList {
	s.CloudName = &v
	return s
}

func (s *ListCloudAccessResponseBodyCloudAccessList) SetSecretId(v string) *ListCloudAccessResponseBodyCloudAccessList {
	s.SecretId = &v
	return s
}

func (s *ListCloudAccessResponseBodyCloudAccessList) SetServiceStatus(v string) *ListCloudAccessResponseBodyCloudAccessList {
	s.ServiceStatus = &v
	return s
}

func (s *ListCloudAccessResponseBodyCloudAccessList) Validate() error {
	return dara.Validate(s)
}
