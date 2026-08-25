// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserProvisioningStatisticsResponseBody
	GetRequestId() *string
	SetUserProvisioningStatistics(v *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) *GetUserProvisioningStatisticsResponseBody
	GetUserProvisioningStatistics() *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics
}

type GetUserProvisioningStatisticsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of the RAM user provisioning.
	UserProvisioningStatistics *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics `json:"UserProvisioningStatistics,omitempty" xml:"UserProvisioningStatistics,omitempty" type:"Struct"`
}

func (s GetUserProvisioningStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserProvisioningStatisticsResponseBody) GetUserProvisioningStatistics() *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	return s.UserProvisioningStatistics
}

func (s *GetUserProvisioningStatisticsResponseBody) SetRequestId(v string) *GetUserProvisioningStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBody) SetUserProvisioningStatistics(v *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) *GetUserProvisioningStatisticsResponseBody {
	s.UserProvisioningStatistics = v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBody) Validate() error {
	if s.UserProvisioningStatistics != nil {
		if err := s.UserProvisioningStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The entity ID, which is the ID of the RAM user provisioning.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The number of failed RAM user provisioning events that are associated with the RAM user provisioning.
	//
	// example:
	//
	// 3
	FailedEventCount *int64 `json:"FailedEventCount,omitempty" xml:"FailedEventCount,omitempty"`
	// The time when the RAM user provisioning was last performed.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	LatestAsyncTime *string `json:"LatestAsyncTime,omitempty" xml:"LatestAsyncTime,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource directory belongs.
	//
	// example:
	//
	// 139665787317****
	OwnerPk *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	// The entity type. The value is fixed as `User Provisioning`.
	//
	// example:
	//
	// User Provisioning
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) GetEntityId() *string {
	return s.EntityId
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) GetFailedEventCount() *int64 {
	return s.FailedEventCount
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) GetLatestAsyncTime() *string {
	return s.LatestAsyncTime
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) GetOwnerPk() *string {
	return s.OwnerPk
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) GetType() *string {
	return s.Type
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetDirectoryId(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetEntityId(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.EntityId = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetFailedEventCount(v int64) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.FailedEventCount = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetLatestAsyncTime(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.LatestAsyncTime = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetOwnerPk(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.OwnerPk = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) SetType(v string) *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics {
	s.Type = &v
	return s
}

func (s *GetUserProvisioningStatisticsResponseBodyUserProvisioningStatistics) Validate() error {
	return dara.Validate(s)
}
