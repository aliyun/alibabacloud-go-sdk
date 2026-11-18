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
	// The update content. You can update the information of up to 20 auxiliary media assets at a time. For the parameter structure, refer to the **UpdateContent*	- table below.
	//
	// >- The update content for `Title`, `Description`, and `Tags` cannot contain emojis.
	//
	// >- If a parameter is passed, the corresponding field is updated. Otherwise, the field is not overwritten or updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"MediaId":"bbc65bba53f6ed90de118a7849****","Title":"Title 1","Description":"Description 1","Tags":"Tag 1,Tag 2"},{"MediaId":"f45cf4eba5cb90233389558c39****","Title2":"Title","Description2":"Description","Tags":"Tag 3,Tag 4"}]
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
