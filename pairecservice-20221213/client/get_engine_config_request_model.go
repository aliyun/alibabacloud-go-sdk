// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetEngineConfigRequest
	GetInstanceId() *string
}

type GetEngineConfigRequest struct {
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEngineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *GetEngineConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEngineConfigRequest) SetInstanceId(v string) *GetEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
