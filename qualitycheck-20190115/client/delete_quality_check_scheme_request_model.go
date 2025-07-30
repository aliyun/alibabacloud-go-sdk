// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityCheckSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteQualityCheckSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *DeleteQualityCheckSchemeRequest
	GetJsonStr() *string
}

type DeleteQualityCheckSchemeRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"schemeId":191}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s DeleteQualityCheckSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityCheckSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteQualityCheckSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *DeleteQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *DeleteQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteQualityCheckSchemeRequest) SetJsonStr(v string) *DeleteQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *DeleteQualityCheckSchemeRequest) Validate() error {
	return dara.Validate(s)
}
