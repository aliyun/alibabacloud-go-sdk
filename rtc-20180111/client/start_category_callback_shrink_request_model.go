// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCategoryCallbackShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartCategoryCallbackShrinkRequest
	GetAppId() *string
	SetCallbackShrink(v string) *StartCategoryCallbackShrinkRequest
	GetCallbackShrink() *string
}

type StartCategoryCallbackShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	CallbackShrink *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
}

func (s StartCategoryCallbackShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCategoryCallbackShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartCategoryCallbackShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartCategoryCallbackShrinkRequest) GetCallbackShrink() *string {
	return s.CallbackShrink
}

func (s *StartCategoryCallbackShrinkRequest) SetAppId(v string) *StartCategoryCallbackShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartCategoryCallbackShrinkRequest) SetCallbackShrink(v string) *StartCategoryCallbackShrinkRequest {
	s.CallbackShrink = &v
	return s
}

func (s *StartCategoryCallbackShrinkRequest) Validate() error {
	return dara.Validate(s)
}
