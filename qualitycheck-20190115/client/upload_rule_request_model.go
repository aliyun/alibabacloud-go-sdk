// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UploadRuleRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UploadRuleRequest
	GetJsonStr() *string
}

type UploadRuleRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {“conditions”:“xxxxx”,"rules":"xxxx"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UploadRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadRuleRequest) GoString() string {
	return s.String()
}

func (s *UploadRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UploadRuleRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UploadRuleRequest) SetBaseMeAgentId(v int64) *UploadRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UploadRuleRequest) SetJsonStr(v string) *UploadRuleRequest {
	s.JsonStr = &v
	return s
}

func (s *UploadRuleRequest) Validate() error {
	return dara.Validate(s)
}
