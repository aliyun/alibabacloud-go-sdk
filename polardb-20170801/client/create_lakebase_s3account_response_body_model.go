// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakebaseS3AccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLakebaseS3AccountResponseBody
	GetRequestId() *string
}

type CreateLakebaseS3AccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLakebaseS3AccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLakebaseS3AccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLakebaseS3AccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLakebaseS3AccountResponseBody) SetRequestId(v string) *CreateLakebaseS3AccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLakebaseS3AccountResponseBody) Validate() error {
	return dara.Validate(s)
}
