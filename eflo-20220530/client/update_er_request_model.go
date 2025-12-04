// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateErRequest
	GetDescription() *string
	SetErId(v string) *UpdateErRequest
	GetErId() *string
	SetErName(v string) *UpdateErRequest
	GetErName() *string
	SetRegionId(v string) *UpdateErRequest
	GetRegionId() *string
}

type UpdateErRequest struct {
	// The description of the document.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun HUB Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Parameter
	//
	// example:
	//
	// er-wulanchabu-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateErRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateErRequest) GoString() string {
	return s.String()
}

func (s *UpdateErRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateErRequest) GetErId() *string {
	return s.ErId
}

func (s *UpdateErRequest) GetErName() *string {
	return s.ErName
}

func (s *UpdateErRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateErRequest) SetDescription(v string) *UpdateErRequest {
	s.Description = &v
	return s
}

func (s *UpdateErRequest) SetErId(v string) *UpdateErRequest {
	s.ErId = &v
	return s
}

func (s *UpdateErRequest) SetErName(v string) *UpdateErRequest {
	s.ErName = &v
	return s
}

func (s *UpdateErRequest) SetRegionId(v string) *UpdateErRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateErRequest) Validate() error {
	return dara.Validate(s)
}
