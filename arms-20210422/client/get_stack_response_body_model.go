// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackResponseBody
	GetRequestId() *string
	SetStackInfo(v []*GetStackResponseBodyStackInfo) *GetStackResponseBody
	GetStackInfo() []*GetStackResponseBodyStackInfo
}

type GetStackResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackInfo []*GetStackResponseBodyStackInfo `json:"StackInfo,omitempty" xml:"StackInfo,omitempty" type:"Repeated"`
}

func (s GetStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackResponseBody) GetStackInfo() []*GetStackResponseBodyStackInfo {
	return s.StackInfo
}

func (s *GetStackResponseBody) SetRequestId(v string) *GetStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBody) SetStackInfo(v []*GetStackResponseBodyStackInfo) *GetStackResponseBody {
	s.StackInfo = v
	return s
}

func (s *GetStackResponseBody) Validate() error {
	if s.StackInfo != nil {
		for _, item := range s.StackInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStackResponseBodyStackInfo struct {
	Api         *string                               `json:"Api,omitempty" xml:"Api,omitempty"`
	Duration    *int64                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Exception   *string                               `json:"Exception,omitempty" xml:"Exception,omitempty"`
	ExtInfo     *GetStackResponseBodyStackInfoExtInfo `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
	Line        *string                               `json:"Line,omitempty" xml:"Line,omitempty"`
	RpcId       *string                               `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceName *string                               `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStackResponseBodyStackInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyStackInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfo) GetApi() *string {
	return s.Api
}

func (s *GetStackResponseBodyStackInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *GetStackResponseBodyStackInfo) GetException() *string {
	return s.Exception
}

func (s *GetStackResponseBodyStackInfo) GetExtInfo() *GetStackResponseBodyStackInfoExtInfo {
	return s.ExtInfo
}

func (s *GetStackResponseBodyStackInfo) GetLine() *string {
	return s.Line
}

func (s *GetStackResponseBodyStackInfo) GetRpcId() *string {
	return s.RpcId
}

func (s *GetStackResponseBodyStackInfo) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetStackResponseBodyStackInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetStackResponseBodyStackInfo) SetApi(v string) *GetStackResponseBodyStackInfo {
	s.Api = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetDuration(v int64) *GetStackResponseBodyStackInfo {
	s.Duration = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetException(v string) *GetStackResponseBodyStackInfo {
	s.Exception = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetExtInfo(v *GetStackResponseBodyStackInfoExtInfo) *GetStackResponseBodyStackInfo {
	s.ExtInfo = v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetLine(v string) *GetStackResponseBodyStackInfo {
	s.Line = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetRpcId(v string) *GetStackResponseBodyStackInfo {
	s.RpcId = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetServiceName(v string) *GetStackResponseBodyStackInfo {
	s.ServiceName = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetStartTime(v int64) *GetStackResponseBodyStackInfo {
	s.StartTime = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) Validate() error {
	if s.ExtInfo != nil {
		if err := s.ExtInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStackResponseBodyStackInfoExtInfo struct {
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStackResponseBodyStackInfoExtInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyStackInfoExtInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfoExtInfo) GetInfo() *string {
	return s.Info
}

func (s *GetStackResponseBodyStackInfoExtInfo) GetType() *string {
	return s.Type
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetInfo(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Info = &v
	return s
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetType(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Type = &v
	return s
}

func (s *GetStackResponseBodyStackInfoExtInfo) Validate() error {
	return dara.Validate(s)
}
