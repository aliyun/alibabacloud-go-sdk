// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCapabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpCode(v int64) *CreateCapabilityResponseBody
	GetHttpCode() *int64
	SetRequestId(v string) *CreateCapabilityResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateCapabilityResponseBody
	GetStatus() *string
}

type CreateCapabilityResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1CC93E65-6734-5060-BEF7-0EB0A4862BCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateCapabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCapabilityResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *CreateCapabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCapabilityResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateCapabilityResponseBody) SetHttpCode(v int64) *CreateCapabilityResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateCapabilityResponseBody) SetRequestId(v string) *CreateCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCapabilityResponseBody) SetStatus(v string) *CreateCapabilityResponseBody {
	s.Status = &v
	return s
}

func (s *CreateCapabilityResponseBody) Validate() error {
	return dara.Validate(s)
}
