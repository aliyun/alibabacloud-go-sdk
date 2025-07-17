// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSilencePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteSilencePolicyRequest
	GetId() *int64
}

type DeleteSilencePolicyRequest struct {
	// The ID of the silence policy.
	//
	// For more information about how to obtain the ID of a silence policy, see [ListSilencePolicies](https://help.aliyun.com/document_detail/2612383.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSilencePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSilencePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSilencePolicyRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteSilencePolicyRequest) SetId(v int64) *DeleteSilencePolicyRequest {
	s.Id = &v
	return s
}

func (s *DeleteSilencePolicyRequest) Validate() error {
	return dara.Validate(s)
}
