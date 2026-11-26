// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityCheckSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListQualityCheckSchemeRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListQualityCheckSchemeRequest
	GetJsonStr() *string
}

type ListQualityCheckSchemeRequest struct {
	// The workspace ID.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For details, see the following information.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//       "name": "Quality Check Scheme A"
	//
	// }
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListQualityCheckSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityCheckSchemeRequest) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListQualityCheckSchemeRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListQualityCheckSchemeRequest) SetBaseMeAgentId(v int64) *ListQualityCheckSchemeRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListQualityCheckSchemeRequest) SetJsonStr(v string) *ListQualityCheckSchemeRequest {
	s.JsonStr = &v
	return s
}

func (s *ListQualityCheckSchemeRequest) Validate() error {
	return dara.Validate(s)
}
