// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityCheckSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateQualityCheckSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateQualityCheckSchemeRequest
	GetJsonStr() *string
}

type CreateQualityCheckSchemeRequest struct {
	// Workspace ID
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// JSON request body. For details, see the parameter description in the request parameters section.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"dataType":1,"name":"质检方案A","type":1}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s CreateQualityCheckSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateQualityCheckSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *CreateQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateQualityCheckSchemeRequest) SetJsonStr(v string) *CreateQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateQualityCheckSchemeRequest) Validate() error {
	return dara.Validate(s)
}
