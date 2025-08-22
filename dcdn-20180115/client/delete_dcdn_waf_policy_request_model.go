// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnWafPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DeleteDcdnWafPolicyRequest
	GetPolicyId() *int64
}

type DeleteDcdnWafPolicyRequest struct {
	// The ID of the protection policy that you want to delete. You can specify only one ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeleteDcdnWafPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnWafPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnWafPolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DeleteDcdnWafPolicyRequest) SetPolicyId(v int64) *DeleteDcdnWafPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteDcdnWafPolicyRequest) Validate() error {
	return dara.Validate(s)
}
