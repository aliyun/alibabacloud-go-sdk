// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteResourceRuleRequest
	GetInstanceId() *string
}

type DeleteResourceRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteResourceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteResourceRuleRequest) SetInstanceId(v string) *DeleteResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteResourceRuleRequest) Validate() error {
	return dara.Validate(s)
}
