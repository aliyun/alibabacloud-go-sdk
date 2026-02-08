// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTTSDemoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditionUrl(v string) *DescribeTTSDemoResponseBody
	GetAuditionUrl() *string
	SetCode(v string) *DescribeTTSDemoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeTTSDemoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeTTSDemoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTTSDemoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTTSDemoResponseBody
	GetSuccess() *bool
}

type DescribeTTSDemoResponseBody struct {
	// Audition file URL
	//
	// example:
	//
	// http://XXX/XXX
	AuditionUrl *string `json:"AuditionUrl,omitempty" xml:"AuditionUrl,omitempty"`
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

func (s DescribeTTSDemoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSDemoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTTSDemoResponseBody) GetAuditionUrl() *string {
	return s.AuditionUrl
}

func (s *DescribeTTSDemoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTTSDemoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeTTSDemoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTTSDemoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTTSDemoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTTSDemoResponseBody) SetAuditionUrl(v string) *DescribeTTSDemoResponseBody {
	s.AuditionUrl = &v
	return s
}

func (s *DescribeTTSDemoResponseBody) SetCode(v string) *DescribeTTSDemoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTTSDemoResponseBody) SetHttpStatusCode(v int32) *DescribeTTSDemoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTTSDemoResponseBody) SetMessage(v string) *DescribeTTSDemoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTTSDemoResponseBody) SetRequestId(v string) *DescribeTTSDemoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTTSDemoResponseBody) SetSuccess(v bool) *DescribeTTSDemoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTTSDemoResponseBody) Validate() error {
	return dara.Validate(s)
}
