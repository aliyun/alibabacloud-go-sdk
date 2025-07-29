// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterInspectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisabledCheckItems(v []*string) *GetClusterInspectConfigResponseBody
	GetDisabledCheckItems() []*string
	SetEnabled(v bool) *GetClusterInspectConfigResponseBody
	GetEnabled() *bool
	SetRecurrence(v string) *GetClusterInspectConfigResponseBody
	GetRecurrence() *string
	SetRequestId(v string) *GetClusterInspectConfigResponseBody
	GetRequestId() *string
}

type GetClusterInspectConfigResponseBody struct {
	// The list of disabled inspection items.
	DisabledCheckItems []*string `json:"disabledCheckItems,omitempty" xml:"disabledCheckItems,omitempty" type:"Repeated"`
	// Specifies whether to enable inspection.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The inspection schedule defined through the RFC5545 Recurrence Rule syntax. You must specify BYHOUR and BYMINUTE. Only FREQ=DAILY is supported. COUNT and UNTIL are not supported.
	//
	// example:
	//
	// FREQ=DAILY;BYHOUR=10;BYMINUTE=15
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 49511F2D-D56A-5C24-B9AE-C8491E09B095
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetClusterInspectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterInspectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterInspectConfigResponseBody) GetDisabledCheckItems() []*string {
	return s.DisabledCheckItems
}

func (s *GetClusterInspectConfigResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetClusterInspectConfigResponseBody) GetRecurrence() *string {
	return s.Recurrence
}

func (s *GetClusterInspectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterInspectConfigResponseBody) SetDisabledCheckItems(v []*string) *GetClusterInspectConfigResponseBody {
	s.DisabledCheckItems = v
	return s
}

func (s *GetClusterInspectConfigResponseBody) SetEnabled(v bool) *GetClusterInspectConfigResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetClusterInspectConfigResponseBody) SetRecurrence(v string) *GetClusterInspectConfigResponseBody {
	s.Recurrence = &v
	return s
}

func (s *GetClusterInspectConfigResponseBody) SetRequestId(v string) *GetClusterInspectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterInspectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
