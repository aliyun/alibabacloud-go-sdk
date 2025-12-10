// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewMCTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v [][]*string) *PreviewMCTableResponseBody
	GetContent() [][]*string
	SetRequestId(v string) *PreviewMCTableResponseBody
	GetRequestId() *string
}

type PreviewMCTableResponseBody struct {
	Content [][]*string `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreviewMCTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewMCTableResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewMCTableResponseBody) GetContent() [][]*string {
	return s.Content
}

func (s *PreviewMCTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewMCTableResponseBody) SetContent(v [][]*string) *PreviewMCTableResponseBody {
	s.Content = v
	return s
}

func (s *PreviewMCTableResponseBody) SetRequestId(v string) *PreviewMCTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewMCTableResponseBody) Validate() error {
	return dara.Validate(s)
}
