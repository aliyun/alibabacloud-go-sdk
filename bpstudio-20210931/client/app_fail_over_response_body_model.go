// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppFailOverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AppFailOverResponseBody
	GetCode() *string
	SetData(v int32) *AppFailOverResponseBody
	GetData() *int32
	SetMessage(v string) *AppFailOverResponseBody
	GetMessage() *string
	SetRequestId(v string) *AppFailOverResponseBody
	GetRequestId() *string
}

type AppFailOverResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The disaster recovery switchover task ID.
	//
	// example:
	//
	// 7441
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9656C816-1E9A-58D2-86D5-710678D61AF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppFailOverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AppFailOverResponseBody) GoString() string {
	return s.String()
}

func (s *AppFailOverResponseBody) GetCode() *string {
	return s.Code
}

func (s *AppFailOverResponseBody) GetData() *int32 {
	return s.Data
}

func (s *AppFailOverResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AppFailOverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AppFailOverResponseBody) SetCode(v string) *AppFailOverResponseBody {
	s.Code = &v
	return s
}

func (s *AppFailOverResponseBody) SetData(v int32) *AppFailOverResponseBody {
	s.Data = &v
	return s
}

func (s *AppFailOverResponseBody) SetMessage(v string) *AppFailOverResponseBody {
	s.Message = &v
	return s
}

func (s *AppFailOverResponseBody) SetRequestId(v string) *AppFailOverResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppFailOverResponseBody) Validate() error {
	return dara.Validate(s)
}
