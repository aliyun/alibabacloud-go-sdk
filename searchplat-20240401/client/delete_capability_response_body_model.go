// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCapabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpCode(v int64) *DeleteCapabilityResponseBody
	GetHttpCode() *int64
	SetRequestId(v string) *DeleteCapabilityResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteCapabilityResponseBody
	GetStatus() *string
}

type DeleteCapabilityResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7CC54C38-D721-4C55-A410-2A94B5A6BE0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteCapabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCapabilityResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *DeleteCapabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCapabilityResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteCapabilityResponseBody) SetHttpCode(v int64) *DeleteCapabilityResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteCapabilityResponseBody) SetRequestId(v string) *DeleteCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCapabilityResponseBody) SetStatus(v string) *DeleteCapabilityResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteCapabilityResponseBody) Validate() error {
	return dara.Validate(s)
}
