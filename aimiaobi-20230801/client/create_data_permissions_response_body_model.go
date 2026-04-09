// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataPermissionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateDataPermissionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDataPermissionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDataPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataPermissionsResponseBody
	GetSuccess() *bool
}

type CreateDataPermissionsResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataPermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataPermissionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDataPermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDataPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataPermissionsResponseBody) SetCode(v string) *CreateDataPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataPermissionsResponseBody) SetHttpStatusCode(v int32) *CreateDataPermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataPermissionsResponseBody) SetMessage(v string) *CreateDataPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataPermissionsResponseBody) SetRequestId(v string) *CreateDataPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataPermissionsResponseBody) SetSuccess(v bool) *CreateDataPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
