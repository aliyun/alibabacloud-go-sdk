// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProhibitedSoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSoftwareIds(v []*string) *DeleteProhibitedSoftwareRequest
	GetSoftwareIds() []*string
}

type DeleteProhibitedSoftwareRequest struct {
	// The IDs of the prohibited software to delete. Duplicate IDs are not allowed. You can specify up to 100 IDs.
	//
	// This parameter is required.
	SoftwareIds []*string `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
}

func (s DeleteProhibitedSoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProhibitedSoftwareRequest) GoString() string {
	return s.String()
}

func (s *DeleteProhibitedSoftwareRequest) GetSoftwareIds() []*string {
	return s.SoftwareIds
}

func (s *DeleteProhibitedSoftwareRequest) SetSoftwareIds(v []*string) *DeleteProhibitedSoftwareRequest {
	s.SoftwareIds = v
	return s
}

func (s *DeleteProhibitedSoftwareRequest) Validate() error {
	return dara.Validate(s)
}
