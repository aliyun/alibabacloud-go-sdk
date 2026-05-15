// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTagListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetQualityRuleTagListRequest
	GetInstanceId() *string
}

type GetQualityRuleTagListRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetQualityRuleTagListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTagListRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTagListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQualityRuleTagListRequest) SetInstanceId(v string) *GetQualityRuleTagListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityRuleTagListRequest) Validate() error {
	return dara.Validate(s)
}
