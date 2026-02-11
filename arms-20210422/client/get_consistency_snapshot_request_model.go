// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsistencySnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetConsistencySnapshotRequest
	GetAppType() *string
	SetCurrentTimestamp(v int64) *GetConsistencySnapshotRequest
	GetCurrentTimestamp() *int64
	SetPid(v string) *GetConsistencySnapshotRequest
	GetPid() *string
	SetProxyUserId(v string) *GetConsistencySnapshotRequest
	GetProxyUserId() *string
	SetRegionId(v string) *GetConsistencySnapshotRequest
	GetRegionId() *string
}

type GetConsistencySnapshotRequest struct {
	// This parameter is required.
	AppType          *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentTimestamp *int64  `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	// This parameter is required.
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsistencySnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsistencySnapshotRequest) GoString() string {
	return s.String()
}

func (s *GetConsistencySnapshotRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetConsistencySnapshotRequest) GetCurrentTimestamp() *int64 {
	return s.CurrentTimestamp
}

func (s *GetConsistencySnapshotRequest) GetPid() *string {
	return s.Pid
}

func (s *GetConsistencySnapshotRequest) GetProxyUserId() *string {
	return s.ProxyUserId
}

func (s *GetConsistencySnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConsistencySnapshotRequest) SetAppType(v string) *GetConsistencySnapshotRequest {
	s.AppType = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetCurrentTimestamp(v int64) *GetConsistencySnapshotRequest {
	s.CurrentTimestamp = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetPid(v string) *GetConsistencySnapshotRequest {
	s.Pid = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetProxyUserId(v string) *GetConsistencySnapshotRequest {
	s.ProxyUserId = &v
	return s
}

func (s *GetConsistencySnapshotRequest) SetRegionId(v string) *GetConsistencySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *GetConsistencySnapshotRequest) Validate() error {
	return dara.Validate(s)
}
