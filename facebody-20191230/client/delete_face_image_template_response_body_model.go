// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceImageTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFaceImageTemplateResponseBody
	GetRequestId() *string
}

type DeleteFaceImageTemplateResponseBody struct {
	// example:
	//
	// F9E76786-3A85-43C3-B79C-3175B4536252
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceImageTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceImageTemplateResponseBody) SetRequestId(v string) *DeleteFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceImageTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
