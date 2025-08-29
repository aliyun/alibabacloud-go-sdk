// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTaskTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTrafficControlTaskTrafficResponseBody
	GetRequestId() *string
	SetTrafficControlTaskTrafficInfo(v *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) *GetTrafficControlTaskTrafficResponseBody
	GetTrafficControlTaskTrafficInfo() *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo
}

type GetTrafficControlTaskTrafficResponseBody struct {
	RequestId                     *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlTaskTrafficInfo *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo `json:"TrafficControlTaskTrafficInfo,omitempty" xml:"TrafficControlTaskTrafficInfo,omitempty" type:"Struct"`
}

func (s GetTrafficControlTaskTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrafficControlTaskTrafficResponseBody) GetTrafficControlTaskTrafficInfo() *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo {
	return s.TrafficControlTaskTrafficInfo
}

func (s *GetTrafficControlTaskTrafficResponseBody) SetRequestId(v string) *GetTrafficControlTaskTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBody) SetTrafficControlTaskTrafficInfo(v *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) *GetTrafficControlTaskTrafficResponseBody {
	s.TrafficControlTaskTrafficInfo = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo struct {
	TargetTraffics []*GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics `json:"TargetTraffics,omitempty" xml:"TargetTraffics,omitempty" type:"Repeated"`
	TaskTraffics   map[string]*TrafficControlTaskTrafficInfoTaskTrafficsValue                             `json:"TaskTraffics,omitempty" xml:"TaskTraffics,omitempty"`
}

func (s GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) GetTargetTraffics() []*GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics {
	return s.TargetTraffics
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) GetTaskTraffics() map[string]*TrafficControlTaskTrafficInfoTaskTrafficsValue {
	return s.TaskTraffics
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) SetTargetTraffics(v []*GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo {
	s.TargetTraffics = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) SetTaskTraffics(v map[string]*TrafficControlTaskTrafficInfoTaskTrafficsValue) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo {
	s.TaskTraffics = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) Validate() error {
	return dara.Validate(s)
}

type GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics struct {
	Data                   map[string]*TrafficControlTaskTrafficInfoTargetTrafficsDataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	TrafficContorlTargetId *string                                                          `json:"TrafficContorlTargetId,omitempty" xml:"TrafficContorlTargetId,omitempty"`
}

func (s GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) GetData() map[string]*TrafficControlTaskTrafficInfoTargetTrafficsDataValue {
	return s.Data
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) GetTrafficContorlTargetId() *string {
	return s.TrafficContorlTargetId
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) SetData(v map[string]*TrafficControlTaskTrafficInfoTargetTrafficsDataValue) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics {
	s.Data = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) SetTrafficContorlTargetId(v string) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics {
	s.TrafficContorlTargetId = &v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) Validate() error {
	return dara.Validate(s)
}
