// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *ResetSnapshotRequest
	GetClientId() *string
	SetDesktopId(v string) *ResetSnapshotRequest
	GetDesktopId() *string
	SetLoginToken(v string) *ResetSnapshotRequest
	GetLoginToken() *string
	SetRegionId(v string) *ResetSnapshotRequest
	GetRegionId() *string
	SetSessionId(v string) *ResetSnapshotRequest
	GetSessionId() *string
	SetSnapshotId(v string) *ResetSnapshotRequest
	GetSnapshotId() *string
	SetStopDesktop(v bool) *ResetSnapshotRequest
	GetStopDesktop() *bool
}

type ResetSnapshotRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// b9d8ddfd-65d4-4857-9e97-56477d1f****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-e964cr92klwqb****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1fdef51b727aa91d6c881658978508114d3f5680fa99a66b2a631d17d5bb4860cccf1173be24d77d5ef1423c83aea****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 05182b8c-bb0d-49d3-963c-ee63a507****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-2zeipxmnhej803x7****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// Specifies whether to stop the cloud computer.
	//
	// example:
	//
	// true
	StopDesktop *bool `json:"StopDesktop,omitempty" xml:"StopDesktop,omitempty"`
}

func (s ResetSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ResetSnapshotRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ResetSnapshotRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ResetSnapshotRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ResetSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetSnapshotRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ResetSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ResetSnapshotRequest) GetStopDesktop() *bool {
	return s.StopDesktop
}

func (s *ResetSnapshotRequest) SetClientId(v string) *ResetSnapshotRequest {
	s.ClientId = &v
	return s
}

func (s *ResetSnapshotRequest) SetDesktopId(v string) *ResetSnapshotRequest {
	s.DesktopId = &v
	return s
}

func (s *ResetSnapshotRequest) SetLoginToken(v string) *ResetSnapshotRequest {
	s.LoginToken = &v
	return s
}

func (s *ResetSnapshotRequest) SetRegionId(v string) *ResetSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSessionId(v string) *ResetSnapshotRequest {
	s.SessionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSnapshotId(v string) *ResetSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *ResetSnapshotRequest) SetStopDesktop(v bool) *ResetSnapshotRequest {
	s.StopDesktop = &v
	return s
}

func (s *ResetSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
