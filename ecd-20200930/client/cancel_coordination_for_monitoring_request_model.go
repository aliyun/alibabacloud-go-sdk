// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCoordinationForMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoIds(v []*string) *CancelCoordinationForMonitoringRequest
	GetCoIds() []*string
	SetEndUserId(v string) *CancelCoordinationForMonitoringRequest
	GetEndUserId() *string
	SetRegionId(v string) *CancelCoordinationForMonitoringRequest
	GetRegionId() *string
	SetUserType(v string) *CancelCoordinationForMonitoringRequest
	GetUserType() *string
}

type CancelCoordinationForMonitoringRequest struct {
	// The IDs of stream collaboration tasks.
	//
	// This parameter is required.
	CoIds []*string `json:"CoIds,omitempty" xml:"CoIds,omitempty" type:"Repeated"`
	// The ID of the end user that initiates stream collaboration. If the initiator is the administrator, skip this parameter.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/436773.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the user.
	//
	// Valid value:
	//
	// 	- TENANT_ADMIN: administrator.
	//
	// example:
	//
	// TENANT_ADMIN
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CancelCoordinationForMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCoordinationForMonitoringRequest) GoString() string {
	return s.String()
}

func (s *CancelCoordinationForMonitoringRequest) GetCoIds() []*string {
	return s.CoIds
}

func (s *CancelCoordinationForMonitoringRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CancelCoordinationForMonitoringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelCoordinationForMonitoringRequest) GetUserType() *string {
	return s.UserType
}

func (s *CancelCoordinationForMonitoringRequest) SetCoIds(v []*string) *CancelCoordinationForMonitoringRequest {
	s.CoIds = v
	return s
}

func (s *CancelCoordinationForMonitoringRequest) SetEndUserId(v string) *CancelCoordinationForMonitoringRequest {
	s.EndUserId = &v
	return s
}

func (s *CancelCoordinationForMonitoringRequest) SetRegionId(v string) *CancelCoordinationForMonitoringRequest {
	s.RegionId = &v
	return s
}

func (s *CancelCoordinationForMonitoringRequest) SetUserType(v string) *CancelCoordinationForMonitoringRequest {
	s.UserType = &v
	return s
}

func (s *CancelCoordinationForMonitoringRequest) Validate() error {
	return dara.Validate(s)
}
