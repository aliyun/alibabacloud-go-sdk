// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDescriptionRequest
	GetDescription() *string
	SetClientToken(v string) *UpdateDescriptionRequest
	GetClientToken() *string
}

type UpdateDescriptionRequest struct {
	// example:
	//
	// aliyunes_name_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The new name of the instance.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B350****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateDescriptionRequest) SetDescription(v string) *UpdateDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateDescriptionRequest) SetClientToken(v string) *UpdateDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
