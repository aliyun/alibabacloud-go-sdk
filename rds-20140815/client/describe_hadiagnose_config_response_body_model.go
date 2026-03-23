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
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
