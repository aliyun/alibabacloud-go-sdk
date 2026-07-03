// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHealthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckHealthResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckHealthResponseBody
	GetStatus() *string
}

type CheckHealthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the component is healthy. A return value of UP indicates healthy. Any other value or an empty value indicates an exception.
	//
	// example:
	//
	// UP
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckHealthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckHealthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckHealthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckHealthResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckHealthResponseBody) SetRequestId(v string) *CheckHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckHealthResponseBody) SetStatus(v string) *CheckHealthResponseBody {
	s.Status = &v
	return s
}

func (s *CheckHealthResponseBody) Validate() error {
	return dara.Validate(s)
}
