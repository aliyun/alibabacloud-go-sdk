// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHostAccountCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetHostAccountCredentialResponseBody
	GetRequestId() *string
}

type ResetHostAccountCredentialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8783-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetHostAccountCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetHostAccountCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *ResetHostAccountCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetHostAccountCredentialResponseBody) SetRequestId(v string) *ResetHostAccountCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetHostAccountCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
