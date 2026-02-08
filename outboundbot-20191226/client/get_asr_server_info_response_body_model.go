// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrServerInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAsrServerInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAsrServerInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAsrServerInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAsrServerInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAsrServerInfoResponseBody
	GetSuccess() *bool
}

type GetAsrServerInfoResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsrServerInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsrServerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsrServerInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAsrServerInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAsrServerInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAsrServerInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsrServerInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAsrServerInfoResponseBody) SetCode(v string) *GetAsrServerInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsrServerInfoResponseBody) SetHttpStatusCode(v int32) *GetAsrServerInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAsrServerInfoResponseBody) SetMessage(v string) *GetAsrServerInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsrServerInfoResponseBody) SetRequestId(v string) *GetAsrServerInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsrServerInfoResponseBody) SetSuccess(v bool) *GetAsrServerInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetAsrServerInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
