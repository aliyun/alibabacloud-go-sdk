// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateEventStatusRequest
	GetAppId() *string
	SetEventIds(v []*string) *UpdateEventStatusRequest
	GetEventIds() []*string
	SetOperationCode(v string) *UpdateEventStatusRequest
	GetOperationCode() *string
	SetOperationParams(v string) *UpdateEventStatusRequest
	GetOperationParams() *string
	SetRegionId(v string) *UpdateEventStatusRequest
	GetRegionId() *string
	SetSource(v string) *UpdateEventStatusRequest
	GetSource() *string
}

type UpdateEventStatusRequest struct {
	// The application ID that identifies the application to which the operation belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The list of risk event IDs.
	EventIds []*string `json:"EventIds,omitempty" xml:"EventIds,omitempty" type:"Repeated"`
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

func (s UpdateEventStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateEventStatusRequest) GetEventIds() []*string {
	return s.EventIds
}

func (s *UpdateEventStatusRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *UpdateEventStatusRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *UpdateEventStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventStatusRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateEventStatusRequest) SetAppId(v string) *UpdateEventStatusRequest {
	s.AppId = &v
	return s
}

func (s *UpdateEventStatusRequest) SetEventIds(v []*string) *UpdateEventStatusRequest {
	s.EventIds = v
	return s
}

func (s *UpdateEventStatusRequest) SetOperationCode(v string) *UpdateEventStatusRequest {
	s.OperationCode = &v
	return s
}

func (s *UpdateEventStatusRequest) SetOperationParams(v string) *UpdateEventStatusRequest {
	s.OperationParams = &v
	return s
}

func (s *UpdateEventStatusRequest) SetRegionId(v string) *UpdateEventStatusRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStatusRequest) SetSource(v string) *UpdateEventStatusRequest {
	s.Source = &v
	return s
}

func (s *UpdateEventStatusRequest) Validate() error {
	return dara.Validate(s)
}
