// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultPolicyVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultPolicyVersionResponseBody
	GetRequestId() *string
}

type SetDefaultPolicyVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultPolicyVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultPolicyVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultPolicyVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultPolicyVersionResponseBody) SetRequestId(v string) *SetDefaultPolicyVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultPolicyVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
