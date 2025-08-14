// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PermissionCheckResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *PermissionCheckResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *PermissionCheckResponseBody
	GetMessage() *string
	SetRequestId(v string) *PermissionCheckResponseBody
	GetRequestId() *string
	SetResultObject(v string) *PermissionCheckResponseBody
	GetResultObject() *string
}

type PermissionCheckResponseBody struct {
	// Return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s PermissionCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PermissionCheckResponseBody) GoString() string {
	return s.String()
}

func (s *PermissionCheckResponseBody) GetCode() *string {
	return s.Code
}

func (s *PermissionCheckResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PermissionCheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PermissionCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PermissionCheckResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *PermissionCheckResponseBody) SetCode(v string) *PermissionCheckResponseBody {
	s.Code = &v
	return s
}

func (s *PermissionCheckResponseBody) SetHttpStatusCode(v string) *PermissionCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PermissionCheckResponseBody) SetMessage(v string) *PermissionCheckResponseBody {
	s.Message = &v
	return s
}

func (s *PermissionCheckResponseBody) SetRequestId(v string) *PermissionCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *PermissionCheckResponseBody) SetResultObject(v string) *PermissionCheckResponseBody {
	s.ResultObject = &v
	return s
}

func (s *PermissionCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
