// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostWebShellResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginUrl(v string) *DescribeHostWebShellResponseBody
	GetLoginUrl() *string
	SetRequestId(v string) *DescribeHostWebShellResponseBody
	GetRequestId() *string
}

type DescribeHostWebShellResponseBody struct {
	// The webshell URL.
	//
	// example:
	//
	// ***
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 842B73C8-5776-4BD9-9872-69C8C46DD7D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHostWebShellResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostWebShellResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostWebShellResponseBody) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *DescribeHostWebShellResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHostWebShellResponseBody) SetLoginUrl(v string) *DescribeHostWebShellResponseBody {
	s.LoginUrl = &v
	return s
}

func (s *DescribeHostWebShellResponseBody) SetRequestId(v string) *DescribeHostWebShellResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHostWebShellResponseBody) Validate() error {
	return dara.Validate(s)
}
