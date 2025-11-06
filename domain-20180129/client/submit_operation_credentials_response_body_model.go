// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitOperationCredentialsResponseBody
	GetRequestId() *string
}

type SubmitOperationCredentialsResponseBody struct {
	// example:
	//
	// 9DFCF6F8-243C-40EC-8035-4B12FEFX7D98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitOperationCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitOperationCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitOperationCredentialsResponseBody) SetRequestId(v string) *SubmitOperationCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitOperationCredentialsResponseBody) Validate() error {
	return dara.Validate(s)
}
