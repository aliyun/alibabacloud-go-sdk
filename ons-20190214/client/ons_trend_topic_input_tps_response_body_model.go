// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTrendTopicInputTpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsTrendTopicInputTpsResponseBodyData) *OnsTrendTopicInputTpsResponseBody
	GetData() *OnsTrendTopicInputTpsResponseBodyData
	SetRequestId(v string) *OnsTrendTopicInputTpsResponseBody
	GetRequestId() *string
}

type OnsTrendTopicInputTpsResponseBody struct {
	// The data returned.
	Data *OnsTrendTopicInputTpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// E213AD8A-0730-4B3D-A35A-340DA47D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTrendTopicInputTpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBody) GetData() *OnsTrendTopicInputTpsResponseBodyData {
	return s.Data
}

func (s *OnsTrendTopicInputTpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTrendTopicInputTpsResponseBody) SetData(v *OnsTrendTopicInputTpsResponseBodyData) *OnsTrendTopicInputTpsResponseBody {
	s.Data = v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBody) SetRequestId(v string) *OnsTrendTopicInputTpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTrendTopicInputTpsResponseBodyData struct {
	Records *OnsTrendTopicInputTpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// The name of the table.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx%test trend of received messages
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The unit of the timestamp.
	//
	// example:
	//
	// time
	XUnit *string `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	// The unit of the Y axis.
	//
	// example:
	//
	// msg
	YUnit *string `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
}

func (s OnsTrendTopicInputTpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBodyData) GetRecords() *OnsTrendTopicInputTpsResponseBodyDataRecords {
	return s.Records
}

func (s *OnsTrendTopicInputTpsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *OnsTrendTopicInputTpsResponseBodyData) GetXUnit() *string {
	return s.XUnit
}

func (s *OnsTrendTopicInputTpsResponseBodyData) GetYUnit() *string {
	return s.YUnit
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetRecords(v *OnsTrendTopicInputTpsResponseBodyDataRecords) *OnsTrendTopicInputTpsResponseBodyData {
	s.Records = v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetTitle(v string) *OnsTrendTopicInputTpsResponseBodyData {
	s.Title = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetXUnit(v string) *OnsTrendTopicInputTpsResponseBodyData {
	s.XUnit = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetYUnit(v string) *OnsTrendTopicInputTpsResponseBodyData {
	s.YUnit = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyData) Validate() error {
	if s.Records != nil {
		if err := s.Records.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTrendTopicInputTpsResponseBodyDataRecords struct {
	StatsDataDo []*OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo `json:"StatsDataDo,omitempty" xml:"StatsDataDo,omitempty" type:"Repeated"`
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecords) GetStatsDataDo() []*OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo {
	return s.StatsDataDo
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecords) SetStatsDataDo(v []*OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) *OnsTrendTopicInputTpsResponseBodyDataRecords {
	s.StatsDataDo = v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecords) Validate() error {
	if s.StatsDataDo != nil {
		for _, item := range s.StatsDataDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo struct {
	X *int64   `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) GetX() *int64 {
	return s.X
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) GetY() *float32 {
	return s.Y
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) SetX(v int64) *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo {
	s.X = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) SetY(v float32) *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo {
	s.Y = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) Validate() error {
	return dara.Validate(s)
}
