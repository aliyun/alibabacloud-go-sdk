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
	// 	- When the value **true*	- is returned for the **Status*	- parameter, the system does not return the ErrorCode parameter.
	//
	// 	- When the value **false*	- is returned for the **Status*	- parameter, the system returns for the ErrorCode parameter the reason why the ApsaraDB for ClickHouse cluster cannot be connected to the ApsaraDB RDS for MySQL instance.
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
	// Indicates whether the ApsaraDB for ClickHouse cluster can be connected to the ApsaraDB RDS for MySQL instance.
	//
	// 	- **true**: The ApsaraDB for ClickHouse cluster can be connected to the ApsaraDB RDS for MySQL instance.
	//
	// 	- **false**: The ApsaraDB for ClickHouse cluster cannot be connected to the ApsaraDB RDS for MySQL instance.
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
