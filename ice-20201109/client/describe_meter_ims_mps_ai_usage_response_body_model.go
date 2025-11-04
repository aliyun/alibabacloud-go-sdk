// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsMpsAiUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeMeterImsMpsAiUsageResponseBodyData) *DescribeMeterImsMpsAiUsageResponseBody
	GetData() []*DescribeMeterImsMpsAiUsageResponseBodyData
	SetRequestId(v string) *DescribeMeterImsMpsAiUsageResponseBody
	GetRequestId() *string
}

type DescribeMeterImsMpsAiUsageResponseBody struct {
	// The usage statistics of IMS on AI processing of MPS.
	Data []*DescribeMeterImsMpsAiUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0622C702-41BE-467E-AF2E-883D4517962E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterImsMpsAiUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMpsAiUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMpsAiUsageResponseBody) GetData() []*DescribeMeterImsMpsAiUsageResponseBodyData {
	return s.Data
}

func (s *DescribeMeterImsMpsAiUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMeterImsMpsAiUsageResponseBody) SetData(v []*DescribeMeterImsMpsAiUsageResponseBodyData) *DescribeMeterImsMpsAiUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterImsMpsAiUsageResponseBody) SetRequestId(v string) *DescribeMeterImsMpsAiUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMeterImsMpsAiUsageResponseBodyData struct {
	// The usage duration, in minutes.
	//
	// example:
	//
	// 644
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The beginning time of usage. The value is a 10-digit timestamp.
	//
	// example:
	//
	// 1656950400
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The AI type. Valid values:
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeMeterImsMpsAiUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsMpsAiUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsMpsAiUsageResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeMeterImsMpsAiUsageResponseBodyData) GetTime() *int64 {
	return s.Time
}

func (s *DescribeMeterImsMpsAiUsageResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeMeterImsMpsAiUsageResponseBodyData) SetDuration(v int64) *DescribeMeterImsMpsAiUsageResponseBodyData {
	s.Duration = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageResponseBodyData) SetTime(v int64) *DescribeMeterImsMpsAiUsageResponseBodyData {
	s.Time = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageResponseBodyData) SetType(v string) *DescribeMeterImsMpsAiUsageResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeMeterImsMpsAiUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
