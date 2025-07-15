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
	// The ID of the stream collaboration. You can obtain the value of this parameter based on the value of `Coid` that is returned by the `ApplyCoordinationForMonitoring` operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// co-0sot77uale3****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// The name of the convenience user account. If you initiate the request as an administrator, you do not need to specify this parameter.
	//
	// example:
	//
	// Alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/436773.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud computer connection task. The first time you initiate the request, you do not need to specify the ID of the cloud computer connection task. If no ticket is returned after you initiate the first request, you must specify the value of taskId that is returned for the first request in the subsequent request.
	//
	// example:
	//
	// 39cc15e5-6998-4b9f-9b2c-7a4cc3e2****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the user.
	//
	// Set the value to TENANT_ADMIN.
	//
	// 	- The value of
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     TENANT_ADMIN
	//
	//     <!-- -->
	//
	//     specifies an administrator.
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
