// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIDiffAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartAIDiffAnalysisResponseBody
	GetCode() *string
	SetData(v string) *StartAIDiffAnalysisResponseBody
	GetData() *string
	SetMessage(v string) *StartAIDiffAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartAIDiffAnalysisResponseBody
	GetRequestId() *string
}

type StartAIDiffAnalysisResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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

func (s StartAIDiffAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAIDiffAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *StartAIDiffAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartAIDiffAnalysisResponseBody) GetData() *string {
	return s.Data
}

func (s *StartAIDiffAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartAIDiffAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAIDiffAnalysisResponseBody) SetCode(v string) *StartAIDiffAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *StartAIDiffAnalysisResponseBody) SetData(v string) *StartAIDiffAnalysisResponseBody {
	s.Data = &v
	return s
}

func (s *StartAIDiffAnalysisResponseBody) SetMessage(v string) *StartAIDiffAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *StartAIDiffAnalysisResponseBody) SetRequestId(v string) *StartAIDiffAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAIDiffAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
