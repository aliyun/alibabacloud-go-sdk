// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetadata interface {
	dara.Model
	String() string
	GoString() string
	SetAttachments(v []*MetadataAttachments) *Metadata
	GetAttachments() []*MetadataAttachments
}

type Metadata struct {
	Attachments []*MetadataAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
}

func (s Metadata) String() string {
	return dara.Prettify(s)
}

func (s Metadata) GoString() string {
	return s.String()
}

func (s *Metadata) GetAttachments() []*MetadataAttachments {
	return s.Attachments
}

func (s *Metadata) SetAttachments(v []*MetadataAttachments) *Metadata {
	s.Attachments = v
	return s
}

func (s *Metadata) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetadataAttachments struct {
	Arn      *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	MimeType *string `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
}

func (s MetadataAttachments) String() string {
	return dara.Prettify(s)
}

func (s MetadataAttachments) GoString() string {
	return s.String()
}

func (s *MetadataAttachments) GetArn() *string {
	return s.Arn
}

func (s *MetadataAttachments) GetMimeType() *string {
	return s.MimeType
}

func (s *MetadataAttachments) SetArn(v string) *MetadataAttachments {
	s.Arn = &v
	return s
}

func (s *MetadataAttachments) SetMimeType(v string) *MetadataAttachments {
	s.MimeType = &v
	return s
}

func (s *MetadataAttachments) Validate() error {
	return dara.Validate(s)
}
