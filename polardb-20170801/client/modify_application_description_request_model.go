// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationDescriptionRequest
	GetApplicationId() *string
	SetDescription(v string) *ModifyApplicationDescriptionRequest
	GetDescription() *string
}

type ModifyApplicationDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my app
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyApplicationDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationDescriptionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyApplicationDescriptionRequest) SetApplicationId(v string) *ModifyApplicationDescriptionRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationDescriptionRequest) SetDescription(v string) *ModifyApplicationDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyApplicationDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
