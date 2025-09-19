// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCheckPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*int64) *DeleteCheckPolicyRequest
	GetPolicyIds() []*int64
	SetPolicyType(v string) *DeleteCheckPolicyRequest
	GetPolicyType() *string
}

type DeleteCheckPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*int64 `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DeleteCheckPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCheckPolicyRequest) GetPolicyIds() []*int64 {
	return s.PolicyIds
}

func (s *DeleteCheckPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteCheckPolicyRequest) SetPolicyIds(v []*int64) *DeleteCheckPolicyRequest {
	s.PolicyIds = v
	return s
}

func (s *DeleteCheckPolicyRequest) SetPolicyType(v string) *DeleteCheckPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *DeleteCheckPolicyRequest) Validate() error {
	return dara.Validate(s)
}
