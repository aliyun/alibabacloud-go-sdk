// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeyDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateKeyDescriptionRequest
	GetDescription() *string
	SetKeyId(v string) *UpdateKeyDescriptionRequest
	GetKeyId() *string
}

type UpdateKeyDescriptionRequest struct {
	// The description of the CMK. This description includes the purpose of the CMK, such as the types of data that you want to protect and applications that can use the CMK.
	//
	// This parameter is required.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s UpdateKeyDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeyDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKeyDescriptionRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *UpdateKeyDescriptionRequest) SetDescription(v string) *UpdateKeyDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateKeyDescriptionRequest) SetKeyId(v string) *UpdateKeyDescriptionRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateKeyDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
