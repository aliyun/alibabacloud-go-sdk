// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateMaterialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileList(v string) *GetTemplateMaterialsRequest
	GetFileList() *string
	SetTemplateId(v string) *GetTemplateMaterialsRequest
	GetTemplateId() *string
}

type GetTemplateMaterialsRequest struct {
	// The materials that you want to query.
	//
	// example:
	//
	// ["music.mp3","config.json","assets/1.jpg"]
	FileList *string `json:"FileList,omitempty" xml:"FileList,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateMaterialsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateMaterialsRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateMaterialsRequest) GetFileList() *string {
	return s.FileList
}

func (s *GetTemplateMaterialsRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateMaterialsRequest) SetFileList(v string) *GetTemplateMaterialsRequest {
	s.FileList = &v
	return s
}

func (s *GetTemplateMaterialsRequest) SetTemplateId(v string) *GetTemplateMaterialsRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateMaterialsRequest) Validate() error {
	return dara.Validate(s)
}
