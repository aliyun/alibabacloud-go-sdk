// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityCheckSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetQualityCheckSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetQualityCheckSchemeRequest
	GetJsonStr() *string
}

type GetQualityCheckSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":"187","ruleRequireInfos":["BusinessNameInfo","RuleCategory"]}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s GetQualityCheckSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetQualityCheckSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *GetQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetQualityCheckSchemeRequest) SetJsonStr(v string) *GetQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *GetQualityCheckSchemeRequest) Validate() error {
	return dara.Validate(s)
}
