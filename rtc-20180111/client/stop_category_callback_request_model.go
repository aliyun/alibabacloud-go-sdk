// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCategoryCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopCategoryCallbackRequest
	GetAppId() *string
	SetCallback(v *StopCategoryCallbackRequestCallback) *StopCategoryCallbackRequest
	GetCallback() *StopCategoryCallbackRequestCallback
}

type StopCategoryCallbackRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Callback *StopCategoryCallbackRequestCallback `json:"Callback,omitempty" xml:"Callback,omitempty" type:"Struct"`
}

func (s StopCategoryCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCategoryCallbackRequest) GoString() string {
	return s.String()
}

func (s *StopCategoryCallbackRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopCategoryCallbackRequest) GetCallback() *StopCategoryCallbackRequestCallback {
	return s.Callback
}

func (s *StopCategoryCallbackRequest) SetAppId(v string) *StopCategoryCallbackRequest {
	s.AppId = &v
	return s
}

func (s *StopCategoryCallbackRequest) SetCallback(v *StopCategoryCallbackRequestCallback) *StopCategoryCallbackRequest {
	s.Callback = v
	return s
}

func (s *StopCategoryCallbackRequest) Validate() error {
	if s.Callback != nil {
		if err := s.Callback.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopCategoryCallbackRequestCallback struct {
	// This parameter is required.
	//
	// example:
	//
	// RecordEvent
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s StopCategoryCallbackRequestCallback) String() string {
	return dara.Prettify(s)
}

func (s StopCategoryCallbackRequestCallback) GoString() string {
	return s.String()
}

func (s *StopCategoryCallbackRequestCallback) GetCategory() *string {
	return s.Category
}

func (s *StopCategoryCallbackRequestCallback) SetCategory(v string) *StopCategoryCallbackRequestCallback {
	s.Category = &v
	return s
}

func (s *StopCategoryCallbackRequestCallback) Validate() error {
	return dara.Validate(s)
}
