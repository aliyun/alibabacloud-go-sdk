// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckClickhouseToRDSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CheckClickhouseToRDSResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *CheckClickhouseToRDSResponseBody
	GetRequestId() *string
	SetStatus(v bool) *CheckClickhouseToRDSResponseBody
	GetStatus() *bool
}

type CheckClickhouseToRDSResponseBody struct {
	// - This parameter is returned only if the connection fails (**Status*	- is **false**).
	//
	// - It indicates the reason for the connection failure.
	//
	// example:
	//
	// NotSameVpc
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A82758F8-E793-5610-BE11-0E46664305C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether a connection can be established. Valid values:
	//
	// - **true**: A connection can be established.
	//
	// - **false**: A connection cannot be established.
	//
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckClickhouseToRDSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckClickhouseToRDSResponseBody) GoString() string {
	return s.String()
}

func (s *CheckClickhouseToRDSResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckClickhouseToRDSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckClickhouseToRDSResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *CheckClickhouseToRDSResponseBody) SetErrorCode(v string) *CheckClickhouseToRDSResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckClickhouseToRDSResponseBody) SetRequestId(v string) *CheckClickhouseToRDSResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckClickhouseToRDSResponseBody) SetStatus(v bool) *CheckClickhouseToRDSResponseBody {
	s.Status = &v
	return s
}

func (s *CheckClickhouseToRDSResponseBody) Validate() error {
	return dara.Validate(s)
}
