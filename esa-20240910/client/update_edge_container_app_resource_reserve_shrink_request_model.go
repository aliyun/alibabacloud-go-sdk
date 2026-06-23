// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppResourceReserveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateEdgeContainerAppResourceReserveShrinkRequest
	GetAppId() *string
	SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveShrinkRequest
	GetDurationTime() *string
	SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveShrinkRequest
	GetEnable() *bool
	SetForever(v bool) *UpdateEdgeContainerAppResourceReserveShrinkRequest
	GetForever() *bool
	SetReserveSetShrink(v string) *UpdateEdgeContainerAppResourceReserveShrinkRequest
	GetReserveSetShrink() *string
}

type UpdateEdgeContainerAppResourceReserveShrinkRequest struct {
	// The application ID. You can call the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation to obtain the application ID.
	//
	// 	Notice: The AppId format is the app- prefix followed by a numeric suffix, with a total length of 20 to 64 characters (example: app-8806886***83794688). Call ListEdgeContainerApps to obtain an existing AppId, or call CreateEdgeContainerApp to create an application first.</notice>.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The reservation end time. This parameter uses UTC time. To convert from UTC+8, add 8 hours. For example, if the current time is 2006-01-02 06:04:05 in UTC+8, enter "2006-01-02T14:04:05Z".
	//
	// example:
	//
	// 2006-01-02T15:04:05Z
	DurationTime *string `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	// Specifies whether to enable resource reservation.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Specifies whether to permanently enable reservation. Once enabled, you cannot set a reservation end time.
	//
	// example:
	//
	// true
	Forever *bool `json:"Forever,omitempty" xml:"Forever,omitempty"`
	// The list of reserved resources.
	ReserveSetShrink *string `json:"ReserveSet,omitempty" xml:"ReserveSet,omitempty"`
}

func (s UpdateEdgeContainerAppResourceReserveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) GetDurationTime() *string {
	return s.DurationTime
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) GetForever() *bool {
	return s.Forever
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) GetReserveSetShrink() *string {
	return s.ReserveSetShrink
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) SetAppId(v string) *UpdateEdgeContainerAppResourceReserveShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveShrinkRequest {
	s.DurationTime = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveShrinkRequest {
	s.Enable = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) SetForever(v bool) *UpdateEdgeContainerAppResourceReserveShrinkRequest {
	s.Forever = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) SetReserveSetShrink(v string) *UpdateEdgeContainerAppResourceReserveShrinkRequest {
	s.ReserveSetShrink = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
