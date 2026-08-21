// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttachedMediaInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateContent(v string) *UpdateAttachedMediaInfosRequest
	GetUpdateContent() *string
}

type UpdateAttachedMediaInfosRequest struct {
	// The update content. You can update the information of up to 20 auxiliary media assets at a time. For the parameter structure, see the **UpdateContent*	- table below.
	//
	// >- The `Title`, `Description`, and `Tags` fields cannot contain emoticons.
	//
	// >- If a parameter is specified, the corresponding field is updated. Otherwise, the corresponding field is not overwritten or updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"MediaId":"bbc65bba53f6ed90de118a7849****","Title":" title1","Description":" description1","Tags":" tag1, tag2"},{"MediaId":"f45cf4eba5cb90233389558c39****","Title2":" title","Description2":" description","Tags":" tag3, tag4"}]
	UpdateContent *string `json:"UpdateContent,omitempty" xml:"UpdateContent,omitempty"`
}

func (s UpdateAttachedMediaInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttachedMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *UpdateAttachedMediaInfosRequest) GetUpdateContent() *string {
	return s.UpdateContent
}

func (s *UpdateAttachedMediaInfosRequest) SetUpdateContent(v string) *UpdateAttachedMediaInfosRequest {
	s.UpdateContent = &v
	return s
}

func (s *UpdateAttachedMediaInfosRequest) Validate() error {
	return dara.Validate(s)
}
