// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLabelGeneratedResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetLabelGeneratedResultRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetLabelGeneratedResultRequest
	GetJsonStr() *string
}

type GetLabelGeneratedResultRequest struct {
	// The ID of the business workspace.
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

func (s GetLabelGeneratedResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLabelGeneratedResultRequest) GoString() string {
	return s.String()
}

func (s *GetLabelGeneratedResultRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetLabelGeneratedResultRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetLabelGeneratedResultRequest) SetBaseMeAgentId(v int64) *GetLabelGeneratedResultRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetLabelGeneratedResultRequest) SetJsonStr(v string) *GetLabelGeneratedResultRequest {
	s.JsonStr = &v
	return s
}

func (s *GetLabelGeneratedResultRequest) Validate() error {
	return dara.Validate(s)
}
