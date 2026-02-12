// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsMessageTraceResponseBodyData) *OnsMessageTraceResponseBody
	GetData() *OnsMessageTraceResponseBodyData
	SetRequestId(v string) *OnsMessageTraceResponseBody
	GetRequestId() *string
}

type OnsMessageTraceResponseBody struct {
	Data *OnsMessageTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// EAE5BE23-37A1-4354-94D6-E44AE17E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessageTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageTraceResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponseBody) GetData() *OnsMessageTraceResponseBodyData {
	return s.Data
}

func (s *OnsMessageTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsMessageTraceResponseBody) SetData(v *OnsMessageTraceResponseBodyData) *OnsMessageTraceResponseBody {
	s.Data = v
	return s
}

func (s *OnsMessageTraceResponseBody) SetRequestId(v string) *OnsMessageTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessageTraceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessageTraceResponseBodyData struct {
	MessageTrack []*OnsMessageTraceResponseBodyDataMessageTrack `json:"MessageTrack,omitempty" xml:"MessageTrack,omitempty" type:"Repeated"`
}

func (s OnsMessageTraceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponseBodyData) GetMessageTrack() []*OnsMessageTraceResponseBodyDataMessageTrack {
	return s.MessageTrack
}

func (s *OnsMessageTraceResponseBodyData) SetMessageTrack(v []*OnsMessageTraceResponseBodyDataMessageTrack) *OnsMessageTraceResponseBodyData {
	s.MessageTrack = v
	return s
}

func (s *OnsMessageTraceResponseBodyData) Validate() error {
	if s.MessageTrack != nil {
		for _, item := range s.MessageTrack {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsMessageTraceResponseBodyDataMessageTrack struct {
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TrackType     *string `json:"TrackType,omitempty" xml:"TrackType,omitempty"`
}

func (s OnsMessageTraceResponseBodyDataMessageTrack) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageTraceResponseBodyDataMessageTrack) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) GetTrackType() *string {
	return s.TrackType
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) SetConsumerGroup(v string) *OnsMessageTraceResponseBodyDataMessageTrack {
	s.ConsumerGroup = &v
	return s
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) SetInstanceId(v string) *OnsMessageTraceResponseBodyDataMessageTrack {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) SetTrackType(v string) *OnsMessageTraceResponseBodyDataMessageTrack {
	s.TrackType = &v
	return s
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) Validate() error {
	return dara.Validate(s)
}
