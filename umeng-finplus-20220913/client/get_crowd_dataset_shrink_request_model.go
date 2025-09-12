// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrowdDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *GetCrowdDatasetShrinkRequest
	GetBodyShrink() *string
}

type GetCrowdDatasetShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrowdDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCrowdDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCrowdDatasetShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *GetCrowdDatasetShrinkRequest) SetBodyShrink(v string) *GetCrowdDatasetShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *GetCrowdDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
