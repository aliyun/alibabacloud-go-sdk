// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProhibitedSoftwareShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSoftwareIdShrink(v string) *GetProhibitedSoftwareShrinkRequest
	GetSoftwareIdShrink() *string
}

type GetProhibitedSoftwareShrinkRequest struct {
	// The prohibited software ID.
	SoftwareIdShrink *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s GetProhibitedSoftwareShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareShrinkRequest) GetSoftwareIdShrink() *string {
	return s.SoftwareIdShrink
}

func (s *GetProhibitedSoftwareShrinkRequest) SetSoftwareIdShrink(v string) *GetProhibitedSoftwareShrinkRequest {
	s.SoftwareIdShrink = &v
	return s
}

func (s *GetProhibitedSoftwareShrinkRequest) Validate() error {
	return dara.Validate(s)
}
