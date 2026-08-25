// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningRdAccountStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserProvisioningRdAccountStatisticsResponseBody
	GetRequestId() *string
	SetUserProvisioningStatistics(v *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) *GetUserProvisioningRdAccountStatisticsResponseBody
	GetUserProvisioningStatistics() *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics
}

type GetUserProvisioningRdAccountStatisticsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of the RAM user provisioning.
	UserProvisioningStatistics *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics `json:"UserProvisioningStatistics,omitempty" xml:"UserProvisioningStatistics,omitempty" type:"Struct"`
}

func (s GetUserProvisioningRdAccountStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningRdAccountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBody) GetUserProvisioningStatistics() *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	return s.UserProvisioningStatistics
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBody) SetRequestId(v string) *GetUserProvisioningRdAccountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBody) SetUserProvisioningStatistics(v *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) *GetUserProvisioningRdAccountStatisticsResponseBody {
	s.UserProvisioningStatistics = v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBody) Validate() error {
	if s.UserProvisioningStatistics != nil {
		if err := s.UserProvisioningStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The entity ID, which is the ID of the member in the resource directory.
	//
	// example:
	//
	// 1743382******
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The number of failed RAM user provisioning events.
	//
	// example:
	//
	// 4
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
	// 1639738******
	OwnerPk *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	// The entity type. The value is fixed as `RD Account`.
	//
	// example:
	//
	// RD Account
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) GetEntityId() *string {
	return s.EntityId
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) GetFailedEventCount() *int64 {
	return s.FailedEventCount
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) GetLatestAsyncTime() *string {
	return s.LatestAsyncTime
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) GetOwnerPk() *string {
	return s.OwnerPk
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) GetType() *string {
	return s.Type
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetDirectoryId(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetEntityId(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.EntityId = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetFailedEventCount(v int64) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.FailedEventCount = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetLatestAsyncTime(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.LatestAsyncTime = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetOwnerPk(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.OwnerPk = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) SetType(v string) *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics {
	s.Type = &v
	return s
}

func (s *GetUserProvisioningRdAccountStatisticsResponseBodyUserProvisioningStatistics) Validate() error {
	return dara.Validate(s)
}
