// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDocumentResult interface {
	dara.Model
	String() string
	GoString() string
	SetDocName(v string) *AddDocumentResult
	GetDocName() *string
	SetDocumentInfo(v *DocumentInfo) *AddDocumentResult
	GetDocumentInfo() *DocumentInfo
	SetErrorMessage(v string) *AddDocumentResult
	GetErrorMessage() *string
	SetSuccess(v bool) *AddDocumentResult
	GetSuccess() *bool
}

type AddDocumentResult struct {
	// example:
	//
	// example.pdf
	DocName      *string       `json:"docName,omitempty" xml:"docName,omitempty"`
	DocumentInfo *DocumentInfo `json:"documentInfo,omitempty" xml:"documentInfo,omitempty"`
	// example:
	//
	// true
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDocumentResult) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentResult) GoString() string {
	return s.String()
}

func (s *AddDocumentResult) GetDocName() *string {
	return s.DocName
}

func (s *AddDocumentResult) GetDocumentInfo() *DocumentInfo {
	return s.DocumentInfo
}

func (s *AddDocumentResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddDocumentResult) GetSuccess() *bool {
	return s.Success
}

func (s *AddDocumentResult) SetDocName(v string) *AddDocumentResult {
	s.DocName = &v
	return s
}

func (s *AddDocumentResult) SetDocumentInfo(v *DocumentInfo) *AddDocumentResult {
	s.DocumentInfo = v
	return s
}

func (s *AddDocumentResult) SetErrorMessage(v string) *AddDocumentResult {
	s.ErrorMessage = &v
	return s
}

func (s *AddDocumentResult) SetSuccess(v bool) *AddDocumentResult {
	s.Success = &v
	return s
}

func (s *AddDocumentResult) Validate() error {
	if s.DocumentInfo != nil {
		if err := s.DocumentInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
