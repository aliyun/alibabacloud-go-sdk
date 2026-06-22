// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualFile interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *ContextualFile
	GetContentType() *string
	SetDatasetName(v string) *ContextualFile
	GetDatasetName() *string
	SetElements(v []*Element) *ContextualFile
	GetElements() []*Element
	SetMediaType(v string) *ContextualFile
	GetMediaType() *string
	SetOSSURI(v string) *ContextualFile
	GetOSSURI() *string
	SetObjectId(v string) *ContextualFile
	GetObjectId() *string
	SetOwnerId(v string) *ContextualFile
	GetOwnerId() *string
	SetProjectName(v string) *ContextualFile
	GetProjectName() *string
	SetURI(v string) *ContextualFile
	GetURI() *string
}

type ContextualFile struct {
	// The Multipurpose Internet Mail Extensions (MIME) type of the file.
	//
	// example:
	//
	// text/x-imm-faq
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// Elements.
	Elements []*Element `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// The media type of the file.
	//
	// example:
	//
	// document
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The URI path of the OSS file. This parameter is used only when the URI is a PDS address.
	//
	// example:
	//
	// oss://test-bucket/test-object.jpg
	OSSURI *string `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	// The identifier of the file in the dataset.
	//
	// example:
	//
	// 0939d7ed-73fa-4009-bbe6-fbbe07b92b2e
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1482910009923706
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The project name.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URI of the file.
	//
	// The format of an OSS URI is oss\\://${bucketname}/${objectname}. ${bucketname} is the name of an OSS bucket in the same region as the current project. ${objectname} is the file path.
	//
	// The format of a PDS URI is pds\\://domains/${domain}/drives/${drive}/files/${file}/revisions/${revision}.
	//
	// example:
	//
	// oss://test-bucket
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s ContextualFile) String() string {
	return dara.Prettify(s)
}

func (s ContextualFile) GoString() string {
	return s.String()
}

func (s *ContextualFile) GetContentType() *string {
	return s.ContentType
}

func (s *ContextualFile) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ContextualFile) GetElements() []*Element {
	return s.Elements
}

func (s *ContextualFile) GetMediaType() *string {
	return s.MediaType
}

func (s *ContextualFile) GetOSSURI() *string {
	return s.OSSURI
}

func (s *ContextualFile) GetObjectId() *string {
	return s.ObjectId
}

func (s *ContextualFile) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ContextualFile) GetProjectName() *string {
	return s.ProjectName
}

func (s *ContextualFile) GetURI() *string {
	return s.URI
}

func (s *ContextualFile) SetContentType(v string) *ContextualFile {
	s.ContentType = &v
	return s
}

func (s *ContextualFile) SetDatasetName(v string) *ContextualFile {
	s.DatasetName = &v
	return s
}

func (s *ContextualFile) SetElements(v []*Element) *ContextualFile {
	s.Elements = v
	return s
}

func (s *ContextualFile) SetMediaType(v string) *ContextualFile {
	s.MediaType = &v
	return s
}

func (s *ContextualFile) SetOSSURI(v string) *ContextualFile {
	s.OSSURI = &v
	return s
}

func (s *ContextualFile) SetObjectId(v string) *ContextualFile {
	s.ObjectId = &v
	return s
}

func (s *ContextualFile) SetOwnerId(v string) *ContextualFile {
	s.OwnerId = &v
	return s
}

func (s *ContextualFile) SetProjectName(v string) *ContextualFile {
	s.ProjectName = &v
	return s
}

func (s *ContextualFile) SetURI(v string) *ContextualFile {
	s.URI = &v
	return s
}

func (s *ContextualFile) Validate() error {
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
