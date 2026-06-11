// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillImportFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *GetSkillImportFileUrlRequest
	GetContentType() *string
	SetNamespaceId(v string) *GetSkillImportFileUrlRequest
	GetNamespaceId() *string
}

type GetSkillImportFileUrlRequest struct {
	// example:
	//
	// application/zip
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetSkillImportFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillImportFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetSkillImportFileUrlRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GetSkillImportFileUrlRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetSkillImportFileUrlRequest) SetContentType(v string) *GetSkillImportFileUrlRequest {
	s.ContentType = &v
	return s
}

func (s *GetSkillImportFileUrlRequest) SetNamespaceId(v string) *GetSkillImportFileUrlRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetSkillImportFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
