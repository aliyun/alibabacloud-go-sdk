// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckTypeToSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateCheckTypeToSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateCheckTypeToSchemeRequest
	GetJsonStr() *string
}

type UpdateCheckTypeToSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateCheckTypeToSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckTypeToSchemeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckTypeToSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateCheckTypeToSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateCheckTypeToSchemeRequest) SetBaseMeAgentId(v int64) *UpdateCheckTypeToSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateCheckTypeToSchemeRequest) SetJsonStr(v string) *UpdateCheckTypeToSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateCheckTypeToSchemeRequest) Validate() error {
	return dara.Validate(s)
}
