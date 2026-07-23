// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyEngineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ApplyEngineConfigRequest
	GetInstanceId() *string
}

type ApplyEngineConfigRequest struct {
	// The instance ID. You can obtain this by calling the [ListInstances](https://help.aliyun.com/document_detail/2411819.html) API.
	//
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ApplyEngineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *ApplyEngineConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ApplyEngineConfigRequest) SetInstanceId(v string) *ApplyEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
