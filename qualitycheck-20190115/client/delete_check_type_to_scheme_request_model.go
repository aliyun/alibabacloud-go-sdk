// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCheckTypeToSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteCheckTypeToSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteCheckTypeToSchemeRequest
	GetJsonStr() *string
}

type DeleteCheckTypeToSchemeRequest struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":"1376","checkType":"4"}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s DeleteCheckTypeToSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckTypeToSchemeRequest) GoString() string {
	return s.String()
}

func (s *DeleteCheckTypeToSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteCheckTypeToSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteCheckTypeToSchemeRequest) SetBaseMeAgentId(v int64) *DeleteCheckTypeToSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteCheckTypeToSchemeRequest) SetJsonStr(v string) *DeleteCheckTypeToSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteCheckTypeToSchemeRequest) Validate() error {
	return dara.Validate(s)
}
