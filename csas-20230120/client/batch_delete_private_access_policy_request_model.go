// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePrivateAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *BatchDeletePrivateAccessPolicyRequest
	GetPolicyIds() []*string
}

type BatchDeletePrivateAccessPolicyRequest struct {
	// The IDs of internal network access policies. You can specify up to 100 internal network access policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s BatchDeletePrivateAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *BatchDeletePrivateAccessPolicyRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *BatchDeletePrivateAccessPolicyRequest) SetPolicyIds(v []*string) *BatchDeletePrivateAccessPolicyRequest {
	s.PolicyIds = v
	return s
}

func (s *BatchDeletePrivateAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
