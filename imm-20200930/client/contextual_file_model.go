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
	ContentType *string    `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DatasetName *string    `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Elements    []*Element `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	MediaType   *string    `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OSSURI      *string    `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	ObjectId    *string    `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	OwnerId     *string    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName *string    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	URI         *string    `json:"URI,omitempty" xml:"URI,omitempty"`
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
	return dara.Validate(s)
}
