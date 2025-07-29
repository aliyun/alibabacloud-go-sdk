// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterInspectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisabledCheckItems(v []*string) *CreateClusterInspectConfigRequest
	GetDisabledCheckItems() []*string
	SetEnabled(v bool) *CreateClusterInspectConfigRequest
	GetEnabled() *bool
	SetRecurrence(v string) *CreateClusterInspectConfigRequest
	GetRecurrence() *string
}

type CreateClusterInspectConfigRequest struct {
	// The list of disabled inspection items.
	DisabledCheckItems []*string `json:"disabledCheckItems,omitempty" xml:"disabledCheckItems,omitempty" type:"Repeated"`
	// Specifies whether to enable cluster inspection.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The inspection period defined using RFC5545 Recurrence Rule. You must specify BYHOUR and BYMINUTE. Only FREQ=DAILY is supported. COUNT or UNTIL is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// FREQ=DAILY;BYHOUR=10;BYMINUTE=15
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
}

func (s CreateClusterInspectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterInspectConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterInspectConfigRequest) GetDisabledCheckItems() []*string {
	return s.DisabledCheckItems
}

func (s *CreateClusterInspectConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateClusterInspectConfigRequest) GetRecurrence() *string {
	return s.Recurrence
}

func (s *CreateClusterInspectConfigRequest) SetDisabledCheckItems(v []*string) *CreateClusterInspectConfigRequest {
	s.DisabledCheckItems = v
	return s
}

func (s *CreateClusterInspectConfigRequest) SetEnabled(v bool) *CreateClusterInspectConfigRequest {
	s.Enabled = &v
	return s
}

func (s *CreateClusterInspectConfigRequest) SetRecurrence(v string) *CreateClusterInspectConfigRequest {
	s.Recurrence = &v
	return s
}

func (s *CreateClusterInspectConfigRequest) Validate() error {
	return dara.Validate(s)
}
