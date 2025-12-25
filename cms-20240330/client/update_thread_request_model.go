// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateThreadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *UpdateThreadRequest
	GetStatus() *string
	SetTitle(v string) *UpdateThreadRequest
	GetTitle() *string
}

type UpdateThreadRequest struct {
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// MY ANTOM Tech Team Annual Dinner Performance Lucky Draw
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateThreadRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateThreadRequest) GoString() string {
	return s.String()
}

func (s *UpdateThreadRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateThreadRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateThreadRequest) SetStatus(v string) *UpdateThreadRequest {
	s.Status = &v
	return s
}

func (s *UpdateThreadRequest) SetTitle(v string) *UpdateThreadRequest {
	s.Title = &v
	return s
}

func (s *UpdateThreadRequest) Validate() error {
	return dara.Validate(s)
}
