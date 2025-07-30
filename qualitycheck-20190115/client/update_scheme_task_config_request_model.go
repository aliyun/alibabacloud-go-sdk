// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSchemeTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateSchemeTaskConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateSchemeTaskConfigRequest
	GetJsonStr() *string
}

type UpdateSchemeTaskConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateSchemeTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSchemeTaskConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateSchemeTaskConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *UpdateSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateSchemeTaskConfigRequest) SetJsonStr(v string) *UpdateSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateSchemeTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
