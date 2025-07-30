// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckTypeToSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateCheckTypeToSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateCheckTypeToSchemeRequest
	GetJsonStr() *string
}

type CreateCheckTypeToSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s CreateCheckTypeToSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckTypeToSchemeRequest) GoString() string {
	return s.String()
}

func (s *CreateCheckTypeToSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateCheckTypeToSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateCheckTypeToSchemeRequest) SetBaseMeAgentId(v int64) *CreateCheckTypeToSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateCheckTypeToSchemeRequest) SetJsonStr(v string) *CreateCheckTypeToSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateCheckTypeToSchemeRequest) Validate() error {
	return dara.Validate(s)
}
