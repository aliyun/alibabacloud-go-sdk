// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContent interface {
	dara.Model
	String() string
	GoString() string
	SetMessageAttachments(v []*AguiMessage) *Content
	GetMessageAttachments() []*AguiMessage
	SetQueryAttachments(v []*QueryAttachment) *Content
	GetQueryAttachments() []*QueryAttachment
	SetTextAttachments(v []*string) *Content
	GetTextAttachments() []*string
}

type Content struct {
	MessageAttachments []*AguiMessage     `json:"MessageAttachments,omitempty" xml:"MessageAttachments,omitempty" type:"Repeated"`
	QueryAttachments   []*QueryAttachment `json:"QueryAttachments,omitempty" xml:"QueryAttachments,omitempty" type:"Repeated"`
	TextAttachments    []*string          `json:"TextAttachments,omitempty" xml:"TextAttachments,omitempty" type:"Repeated"`
}

func (s Content) String() string {
	return dara.Prettify(s)
}

func (s Content) GoString() string {
	return s.String()
}

func (s *Content) GetMessageAttachments() []*AguiMessage {
	return s.MessageAttachments
}

func (s *Content) GetQueryAttachments() []*QueryAttachment {
	return s.QueryAttachments
}

func (s *Content) GetTextAttachments() []*string {
	return s.TextAttachments
}

func (s *Content) SetMessageAttachments(v []*AguiMessage) *Content {
	s.MessageAttachments = v
	return s
}

func (s *Content) SetQueryAttachments(v []*QueryAttachment) *Content {
	s.QueryAttachments = v
	return s
}

func (s *Content) SetTextAttachments(v []*string) *Content {
	s.TextAttachments = v
	return s
}

func (s *Content) Validate() error {
	if s.MessageAttachments != nil {
		for _, item := range s.MessageAttachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryAttachments != nil {
		for _, item := range s.QueryAttachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
