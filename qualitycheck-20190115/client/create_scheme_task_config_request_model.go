// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemeTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateSchemeTaskConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateSchemeTaskConfigRequest
	GetJsonStr() *string
}

type CreateSchemeTaskConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s CreateSchemeTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSchemeTaskConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateSchemeTaskConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *CreateSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateSchemeTaskConfigRequest) SetJsonStr(v string) *CreateSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateSchemeTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
