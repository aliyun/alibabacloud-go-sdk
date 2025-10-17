// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCategoryCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartCategoryCallbackRequest
	GetAppId() *string
	SetCallback(v *StartCategoryCallbackRequestCallback) *StartCategoryCallbackRequest
	GetCallback() *StartCategoryCallbackRequestCallback
}

type StartCategoryCallbackRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Callback *StartCategoryCallbackRequestCallback `json:"Callback,omitempty" xml:"Callback,omitempty" type:"Struct"`
}

func (s StartCategoryCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCategoryCallbackRequest) GoString() string {
	return s.String()
}

func (s *StartCategoryCallbackRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartCategoryCallbackRequest) GetCallback() *StartCategoryCallbackRequestCallback {
	return s.Callback
}

func (s *StartCategoryCallbackRequest) SetAppId(v string) *StartCategoryCallbackRequest {
	s.AppId = &v
	return s
}

func (s *StartCategoryCallbackRequest) SetCallback(v *StartCategoryCallbackRequestCallback) *StartCategoryCallbackRequest {
	s.Callback = v
	return s
}

func (s *StartCategoryCallbackRequest) Validate() error {
	if s.Callback != nil {
		if err := s.Callback.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartCategoryCallbackRequestCallback struct {
	// This parameter is required.
	//
	// example:
	//
	// RecordEvent
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s StartCategoryCallbackRequestCallback) String() string {
	return dara.Prettify(s)
}

func (s StartCategoryCallbackRequestCallback) GoString() string {
	return s.String()
}

func (s *StartCategoryCallbackRequestCallback) GetCategory() *string {
	return s.Category
}

func (s *StartCategoryCallbackRequestCallback) SetCategory(v string) *StartCategoryCallbackRequestCallback {
	s.Category = &v
	return s
}

func (s *StartCategoryCallbackRequestCallback) Validate() error {
	return dara.Validate(s)
}
