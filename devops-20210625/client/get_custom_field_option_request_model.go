// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomFieldOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceIdentifier(v string) *GetCustomFieldOptionRequest
	GetSpaceIdentifier() *string
	SetSpaceType(v string) *GetCustomFieldOptionRequest
	GetSpaceType() *string
	SetWorkitemTypeIdentifier(v string) *GetCustomFieldOptionRequest
	GetWorkitemTypeIdentifier() *string
}

type GetCustomFieldOptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Project
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9uyxxxx1re573f561dxxxxx
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s GetCustomFieldOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldOptionRequest) GoString() string {
	return s.String()
}

func (s *GetCustomFieldOptionRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *GetCustomFieldOptionRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *GetCustomFieldOptionRequest) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *GetCustomFieldOptionRequest) SetSpaceIdentifier(v string) *GetCustomFieldOptionRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetCustomFieldOptionRequest) SetSpaceType(v string) *GetCustomFieldOptionRequest {
	s.SpaceType = &v
	return s
}

func (s *GetCustomFieldOptionRequest) SetWorkitemTypeIdentifier(v string) *GetCustomFieldOptionRequest {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *GetCustomFieldOptionRequest) Validate() error {
	return dara.Validate(s)
}
