// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableApiKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DisableApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableApiKeyResponseBody
	GetSuccess() *bool
}

type DisableApiKeyResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36045E0A-551D-592D-B1BC-4C56596CE59E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DisableApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableApiKeyResponseBody) SetCode(v string) *DisableApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DisableApiKeyResponseBody) SetHttpStatusCode(v int32) *DisableApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableApiKeyResponseBody) SetMessage(v string) *DisableApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DisableApiKeyResponseBody) SetRequestId(v string) *DisableApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApiKeyResponseBody) SetSuccess(v bool) *DisableApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DisableApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
