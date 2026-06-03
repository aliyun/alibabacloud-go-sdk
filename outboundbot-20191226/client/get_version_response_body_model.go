// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVersionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVersionResponseBody
	GetSuccess() *bool
	SetVersion(v string) *GetVersionResponseBody
	GetVersion() *string
}

type GetVersionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1364f208-982d-4d0c-89aa-d56e22b47589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 2018-12-13
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVersionResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetVersionResponseBody) SetCode(v string) *GetVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetVersionResponseBody) SetHttpStatusCode(v int32) *GetVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVersionResponseBody) SetMessage(v string) *GetVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetVersionResponseBody) SetRequestId(v string) *GetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVersionResponseBody) SetSuccess(v bool) *GetVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetVersionResponseBody) SetVersion(v string) *GetVersionResponseBody {
	s.Version = &v
	return s
}

func (s *GetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
