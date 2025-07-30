// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemeTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteSchemeTaskConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteSchemeTaskConfigRequest
	GetJsonStr() *string
}

type DeleteSchemeTaskConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":"329"}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s DeleteSchemeTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchemeTaskConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteSchemeTaskConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *DeleteSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteSchemeTaskConfigRequest) SetJsonStr(v string) *DeleteSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteSchemeTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
