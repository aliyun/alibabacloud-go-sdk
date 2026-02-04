// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGlobalizationConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceGlobalizationConfigRequest
	GetInstanceId() *string
}

type GetInstanceGlobalizationConfigRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceGlobalizationConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGlobalizationConfigRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceGlobalizationConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceGlobalizationConfigRequest) SetInstanceId(v string) *GetInstanceGlobalizationConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceGlobalizationConfigRequest) Validate() error {
	return dara.Validate(s)
}
