// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDataPermissionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDataPermissionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDataPermissionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDataPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataPermissionsResponseBody
	GetSuccess() *bool
}

type DeleteDataPermissionsResponseBody struct {
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

func (s DeleteDataPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataPermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDataPermissionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDataPermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDataPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataPermissionsResponseBody) SetCode(v string) *DeleteDataPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataPermissionsResponseBody) SetHttpStatusCode(v int32) *DeleteDataPermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDataPermissionsResponseBody) SetMessage(v string) *DeleteDataPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataPermissionsResponseBody) SetRequestId(v string) *DeleteDataPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataPermissionsResponseBody) SetSuccess(v bool) *DeleteDataPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
