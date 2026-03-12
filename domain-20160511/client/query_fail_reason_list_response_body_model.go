// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryFailReasonListResponseBodyData) *QueryFailReasonListResponseBody
	GetData() *QueryFailReasonListResponseBodyData
	SetRequestId(v string) *QueryFailReasonListResponseBody
	GetRequestId() *string
}

type QueryFailReasonListResponseBody struct {
	Data      *QueryFailReasonListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFailReasonListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFailReasonListResponseBody) GetData() *QueryFailReasonListResponseBodyData {
	return s.Data
}

func (s *QueryFailReasonListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFailReasonListResponseBody) SetData(v *QueryFailReasonListResponseBodyData) *QueryFailReasonListResponseBody {
	s.Data = v
	return s
}

func (s *QueryFailReasonListResponseBody) SetRequestId(v string) *QueryFailReasonListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFailReasonListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFailReasonListResponseBodyData struct {
	FailRecord []*QueryFailReasonListResponseBodyDataFailRecord `json:"FailRecord,omitempty" xml:"FailRecord,omitempty" type:"Repeated"`
}

func (s QueryFailReasonListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFailReasonListResponseBodyData) GetFailRecord() []*QueryFailReasonListResponseBodyDataFailRecord {
	return s.FailRecord
}

func (s *QueryFailReasonListResponseBodyData) SetFailRecord(v []*QueryFailReasonListResponseBodyDataFailRecord) *QueryFailReasonListResponseBodyData {
	s.FailRecord = v
	return s
}

func (s *QueryFailReasonListResponseBodyData) Validate() error {
	if s.FailRecord != nil {
		for _, item := range s.FailRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFailReasonListResponseBodyDataFailRecord struct {
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty"`
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s QueryFailReasonListResponseBodyDataFailRecord) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonListResponseBodyDataFailRecord) GoString() string {
	return s.String()
}

func (s *QueryFailReasonListResponseBodyDataFailRecord) GetDate() *string {
	return s.Date
}

func (s *QueryFailReasonListResponseBodyDataFailRecord) GetFailReason() *string {
	return s.FailReason
}

func (s *QueryFailReasonListResponseBodyDataFailRecord) SetDate(v string) *QueryFailReasonListResponseBodyDataFailRecord {
	s.Date = &v
	return s
}

func (s *QueryFailReasonListResponseBodyDataFailRecord) SetFailReason(v string) *QueryFailReasonListResponseBodyDataFailRecord {
	s.FailReason = &v
	return s
}

func (s *QueryFailReasonListResponseBodyDataFailRecord) Validate() error {
	return dara.Validate(s)
}
