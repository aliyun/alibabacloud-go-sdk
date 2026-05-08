// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInteractTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *InteractTextResponseBody
	GetEnd() *bool
	SetIndex(v int32) *InteractTextResponseBody
	GetIndex() *int32
	SetMessage(v string) *InteractTextResponseBody
	GetMessage() *string
	SetRelatedImages(v []*string) *InteractTextResponseBody
	GetRelatedImages() []*string
	SetRelatedVideos(v []*string) *InteractTextResponseBody
	GetRelatedVideos() []*string
	SetSessionId(v string) *InteractTextResponseBody
	GetSessionId() *string
	SetType(v int32) *InteractTextResponseBody
	GetType() *int32
}

type InteractTextResponseBody struct {
	// example:
	//
	// false
	End *bool `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	Index         *int32    `json:"index,omitempty" xml:"index,omitempty"`
	Message       *string   `json:"message,omitempty" xml:"message,omitempty"`
	RelatedImages []*string `json:"relatedImages,omitempty" xml:"relatedImages,omitempty" type:"Repeated"`
	RelatedVideos []*string `json:"relatedVideos,omitempty" xml:"relatedVideos,omitempty" type:"Repeated"`
	// example:
	//
	// 79e954faffe2415ebd18188ba787d78e
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InteractTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InteractTextResponseBody) GoString() string {
	return s.String()
}

func (s *InteractTextResponseBody) GetEnd() *bool {
	return s.End
}

func (s *InteractTextResponseBody) GetIndex() *int32 {
	return s.Index
}

func (s *InteractTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InteractTextResponseBody) GetRelatedImages() []*string {
	return s.RelatedImages
}

func (s *InteractTextResponseBody) GetRelatedVideos() []*string {
	return s.RelatedVideos
}

func (s *InteractTextResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *InteractTextResponseBody) GetType() *int32 {
	return s.Type
}

func (s *InteractTextResponseBody) SetEnd(v bool) *InteractTextResponseBody {
	s.End = &v
	return s
}

func (s *InteractTextResponseBody) SetIndex(v int32) *InteractTextResponseBody {
	s.Index = &v
	return s
}

func (s *InteractTextResponseBody) SetMessage(v string) *InteractTextResponseBody {
	s.Message = &v
	return s
}

func (s *InteractTextResponseBody) SetRelatedImages(v []*string) *InteractTextResponseBody {
	s.RelatedImages = v
	return s
}

func (s *InteractTextResponseBody) SetRelatedVideos(v []*string) *InteractTextResponseBody {
	s.RelatedVideos = v
	return s
}

func (s *InteractTextResponseBody) SetSessionId(v string) *InteractTextResponseBody {
	s.SessionId = &v
	return s
}

func (s *InteractTextResponseBody) SetType(v int32) *InteractTextResponseBody {
	s.Type = &v
	return s
}

func (s *InteractTextResponseBody) Validate() error {
	return dara.Validate(s)
}
