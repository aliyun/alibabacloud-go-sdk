// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenterPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCenterPolicyResponseBody
	GetRequestId() *string
}

type DeleteCenterPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 72E47B1E-6B11-5A11-A27C-7A80F866****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenterPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenterPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCenterPolicyResponseBody) SetRequestId(v string) *DeleteCenterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCenterPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
