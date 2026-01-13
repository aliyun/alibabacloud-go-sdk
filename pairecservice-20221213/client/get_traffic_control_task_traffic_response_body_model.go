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
	// example:
	//
	// 6CF1E160-3F36-5E73-A170-C75504F05BBC
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
	if s.TrafficControlTaskTrafficInfo != nil {
		if err := s.TrafficControlTaskTrafficInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.TargetTraffics != nil {
		for _, item := range s.TargetTraffics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics struct {
	Data                   map[string]*TrafficControlTaskTrafficInfoTargetTrafficsDataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	TrafficControlTargetId *string                                                          `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
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

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) GetTrafficControlTargetId() *string {
	return s.TrafficControlTargetId
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) SetData(v map[string]*TrafficControlTaskTrafficInfoTargetTrafficsDataValue) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics {
	s.Data = v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) SetTrafficControlTargetId(v string) *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics {
	s.TrafficControlTargetId = &v
	return s
}

func (s *GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfoTargetTraffics) Validate() error {
	return dara.Validate(s)
}
