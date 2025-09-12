// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCrowdDatasetAndBindingDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *SaveCrowdDatasetAndBindingDatasetShrinkRequest
	GetBodyShrink() *string
}

type SaveCrowdDatasetAndBindingDatasetShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveCrowdDatasetAndBindingDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveCrowdDatasetAndBindingDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveCrowdDatasetAndBindingDatasetShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *SaveCrowdDatasetAndBindingDatasetShrinkRequest) SetBodyShrink(v string) *SaveCrowdDatasetAndBindingDatasetShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *SaveCrowdDatasetAndBindingDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
