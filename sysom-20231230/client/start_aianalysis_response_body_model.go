// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartAIAnalysisResponseBody
	GetCode() *string
	SetData(v *StartAIAnalysisResponseBodyData) *StartAIAnalysisResponseBody
	GetData() *StartAIAnalysisResponseBodyData
	SetMessage(v string) *StartAIAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartAIAnalysisResponseBody
	GetRequestId() *string
}

type StartAIAnalysisResponseBody struct {
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *StartAIAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartAIAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAIAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartAIAnalysisResponseBody) GetData() *StartAIAnalysisResponseBodyData {
	return s.Data
}

func (s *StartAIAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartAIAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAIAnalysisResponseBody) SetCode(v string) *StartAIAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *StartAIAnalysisResponseBody) SetData(v *StartAIAnalysisResponseBodyData) *StartAIAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *StartAIAnalysisResponseBody) SetMessage(v string) *StartAIAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *StartAIAnalysisResponseBody) SetRequestId(v string) *StartAIAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAIAnalysisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartAIAnalysisResponseBodyData struct {
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysis_id,omitempty" xml:"analysis_id,omitempty"`
}

func (s StartAIAnalysisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartAIAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisResponseBodyData) GetAnalysisId() *string {
	return s.AnalysisId
}

func (s *StartAIAnalysisResponseBodyData) SetAnalysisId(v string) *StartAIAnalysisResponseBodyData {
	s.AnalysisId = &v
	return s
}

func (s *StartAIAnalysisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
