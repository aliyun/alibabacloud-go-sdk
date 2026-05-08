// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentType(v string) *DocumentInfo
	GetDocumentType() *string
	SetId(v string) *DocumentInfo
	GetId() *string
	SetName(v string) *DocumentInfo
	GetName() *string
	SetProcessStatus(v string) *DocumentInfo
	GetProcessStatus() *string
}

type DocumentInfo struct {
	DocumentType  *string `json:"documentType,omitempty" xml:"documentType,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	ProcessStatus *string `json:"processStatus,omitempty" xml:"processStatus,omitempty"`
}

func (s DocumentInfo) String() string {
	return dara.Prettify(s)
}

func (s DocumentInfo) GoString() string {
	return s.String()
}

func (s *DocumentInfo) GetDocumentType() *string {
	return s.DocumentType
}

func (s *DocumentInfo) GetId() *string {
	return s.Id
}

func (s *DocumentInfo) GetName() *string {
	return s.Name
}

func (s *DocumentInfo) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *DocumentInfo) SetDocumentType(v string) *DocumentInfo {
	s.DocumentType = &v
	return s
}

func (s *DocumentInfo) SetId(v string) *DocumentInfo {
	s.Id = &v
	return s
}

func (s *DocumentInfo) SetName(v string) *DocumentInfo {
	s.Name = &v
	return s
}

func (s *DocumentInfo) SetProcessStatus(v string) *DocumentInfo {
	s.ProcessStatus = &v
	return s
}

func (s *DocumentInfo) Validate() error {
	return dara.Validate(s)
}
