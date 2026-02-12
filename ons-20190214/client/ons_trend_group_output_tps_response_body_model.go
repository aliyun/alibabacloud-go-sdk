// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTrendGroupOutputTpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsTrendGroupOutputTpsResponseBodyData) *OnsTrendGroupOutputTpsResponseBody
	GetData() *OnsTrendGroupOutputTpsResponseBodyData
	SetRequestId(v string) *OnsTrendGroupOutputTpsResponseBody
	GetRequestId() *string
}

type OnsTrendGroupOutputTpsResponseBody struct {
	// The data returned.
	Data *OnsTrendGroupOutputTpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// CE57AEDC-8FD2-43ED-8E3B-1F878077****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBody) GetData() *OnsTrendGroupOutputTpsResponseBodyData {
	return s.Data
}

func (s *OnsTrendGroupOutputTpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTrendGroupOutputTpsResponseBody) SetData(v *OnsTrendGroupOutputTpsResponseBodyData) *OnsTrendGroupOutputTpsResponseBody {
	s.Data = v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBody) SetRequestId(v string) *OnsTrendGroupOutputTpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTrendGroupOutputTpsResponseBodyData struct {
	Records *OnsTrendGroupOutputTpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// The name of the table.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx%test@MQ_INST_111111111111_DOxxxxxx%GID_test trend chart of delivered messages
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The unit of the timestamp.
	//
	// example:
	//
	// time
	XUnit *string `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	// The total number of messages.
	//
	// example:
	//
	// msg
	YUnit *string `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) GetRecords() *OnsTrendGroupOutputTpsResponseBodyDataRecords {
	return s.Records
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) GetXUnit() *string {
	return s.XUnit
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) GetYUnit() *string {
	return s.YUnit
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetRecords(v *OnsTrendGroupOutputTpsResponseBodyDataRecords) *OnsTrendGroupOutputTpsResponseBodyData {
	s.Records = v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetTitle(v string) *OnsTrendGroupOutputTpsResponseBodyData {
	s.Title = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetXUnit(v string) *OnsTrendGroupOutputTpsResponseBodyData {
	s.XUnit = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetYUnit(v string) *OnsTrendGroupOutputTpsResponseBodyData {
	s.YUnit = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) Validate() error {
	if s.Records != nil {
		if err := s.Records.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsTrendGroupOutputTpsResponseBodyDataRecords struct {
	StatsDataDo []*OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo `json:"StatsDataDo,omitempty" xml:"StatsDataDo,omitempty" type:"Repeated"`
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecords) GetStatsDataDo() []*OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo {
	return s.StatsDataDo
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecords) SetStatsDataDo(v []*OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) *OnsTrendGroupOutputTpsResponseBodyDataRecords {
	s.StatsDataDo = v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecords) Validate() error {
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

type OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo struct {
	X *int64   `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) GetX() *int64 {
	return s.X
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) GetY() *float32 {
	return s.Y
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) SetX(v int64) *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo {
	s.X = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) SetY(v float32) *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo {
	s.Y = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) Validate() error {
	return dara.Validate(s)
}
