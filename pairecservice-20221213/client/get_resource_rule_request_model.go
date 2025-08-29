// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetResourceRuleRequest
	GetInstanceId() *string
}

type GetResourceRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetResourceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetResourceRuleRequest) SetInstanceId(v string) *GetResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetResourceRuleRequest) Validate() error {
	return dara.Validate(s)
}
