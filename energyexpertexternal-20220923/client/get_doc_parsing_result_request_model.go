// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocParsingResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReturnFormat(v string) *GetDocParsingResultRequest
	GetReturnFormat() *string
	SetTaskId(v string) *GetDocParsingResultRequest
	GetTaskId() *string
}

type GetDocParsingResultRequest struct {
	// - The document parsing result supports two formats: markdown and json.
	//
	// - By default, the result is returned in markdown format.
	//
	// example:
	//
	// md
	ReturnFormat *string `json:"returnFormat,omitempty" xml:"returnFormat,omitempty"`
	// - Task ID.
	//
	// - The taskId is obtained from the SubmitDocParsingTaskAdvance or SubmitDocParsingTask interfaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2c22388d-e2ed-44fe-99e6-99922f15e7bb
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetDocParsingResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocParsingResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocParsingResultRequest) GetReturnFormat() *string {
	return s.ReturnFormat
}

func (s *GetDocParsingResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocParsingResultRequest) SetReturnFormat(v string) *GetDocParsingResultRequest {
	s.ReturnFormat = &v
	return s
}

func (s *GetDocParsingResultRequest) SetTaskId(v string) *GetDocParsingResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetDocParsingResultRequest) Validate() error {
	return dara.Validate(s)
}
