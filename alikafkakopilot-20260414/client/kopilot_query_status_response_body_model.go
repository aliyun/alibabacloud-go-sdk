// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotQueryStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *KopilotQueryStatusResponseBody
	GetCode() *int64
	SetData(v *KopilotQueryStatusResponseBodyData) *KopilotQueryStatusResponseBody
	GetData() *KopilotQueryStatusResponseBodyData
	SetRequestId(v string) *KopilotQueryStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *KopilotQueryStatusResponseBody
	GetSuccess() *bool
}

type KopilotQueryStatusResponseBody struct {
	Code      *int64                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *KopilotQueryStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s KopilotQueryStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KopilotQueryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *KopilotQueryStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *KopilotQueryStatusResponseBody) GetData() *KopilotQueryStatusResponseBodyData {
	return s.Data
}

func (s *KopilotQueryStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KopilotQueryStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *KopilotQueryStatusResponseBody) SetCode(v int64) *KopilotQueryStatusResponseBody {
	s.Code = &v
	return s
}

func (s *KopilotQueryStatusResponseBody) SetData(v *KopilotQueryStatusResponseBodyData) *KopilotQueryStatusResponseBody {
	s.Data = v
	return s
}

func (s *KopilotQueryStatusResponseBody) SetRequestId(v string) *KopilotQueryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *KopilotQueryStatusResponseBody) SetSuccess(v bool) *KopilotQueryStatusResponseBody {
	s.Success = &v
	return s
}

func (s *KopilotQueryStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotQueryStatusResponseBodyData struct {
	ActivateTime *int64  `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LifeStatus   *string `json:"LifeStatus,omitempty" xml:"LifeStatus,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Uid          *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s KopilotQueryStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s KopilotQueryStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *KopilotQueryStatusResponseBodyData) GetActivateTime() *int64 {
	return s.ActivateTime
}

func (s *KopilotQueryStatusResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *KopilotQueryStatusResponseBodyData) GetLifeStatus() *string {
	return s.LifeStatus
}

func (s *KopilotQueryStatusResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotQueryStatusResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *KopilotQueryStatusResponseBodyData) SetActivateTime(v int64) *KopilotQueryStatusResponseBodyData {
	s.ActivateTime = &v
	return s
}

func (s *KopilotQueryStatusResponseBodyData) SetInstanceId(v string) *KopilotQueryStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *KopilotQueryStatusResponseBodyData) SetLifeStatus(v string) *KopilotQueryStatusResponseBodyData {
	s.LifeStatus = &v
	return s
}

func (s *KopilotQueryStatusResponseBodyData) SetRegionId(v string) *KopilotQueryStatusResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *KopilotQueryStatusResponseBodyData) SetUid(v string) *KopilotQueryStatusResponseBodyData {
	s.Uid = &v
	return s
}

func (s *KopilotQueryStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
