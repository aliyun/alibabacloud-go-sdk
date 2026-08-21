// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateContent(v string) *UpdateImageInfosRequest
	GetUpdateContent() *string
}

type UpdateImageInfosRequest struct {
	// The update content. You can modify the information of up to 20 images at a time. For the parameter structure, see the **UpdateContent*	- table below.
	//
	// > - The Title, Description, and Tags fields cannot contain emoticons.
	//
	// >- If a parameter is specified, the corresponding field is updated. Otherwise, the corresponding field is not overwritten or updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ImageId":"ff8fe57e3461416c6a267a4e09****","Title":" title","Description":" description","Tags":" tag1, tag2"}]
	UpdateContent *string `json:"UpdateContent,omitempty" xml:"UpdateContent,omitempty"`
}

func (s UpdateImageInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageInfosRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageInfosRequest) GetUpdateContent() *string {
	return s.UpdateContent
}

func (s *UpdateImageInfosRequest) SetUpdateContent(v string) *UpdateImageInfosRequest {
	s.UpdateContent = &v
	return s
}

func (s *UpdateImageInfosRequest) Validate() error {
	return dara.Validate(s)
}
