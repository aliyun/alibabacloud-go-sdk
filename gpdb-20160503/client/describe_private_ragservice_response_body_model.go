// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateRAGServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaCert(v string) *DescribePrivateRAGServiceResponseBody
	GetCaCert() *string
	SetRequestId(v string) *DescribePrivateRAGServiceResponseBody
	GetRequestId() *string
}

type DescribePrivateRAGServiceResponseBody struct {
	// CA certificate Info. The returned OSS link, valid for 2 hours.
	//
	// example:
	//
	// https://oss-xxx
	CaCert *string `json:"CaCert,omitempty" xml:"CaCert,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrivateRAGServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateRAGServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrivateRAGServiceResponseBody) GetCaCert() *string {
	return s.CaCert
}

func (s *DescribePrivateRAGServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrivateRAGServiceResponseBody) SetCaCert(v string) *DescribePrivateRAGServiceResponseBody {
	s.CaCert = &v
	return s
}

func (s *DescribePrivateRAGServiceResponseBody) SetRequestId(v string) *DescribePrivateRAGServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrivateRAGServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
