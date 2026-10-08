// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHADiagnoseConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeHADiagnoseConfigResponseBody
	GetRequestId() *string
	SetTcpConnectionType(v string) *DescribeHADiagnoseConfigResponseBody
	GetTcpConnectionType() *string
}

type DescribeHADiagnoseConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06B220E2-EAC5-4DBE-A1FC-1B62DB6A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The availability check method that Alibaba Cloud uses for the ApsaraDB RDS instance. Valid values:
	//
	// - **LONG**: persistent connection.
	//
	// - **SHORT**: short-lived connection.
	//
	// example:
	//
	// LONG
	TcpConnectionType *string `json:"TcpConnectionType,omitempty" xml:"TcpConnectionType,omitempty"`
}

func (s DescribeHADiagnoseConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHADiagnoseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHADiagnoseConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHADiagnoseConfigResponseBody) GetTcpConnectionType() *string {
	return s.TcpConnectionType
}

func (s *DescribeHADiagnoseConfigResponseBody) SetRequestId(v string) *DescribeHADiagnoseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHADiagnoseConfigResponseBody) SetTcpConnectionType(v string) *DescribeHADiagnoseConfigResponseBody {
	s.TcpConnectionType = &v
	return s
}

func (s *DescribeHADiagnoseConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
