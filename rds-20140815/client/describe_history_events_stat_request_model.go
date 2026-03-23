// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveStatus(v string) *DescribeHistoryEventsStatRequest
	GetArchiveStatus() *string
	SetFromStartTime(v string) *DescribeHistoryEventsStatRequest
	GetFromStartTime() *string
	SetRegionId(v string) *DescribeHistoryEventsStatRequest
	GetRegionId() *string
	SetSecurityToken(v string) *DescribeHistoryEventsStatRequest
	GetSecurityToken() *string
	SetToStartTime(v string) *DescribeHistoryEventsStatRequest
	GetToStartTime() *string
}

type DescribeHistoryEventsStatRequest struct {
	ArchiveStatus *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ToStartTime   *string `json:"ToStartTime,omitempty" xml:"ToStartTime,omitempty"`
}

func (s DescribeHistoryEventsStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsStatRequest) GetArchiveStatus() *string {
	return s.ArchiveStatus
}

func (s *DescribeHistoryEventsStatRequest) GetFromStartTime() *string {
	return s.FromStartTime
}

func (s *DescribeHistoryEventsStatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryEventsStatRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeHistoryEventsStatRequest) GetToStartTime() *string {
	return s.ToStartTime
}

func (s *DescribeHistoryEventsStatRequest) SetArchiveStatus(v string) *DescribeHistoryEventsStatRequest {
	s.ArchiveStatus = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) SetFromStartTime(v string) *DescribeHistoryEventsStatRequest {
	s.FromStartTime = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) SetRegionId(v string) *DescribeHistoryEventsStatRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) SetSecurityToken(v string) *DescribeHistoryEventsStatRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) SetToStartTime(v string) *DescribeHistoryEventsStatRequest {
	s.ToStartTime = &v
	return s
}

func (s *DescribeHistoryEventsStatRequest) Validate() error {
	return dara.Validate(s)
}
