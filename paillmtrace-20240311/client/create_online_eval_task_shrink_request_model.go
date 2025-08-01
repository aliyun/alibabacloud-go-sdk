// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineEvalTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateOnlineEvalTaskShrinkRequest
	GetBodyShrink() *string
}

type CreateOnlineEvalTaskShrinkRequest struct {
	// The request data.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOnlineEvalTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineEvalTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateOnlineEvalTaskShrinkRequest) SetBodyShrink(v string) *CreateOnlineEvalTaskShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateOnlineEvalTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
