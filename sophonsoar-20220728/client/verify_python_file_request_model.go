// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPythonFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *VerifyPythonFileRequest
	GetContent() *string
}

type VerifyPythonFileRequest struct {
	// The Python code snippet.
	//
	// This parameter is required.
	//
	// example:
	//
	// import logging
	//
	// def execute (params):
	//
	//   success=True
	//
	//   message=\\"OK\\"
	//
	//   data=[]
	//
	//   return (success,message,data)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s VerifyPythonFileRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyPythonFileRequest) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileRequest) GetContent() *string {
	return s.Content
}

func (s *VerifyPythonFileRequest) SetContent(v string) *VerifyPythonFileRequest {
	s.Content = &v
	return s
}

func (s *VerifyPythonFileRequest) Validate() error {
	return dara.Validate(s)
}
