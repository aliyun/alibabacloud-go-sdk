// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLabelAnalysisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetLabelAnalysisResultRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetLabelAnalysisResultRequest
	GetJsonStr() *string
}

type GetLabelAnalysisResultRequest struct {
	// The business space ID.
	//
	// example:
	//
	// 12345
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// The complete JSON string. For more information, see the following detailed description.
	//
	// example:
	//
	// “”
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetLabelAnalysisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLabelAnalysisResultRequest) GoString() string {
	return s.String()
}

func (s *GetLabelAnalysisResultRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetLabelAnalysisResultRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetLabelAnalysisResultRequest) SetBaseMeAgentId(v int64) *GetLabelAnalysisResultRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetLabelAnalysisResultRequest) SetJsonStr(v string) *GetLabelAnalysisResultRequest {
	s.JsonStr = &v
	return s
}

func (s *GetLabelAnalysisResultRequest) Validate() error {
	return dara.Validate(s)
}
