// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteControlPolicyResponseBody
	GetRequestId() *string
}

type DeleteControlPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteControlPolicyResponseBody) SetRequestId(v string) *DeleteControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
