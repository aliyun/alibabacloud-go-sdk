// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DescribeDcdnWafPolicyRequest
	GetPolicyId() *int64
}

type DescribeDcdnWafPolicyRequest struct {
	// The ID of the protection policy. You can specify only one ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeDcdnWafPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeDcdnWafPolicyRequest) SetPolicyId(v int64) *DescribeDcdnWafPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeDcdnWafPolicyRequest) Validate() error {
	return dara.Validate(s)
}
