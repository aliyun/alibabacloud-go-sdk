// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DeletePolicyV2Request
	GetPolicyId() *string
}

type DeletePolicyV2Request struct {
	// The ID of the backup policy.
	//
	// example:
	//
	// po-000************2l6
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeletePolicyV2Request) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyV2Request) GoString() string {
	return s.String()
}

func (s *DeletePolicyV2Request) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeletePolicyV2Request) SetPolicyId(v string) *DeletePolicyV2Request {
	s.PolicyId = &v
	return s
}

func (s *DeletePolicyV2Request) Validate() error {
	return dara.Validate(s)
}
