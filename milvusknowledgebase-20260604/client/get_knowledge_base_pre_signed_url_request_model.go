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
	// The list of files to be uploaded. You can specify 1 to 100 files.
	Documents []*GetKnowledgeBasePreSignedUrlRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The validity period of the pre-signed URL in seconds. Default value: `3600`.
	//
	// example:
	//
	// 3600
	ExpiresIn *int32 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// The ID of the knowledge base. This parameter and datasetId are mutually exclusive. If both are specified, this parameter takes precedence.
	//
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
	// The display name of the file. If not specified, the file name from Path is used by default.
	//
	// example:
	//
	// CHANGELOG.md
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The **file name or relative path*	- for local upload scenarios. The path cannot start with `direct_upload/` or `uploaded/`, cannot contain empty segments, `.`, or `..`, and must be less than or equal to 1,024 bytes in length.
	//
	// example:
	//
	// contract-2026.md
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the file in bytes.
	//
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
