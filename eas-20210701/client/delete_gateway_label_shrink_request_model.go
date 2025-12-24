// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayLabelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelKeysShrink(v string) *DeleteGatewayLabelShrinkRequest
	GetLabelKeysShrink() *string
}

type DeleteGatewayLabelShrinkRequest struct {
	// This parameter is required.
	LabelKeysShrink *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s DeleteGatewayLabelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayLabelShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayLabelShrinkRequest) GetLabelKeysShrink() *string {
	return s.LabelKeysShrink
}

func (s *DeleteGatewayLabelShrinkRequest) SetLabelKeysShrink(v string) *DeleteGatewayLabelShrinkRequest {
	s.LabelKeysShrink = &v
	return s
}

func (s *DeleteGatewayLabelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
