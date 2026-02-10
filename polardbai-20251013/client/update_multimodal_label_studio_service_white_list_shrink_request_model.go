// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultimodalLabelStudioServiceWhiteListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest
	GetDBClusterId() *string
	SetWhiteListShrink(v string) *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest
	GetWhiteListShrink() *string
}

type UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	WhiteListShrink *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest) GetWhiteListShrink() *string {
	return s.WhiteListShrink
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest) SetDBClusterId(v string) *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest) SetWhiteListShrink(v string) *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest {
	s.WhiteListShrink = &v
	return s
}

func (s *UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
