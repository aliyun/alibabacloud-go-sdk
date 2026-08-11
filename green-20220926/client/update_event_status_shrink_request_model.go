// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateEventStatusShrinkRequest
	GetAppId() *string
	SetEventIdsShrink(v string) *UpdateEventStatusShrinkRequest
	GetEventIdsShrink() *string
	SetOperationCode(v string) *UpdateEventStatusShrinkRequest
	GetOperationCode() *string
	SetOperationParams(v string) *UpdateEventStatusShrinkRequest
	GetOperationParams() *string
	SetRegionId(v string) *UpdateEventStatusShrinkRequest
	GetRegionId() *string
	SetSource(v string) *UpdateEventStatusShrinkRequest
	GetSource() *string
}

type UpdateEventStatusShrinkRequest struct {
	// The application ID that identifies the application to which the operation belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The list of risk event IDs.
	EventIdsShrink *string `json:"EventIds,omitempty" xml:"EventIds,omitempty"`
	// The operation code that defines the specific type of event status change operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// resolve
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The operation parameters that contain additional parameter information required to execute the operation.
	//
	// example:
	//
	// {}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The operation source that identifies the source system or module that triggered this status update request.
	//
	// example:
	//
	// xx
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateEventStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventStatusShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateEventStatusShrinkRequest) GetEventIdsShrink() *string {
	return s.EventIdsShrink
}

func (s *UpdateEventStatusShrinkRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *UpdateEventStatusShrinkRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *UpdateEventStatusShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStatusShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateEventStatusShrinkRequest) SetAppId(v string) *UpdateEventStatusShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateEventStatusShrinkRequest) SetEventIdsShrink(v string) *UpdateEventStatusShrinkRequest {
	s.EventIdsShrink = &v
	return s
}

func (s *UpdateEventStatusShrinkRequest) SetOperationCode(v string) *UpdateEventStatusShrinkRequest {
	s.OperationCode = &v
	return s
}

func (s *UpdateEventStatusShrinkRequest) SetOperationParams(v string) *UpdateEventStatusShrinkRequest {
	s.OperationParams = &v
	return s
}

func (s *UpdateEventStatusShrinkRequest) SetRegionId(v string) *UpdateEventStatusShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStatusShrinkRequest) SetSource(v string) *UpdateEventStatusShrinkRequest {
	s.Source = &v
	return s
}

func (s *UpdateEventStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
