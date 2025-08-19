// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePayerForAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePayerForAccountResponseBody
	GetRequestId() *string
}

type UpdatePayerForAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePayerForAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePayerForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePayerForAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePayerForAccountResponseBody) SetRequestId(v string) *UpdatePayerForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePayerForAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
