// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCategoryCallbackShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopCategoryCallbackShrinkRequest
	GetAppId() *string
	SetCallbackShrink(v string) *StopCategoryCallbackShrinkRequest
	GetCallbackShrink() *string
}

type StopCategoryCallbackShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	CallbackShrink *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
}

func (s StopCategoryCallbackShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCategoryCallbackShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopCategoryCallbackShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopCategoryCallbackShrinkRequest) GetCallbackShrink() *string {
	return s.CallbackShrink
}

func (s *StopCategoryCallbackShrinkRequest) SetAppId(v string) *StopCategoryCallbackShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StopCategoryCallbackShrinkRequest) SetCallbackShrink(v string) *StopCategoryCallbackShrinkRequest {
	s.CallbackShrink = &v
	return s
}

func (s *StopCategoryCallbackShrinkRequest) Validate() error {
	return dara.Validate(s)
}
