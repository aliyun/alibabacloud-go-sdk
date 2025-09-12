// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYzdInstanceTaskResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *GetYzdInstanceTaskResultShrinkRequest
	GetBodyShrink() *string
}

type GetYzdInstanceTaskResultShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYzdInstanceTaskResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYzdInstanceTaskResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetYzdInstanceTaskResultShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *GetYzdInstanceTaskResultShrinkRequest) SetBodyShrink(v string) *GetYzdInstanceTaskResultShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *GetYzdInstanceTaskResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
