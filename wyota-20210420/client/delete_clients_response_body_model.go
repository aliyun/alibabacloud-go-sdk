// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteClientsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteClientsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteClientsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteClientsResponseBody
	GetSuccess() *bool
}

type DeleteClientsResponseBody struct {
	// example:
	//
	// TERMINAL_NOT_FOUND
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// terminal not found
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C5DCE54A-B266-522E-A6ED-468AF45F5AAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteClientsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteClientsResponseBody) SetCode(v string) *DeleteClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClientsResponseBody) SetHttpStatusCode(v int32) *DeleteClientsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteClientsResponseBody) SetMessage(v string) *DeleteClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClientsResponseBody) SetRequestId(v string) *DeleteClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientsResponseBody) SetSuccess(v bool) *DeleteClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteClientsResponseBody) Validate() error {
	return dara.Validate(s)
}
