// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProhibitedSoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSoftwareId(v *GetProhibitedSoftwareRequestSoftwareId) *GetProhibitedSoftwareRequest
	GetSoftwareId() *GetProhibitedSoftwareRequestSoftwareId
}

type GetProhibitedSoftwareRequest struct {
	// The prohibited software ID.
	SoftwareId *GetProhibitedSoftwareRequestSoftwareId `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty" type:"Struct"`
}

func (s GetProhibitedSoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareRequest) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareRequest) GetSoftwareId() *GetProhibitedSoftwareRequestSoftwareId {
	return s.SoftwareId
}

func (s *GetProhibitedSoftwareRequest) SetSoftwareId(v *GetProhibitedSoftwareRequestSoftwareId) *GetProhibitedSoftwareRequest {
	s.SoftwareId = v
	return s
}

func (s *GetProhibitedSoftwareRequest) Validate() error {
	if s.SoftwareId != nil {
		if err := s.SoftwareId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProhibitedSoftwareRequestSoftwareId struct {
	// Indicates whether the prohibited software is a system built-in prohibited software. Valid values:
	//
	// - **true**: A system built-in prohibited software that is shared across all Alibaba Cloud accounts and cannot be modified or deleted.
	//
	// - **false**: A custom prohibited software under the current Alibaba Cloud account.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The prohibited software ID. You can obtain the value from the following operations:
	//
	// - [ListProhibitedSoftware](~~ListProhibitedSoftware~~): Lists prohibited software.
	//
	// - [CreateProhibitedSoftware](~~CreateProhibitedSoftware~~): Creates custom prohibited software.
	//
	// example:
	//
	// swb-3e6a1f9c4b28****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s GetProhibitedSoftwareRequestSoftwareId) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareRequestSoftwareId) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareRequestSoftwareId) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetProhibitedSoftwareRequestSoftwareId) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *GetProhibitedSoftwareRequestSoftwareId) SetIsDefault(v bool) *GetProhibitedSoftwareRequestSoftwareId {
	s.IsDefault = &v
	return s
}

func (s *GetProhibitedSoftwareRequestSoftwareId) SetSoftwareId(v string) *GetProhibitedSoftwareRequestSoftwareId {
	s.SoftwareId = &v
	return s
}

func (s *GetProhibitedSoftwareRequestSoftwareId) Validate() error {
	return dara.Validate(s)
}
