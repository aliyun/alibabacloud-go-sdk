// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *AnalyzeLabelRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *AnalyzeLabelRequest
	GetJsonStr() *string
}

type AnalyzeLabelRequest struct {
	// The business workspace ID.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following details.
	//
	// example:
	//
	// “”
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AnalyzeLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeLabelRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeLabelRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *AnalyzeLabelRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *AnalyzeLabelRequest) SetBaseMeAgentId(v int64) *AnalyzeLabelRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AnalyzeLabelRequest) SetJsonStr(v string) *AnalyzeLabelRequest {
	s.JsonStr = &v
	return s
}

func (s *AnalyzeLabelRequest) Validate() error {
	return dara.Validate(s)
}
