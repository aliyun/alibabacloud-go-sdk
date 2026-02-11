// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataConsistencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *CheckDataConsistencyRequest
	GetAppType() *string
	SetCurrentTimestamp(v int64) *CheckDataConsistencyRequest
	GetCurrentTimestamp() *int64
	SetPid(v string) *CheckDataConsistencyRequest
	GetPid() *string
	SetProxyUserId(v string) *CheckDataConsistencyRequest
	GetProxyUserId() *string
	SetRegionId(v string) *CheckDataConsistencyRequest
	GetRegionId() *string
}

type CheckDataConsistencyRequest struct {
	// This parameter is required.
	AppType          *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentTimestamp *int64  `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	// This parameter is required.
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	ProxyUserId *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckDataConsistencyRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDataConsistencyRequest) GoString() string {
	return s.String()
}

func (s *CheckDataConsistencyRequest) GetAppType() *string {
	return s.AppType
}

func (s *CheckDataConsistencyRequest) GetCurrentTimestamp() *int64 {
	return s.CurrentTimestamp
}

func (s *CheckDataConsistencyRequest) GetPid() *string {
	return s.Pid
}

func (s *CheckDataConsistencyRequest) GetProxyUserId() *string {
	return s.ProxyUserId
}

func (s *CheckDataConsistencyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckDataConsistencyRequest) SetAppType(v string) *CheckDataConsistencyRequest {
	s.AppType = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetCurrentTimestamp(v int64) *CheckDataConsistencyRequest {
	s.CurrentTimestamp = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetPid(v string) *CheckDataConsistencyRequest {
	s.Pid = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetProxyUserId(v string) *CheckDataConsistencyRequest {
	s.ProxyUserId = &v
	return s
}

func (s *CheckDataConsistencyRequest) SetRegionId(v string) *CheckDataConsistencyRequest {
	s.RegionId = &v
	return s
}

func (s *CheckDataConsistencyRequest) Validate() error {
	return dara.Validate(s)
}
