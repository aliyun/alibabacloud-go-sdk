// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarningConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateWarningConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateWarningConfigRequest
	GetJsonStr() *string
}

type UpdateWarningConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"ridList":[18130],"configName":"0310","channels":[{"type":1,"url":"https://sca.console.aliyun.com/#/warningConfig"}],"configId":29}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateWarningConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateWarningConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateWarningConfigRequest) SetBaseMeAgentId(v int64) *UpdateWarningConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateWarningConfigRequest) SetJsonStr(v string) *UpdateWarningConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateWarningConfigRequest) Validate() error {
	return dara.Validate(s)
}
