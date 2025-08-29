// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRuleItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteResourceRuleItemRequest
	GetInstanceId() *string
}

type DeleteResourceRuleItemRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteResourceRuleItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRuleItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleItemRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteResourceRuleItemRequest) SetInstanceId(v string) *DeleteResourceRuleItemRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteResourceRuleItemRequest) Validate() error {
	return dara.Validate(s)
}
