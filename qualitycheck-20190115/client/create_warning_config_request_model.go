// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarningConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateWarningConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateWarningConfigRequest
	GetJsonStr() *string
}

type CreateWarningConfigRequest struct {
	// Workspace ID
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// Complete JSON string information. See the following details for specific content.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ridList":[15670],"configName":"投诉预警","channels":[{"type":1,"url":"https://xxx"}]}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s CreateWarningConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateWarningConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateWarningConfigRequest) SetBaseMeAgentId(v int64) *CreateWarningConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateWarningConfigRequest) SetJsonStr(v string) *CreateWarningConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateWarningConfigRequest) Validate() error {
	return dara.Validate(s)
}
