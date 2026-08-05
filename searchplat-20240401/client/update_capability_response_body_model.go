// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCapabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpCode(v int64) *UpdateCapabilityResponseBody
	GetHttpCode() *int64
	SetRequestId(v string) *UpdateCapabilityResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateCapabilityResponseBody
	GetStatus() *string
}

type UpdateCapabilityResponseBody struct {
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
	// 5950143C-B8F0-5758-A08A-66F302FD587F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateCapabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCapabilityResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *UpdateCapabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCapabilityResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateCapabilityResponseBody) SetHttpCode(v int64) *UpdateCapabilityResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateCapabilityResponseBody) SetRequestId(v string) *UpdateCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCapabilityResponseBody) SetStatus(v string) *UpdateCapabilityResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateCapabilityResponseBody) Validate() error {
	return dara.Validate(s)
}
