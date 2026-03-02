// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStackInfo interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *StackInfo
	GetApi() *string
	SetDuration(v int64) *StackInfo
	GetDuration() *int64
	SetException(v string) *StackInfo
	GetException() *string
	SetExtInfo(v *StackInfoExtInfo) *StackInfo
	GetExtInfo() *StackInfoExtInfo
	SetLine(v string) *StackInfo
	GetLine() *string
	SetRpcId(v string) *StackInfo
	GetRpcId() *string
	SetServiceName(v string) *StackInfo
	GetServiceName() *string
	SetStartTime(v int64) *StackInfo
	GetStartTime() *int64
}

type StackInfo struct {
	Api         *string           `json:"api,omitempty" xml:"api,omitempty"`
	Duration    *int64            `json:"duration,omitempty" xml:"duration,omitempty"`
	Exception   *string           `json:"exception,omitempty" xml:"exception,omitempty"`
	ExtInfo     *StackInfoExtInfo `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	Line        *string           `json:"line,omitempty" xml:"line,omitempty"`
	RpcId       *string           `json:"rpcId,omitempty" xml:"rpcId,omitempty"`
	ServiceName *string           `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	StartTime   *int64            `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s StackInfo) String() string {
	return dara.Prettify(s)
}

func (s StackInfo) GoString() string {
	return s.String()
}

func (s *StackInfo) GetApi() *string {
	return s.Api
}

func (s *StackInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *StackInfo) GetException() *string {
	return s.Exception
}

func (s *StackInfo) GetExtInfo() *StackInfoExtInfo {
	return s.ExtInfo
}

func (s *StackInfo) GetLine() *string {
	return s.Line
}

func (s *StackInfo) GetRpcId() *string {
	return s.RpcId
}

func (s *StackInfo) GetServiceName() *string {
	return s.ServiceName
}

func (s *StackInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *StackInfo) SetApi(v string) *StackInfo {
	s.Api = &v
	return s
}

func (s *StackInfo) SetDuration(v int64) *StackInfo {
	s.Duration = &v
	return s
}

func (s *StackInfo) SetException(v string) *StackInfo {
	s.Exception = &v
	return s
}

func (s *StackInfo) SetExtInfo(v *StackInfoExtInfo) *StackInfo {
	s.ExtInfo = v
	return s
}

func (s *StackInfo) SetLine(v string) *StackInfo {
	s.Line = &v
	return s
}

func (s *StackInfo) SetRpcId(v string) *StackInfo {
	s.RpcId = &v
	return s
}

func (s *StackInfo) SetServiceName(v string) *StackInfo {
	s.ServiceName = &v
	return s
}

func (s *StackInfo) SetStartTime(v int64) *StackInfo {
	s.StartTime = &v
	return s
}

func (s *StackInfo) Validate() error {
	if s.ExtInfo != nil {
		if err := s.ExtInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
