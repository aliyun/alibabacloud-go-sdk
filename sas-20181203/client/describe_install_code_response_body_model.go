// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstallCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstallCodeResponseBody
	GetCode() *string
	SetRequestId(v string) *DescribeInstallCodeResponseBody
	GetRequestId() *string
}

type DescribeInstallCodeResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB393****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstallCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstallCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstallCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstallCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstallCodeResponseBody) SetCode(v string) *DescribeInstallCodeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstallCodeResponseBody) SetRequestId(v string) *DescribeInstallCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstallCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
