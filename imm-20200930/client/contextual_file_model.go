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
	// Name of the dataset
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// Elements
	Elements []*Element `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// Media type of the current file
	//
	// Valid values:
	//
	// 	- image
	//
	// 	- other
	//
	// 	- document
	//
	// 	- archive
	//
	// 	- audio
	//
	// 	- video
	//
	// example:
	//
	// document
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The URI of the OSS object. This parameter is available only if the value of the URI parameter is the URI of a file in Photo and Drive Service.
	//
	// example:
	//
	// oss://test-bucket/test-object.jpg
	OSSURI *string `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	// The identifier of the corresponding file that exists in the dataset.
	//
	// example:
	//
	// 0939d7ed-73fa-4009-bbe6-fbbe07b92b2e
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// User ID
	//
	// example:
	//
	// 1482910009923706
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Name of the project
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// URI of the file. Specify the OSS URI in the oss://${bucketname}/${objectname} format, where ${bucketname} is the name of the bucket in the same region as the current project and ${objectname} is the path of the object. The URI of a file in Photo and Drive Service follows the pds://domains/${domain}/drives/${drive}/files/${file}/revisions/${revision} format.
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
