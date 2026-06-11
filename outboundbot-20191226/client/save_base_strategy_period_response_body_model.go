// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBaseStrategyPeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveBaseStrategyPeriodResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SaveBaseStrategyPeriodResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveBaseStrategyPeriodResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveBaseStrategyPeriodResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveBaseStrategyPeriodResponseBody
	GetSuccess() *bool
}

type SaveBaseStrategyPeriodResponseBody struct {
	// Response code
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
	// Response message
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

func (s SaveBaseStrategyPeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBaseStrategyPeriodResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBaseStrategyPeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveBaseStrategyPeriodResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveBaseStrategyPeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveBaseStrategyPeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBaseStrategyPeriodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveBaseStrategyPeriodResponseBody) SetCode(v string) *SaveBaseStrategyPeriodResponseBody {
	s.Code = &v
	return s
}

func (s *SaveBaseStrategyPeriodResponseBody) SetHttpStatusCode(v int32) *SaveBaseStrategyPeriodResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveBaseStrategyPeriodResponseBody) SetMessage(v string) *SaveBaseStrategyPeriodResponseBody {
	s.Message = &v
	return s
}

func (s *SaveBaseStrategyPeriodResponseBody) SetRequestId(v string) *SaveBaseStrategyPeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBaseStrategyPeriodResponseBody) SetSuccess(v bool) *SaveBaseStrategyPeriodResponseBody {
	s.Success = &v
	return s
}

func (s *SaveBaseStrategyPeriodResponseBody) Validate() error {
	return dara.Validate(s)
}
