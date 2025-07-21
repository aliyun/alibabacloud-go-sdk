// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionByCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetPermissionByCodeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetPermissionByCodeResponseBody
	GetCode() *string
	SetMessage(v string) *GetPermissionByCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPermissionByCodeResponseBody
	GetRequestId() *string
}

type GetPermissionByCodeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error description information.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPermissionByCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionByCodeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetPermissionByCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPermissionByCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPermissionByCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPermissionByCodeResponseBody) SetAccessDeniedDetail(v string) *GetPermissionByCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetPermissionByCodeResponseBody) SetCode(v string) *GetPermissionByCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetPermissionByCodeResponseBody) SetMessage(v string) *GetPermissionByCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetPermissionByCodeResponseBody) SetRequestId(v string) *GetPermissionByCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPermissionByCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
