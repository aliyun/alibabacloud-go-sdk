// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseS3AccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLakebaseS3AccountResponseBody
	GetRequestId() *string
	SetUserAccAk(v string) *DescribeLakebaseS3AccountResponseBody
	GetUserAccAk() *string
	SetUserAccPolicy(v string) *DescribeLakebaseS3AccountResponseBody
	GetUserAccPolicy() *string
	SetUserAccSk(v string) *DescribeLakebaseS3AccountResponseBody
	GetUserAccSk() *string
}

type DescribeLakebaseS3AccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Access Key of the S3 account.
	//
	// example:
	//
	// accname
	UserAccAk *string `json:"UserAccAk,omitempty" xml:"UserAccAk,omitempty"`
	// The policy document of the S3 account.
	//
	// example:
	//
	// {"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["s3:*"],"Resource":["*"]}]}
	UserAccPolicy *string `json:"UserAccPolicy,omitempty" xml:"UserAccPolicy,omitempty"`
	// The Secret Key of the S3 account (masked).
	//
	// example:
	//
	// password***
	UserAccSk *string `json:"UserAccSk,omitempty" xml:"UserAccSk,omitempty"`
}

func (s DescribeLakebaseS3AccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseS3AccountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseS3AccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLakebaseS3AccountResponseBody) GetUserAccAk() *string {
	return s.UserAccAk
}

func (s *DescribeLakebaseS3AccountResponseBody) GetUserAccPolicy() *string {
	return s.UserAccPolicy
}

func (s *DescribeLakebaseS3AccountResponseBody) GetUserAccSk() *string {
	return s.UserAccSk
}

func (s *DescribeLakebaseS3AccountResponseBody) SetRequestId(v string) *DescribeLakebaseS3AccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLakebaseS3AccountResponseBody) SetUserAccAk(v string) *DescribeLakebaseS3AccountResponseBody {
	s.UserAccAk = &v
	return s
}

func (s *DescribeLakebaseS3AccountResponseBody) SetUserAccPolicy(v string) *DescribeLakebaseS3AccountResponseBody {
	s.UserAccPolicy = &v
	return s
}

func (s *DescribeLakebaseS3AccountResponseBody) SetUserAccSk(v string) *DescribeLakebaseS3AccountResponseBody {
	s.UserAccSk = &v
	return s
}

func (s *DescribeLakebaseS3AccountResponseBody) Validate() error {
	return dara.Validate(s)
}
