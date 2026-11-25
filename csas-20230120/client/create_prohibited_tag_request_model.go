// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProhibitedTagRequest
	GetDescription() *string
	SetName(v string) *CreateProhibitedTagRequest
	GetName() *string
}

type CreateProhibitedTagRequest struct {
	// The description of the prohibited software tag. The description can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase English letters, digits, spaces, periods (.), underscores (_), and hyphens (-). This parameter can be left empty.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the prohibited software tag. Fuzzy match is supported. The name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase English letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest_616bcc13
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateProhibitedTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedTagRequest) GoString() string {
	return s.String()
}

func (s *CreateProhibitedTagRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProhibitedTagRequest) GetName() *string {
	return s.Name
}

func (s *CreateProhibitedTagRequest) SetDescription(v string) *CreateProhibitedTagRequest {
	s.Description = &v
	return s
}

func (s *CreateProhibitedTagRequest) SetName(v string) *CreateProhibitedTagRequest {
	s.Name = &v
	return s
}

func (s *CreateProhibitedTagRequest) Validate() error {
	return dara.Validate(s)
}
