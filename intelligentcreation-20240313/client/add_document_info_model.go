// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDocumentInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentType(v string) *AddDocumentInfo
	GetDocumentType() *string
	SetName(v string) *AddDocumentInfo
	GetName() *string
	SetUrl(v string) *AddDocumentInfo
	GetUrl() *string
}

type AddDocumentInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// pdf
	DocumentType *string `json:"documentType,omitempty" xml:"documentType,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddDocumentInfo) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentInfo) GoString() string {
	return s.String()
}

func (s *AddDocumentInfo) GetDocumentType() *string {
	return s.DocumentType
}

func (s *AddDocumentInfo) GetName() *string {
	return s.Name
}

func (s *AddDocumentInfo) GetUrl() *string {
	return s.Url
}

func (s *AddDocumentInfo) SetDocumentType(v string) *AddDocumentInfo {
	s.DocumentType = &v
	return s
}

func (s *AddDocumentInfo) SetName(v string) *AddDocumentInfo {
	s.Name = &v
	return s
}

func (s *AddDocumentInfo) SetUrl(v string) *AddDocumentInfo {
	s.Url = &v
	return s
}

func (s *AddDocumentInfo) Validate() error {
	return dara.Validate(s)
}
