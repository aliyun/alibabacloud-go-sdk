// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityCheckSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateQualityCheckSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateQualityCheckSchemeRequest
	GetJsonStr() *string
}

type UpdateQualityCheckSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	JsonStr       *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s UpdateQualityCheckSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateQualityCheckSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *UpdateQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateQualityCheckSchemeRequest) SetJsonStr(v string) *UpdateQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateQualityCheckSchemeRequest) Validate() error {
	return dara.Validate(s)
}
