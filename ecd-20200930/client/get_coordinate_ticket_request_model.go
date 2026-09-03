// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoordinateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoId(v string) *GetCoordinateTicketRequest
	GetCoId() *string
	SetEndUserId(v string) *GetCoordinateTicketRequest
	GetEndUserId() *string
	SetRegionId(v string) *GetCoordinateTicketRequest
	GetRegionId() *string
	SetTaskId(v string) *GetCoordinateTicketRequest
	GetTaskId() *string
	SetUserType(v string) *GetCoordinateTicketRequest
	GetUserType() *string
}

type GetCoordinateTicketRequest struct {
	// The coordination flow ID. This value is the `Coid` returned by the [ApplyCoordinationForMonitoring](~~ApplyCoordinationForMonitoring~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// co-0sot77uale3****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// The username of the end user. This parameter is not required on the administrator side.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cloud computer connection task ID. This parameter is not required for the first request. If the first request does not return a Ticket, specify the `TaskId` returned by the first request in subsequent requests.
	//
	// example:
	//
	// 39cc15e5-6998-4b9f-9b2c-7a4cc3e2****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The user type.
	//
	// This parameter is required.
	//
	// example:
	//
	// TENANT_ADMIN
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetCoordinateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCoordinateTicketRequest) GoString() string {
	return s.String()
}

func (s *GetCoordinateTicketRequest) GetCoId() *string {
	return s.CoId
}

func (s *GetCoordinateTicketRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetCoordinateTicketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCoordinateTicketRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCoordinateTicketRequest) GetUserType() *string {
	return s.UserType
}

func (s *GetCoordinateTicketRequest) SetCoId(v string) *GetCoordinateTicketRequest {
	s.CoId = &v
	return s
}

func (s *GetCoordinateTicketRequest) SetEndUserId(v string) *GetCoordinateTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetCoordinateTicketRequest) SetRegionId(v string) *GetCoordinateTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetCoordinateTicketRequest) SetTaskId(v string) *GetCoordinateTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetCoordinateTicketRequest) SetUserType(v string) *GetCoordinateTicketRequest {
	s.UserType = &v
	return s
}

func (s *GetCoordinateTicketRequest) Validate() error {
	return dara.Validate(s)
}
