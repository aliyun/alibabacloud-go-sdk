// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventCenterRuleNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListEventCenterRuleNameRequest
	GetInstanceId() *string
}

type ListEventCenterRuleNameRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListEventCenterRuleNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventCenterRuleNameRequest) GoString() string {
	return s.String()
}

func (s *ListEventCenterRuleNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEventCenterRuleNameRequest) SetInstanceId(v string) *ListEventCenterRuleNameRequest {
	s.InstanceId = &v
	return s
}

func (s *ListEventCenterRuleNameRequest) Validate() error {
	return dara.Validate(s)
}
