// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBasePreSignedUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocuments(v []*GetKnowledgeBasePreSignedUrlRequestDocuments) *GetKnowledgeBasePreSignedUrlRequest
	GetDocuments() []*GetKnowledgeBasePreSignedUrlRequestDocuments
	SetExpiresIn(v int32) *GetKnowledgeBasePreSignedUrlRequest
	GetExpiresIn() *int32
	SetKnowledgeBaseId(v string) *GetKnowledgeBasePreSignedUrlRequest
	GetKnowledgeBaseId() *string
}

type GetKnowledgeBasePreSignedUrlRequest struct {
	Documents []*GetKnowledgeBasePreSignedUrlRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// example:
	//
	// 3600
	ExpiresIn *int32 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// example:
	//
	// kb-3bd02617e9be335f
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
}

func (s GetKnowledgeBasePreSignedUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBasePreSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBasePreSignedUrlRequest) GetDocuments() []*GetKnowledgeBasePreSignedUrlRequestDocuments {
	return s.Documents
}

func (s *GetKnowledgeBasePreSignedUrlRequest) GetExpiresIn() *int32 {
	return s.ExpiresIn
}

func (s *GetKnowledgeBasePreSignedUrlRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBasePreSignedUrlRequest) SetDocuments(v []*GetKnowledgeBasePreSignedUrlRequestDocuments) *GetKnowledgeBasePreSignedUrlRequest {
	s.Documents = v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlRequest) SetExpiresIn(v int32) *GetKnowledgeBasePreSignedUrlRequest {
	s.ExpiresIn = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlRequest) SetKnowledgeBaseId(v string) *GetKnowledgeBasePreSignedUrlRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlRequest) Validate() error {
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKnowledgeBasePreSignedUrlRequestDocuments struct {
	// example:
	//
	// CHANGELOG.md
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 本地上传时为预签名上传使用的批次相对路径；不同 ImportType 下含义由导入类型定义。
	//
	// example:
	//
	// contract-2026.md
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 1024
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetKnowledgeBasePreSignedUrlRequestDocuments) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBasePreSignedUrlRequestDocuments) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBasePreSignedUrlRequestDocuments) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeBasePreSignedUrlRequestDocuments) GetPath() *string {
	return s.Path
}

func (s *GetKnowledgeBasePreSignedUrlRequestDocuments) GetSize() *int64 {
	return s.Size
}

func (s *GetKnowledgeBasePreSignedUrlRequestDocuments) SetName(v string) *GetKnowledgeBasePreSignedUrlRequestDocuments {
	s.Name = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlRequestDocuments) SetPath(v string) *GetKnowledgeBasePreSignedUrlRequestDocuments {
	s.Path = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlRequestDocuments) SetSize(v int64) *GetKnowledgeBasePreSignedUrlRequestDocuments {
	s.Size = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlRequestDocuments) Validate() error {
	return dara.Validate(s)
}
