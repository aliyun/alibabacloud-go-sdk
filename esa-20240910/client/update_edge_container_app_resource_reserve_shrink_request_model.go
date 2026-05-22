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
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the reservation. The input time is UTC. It takes +8 hours to enter Beijing time. For example, if the current time is 2006-01-02 06:04:05, you need to enter "2006-01-02T14:04:05Z".
	//
	// example:
	//
	// 2006-01-02T15:04:05Z
	DurationTime *string `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	// Whether to enable resource reservation.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Whether to permanently enable the reservation. Once it is enabled, you are not allowed to set the reservation deadline.
	//
	// example:
	//
	// true
	Forever *bool `json:"Forever,omitempty" xml:"Forever,omitempty"`
	// Reserved resource list.
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
