// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailingReasonListForQualificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryFailingReasonListForQualificationResponseBodyData) *QueryFailingReasonListForQualificationResponseBody
	GetData() []*QueryFailingReasonListForQualificationResponseBodyData
	SetRequestId(v string) *QueryFailingReasonListForQualificationResponseBody
	GetRequestId() *string
}

type QueryFailingReasonListForQualificationResponseBody struct {
	// List of domain name qualification verification failures.
	Data []*QueryFailingReasonListForQualificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFailingReasonListForQualificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFailingReasonListForQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFailingReasonListForQualificationResponseBody) GetData() []*QueryFailingReasonListForQualificationResponseBodyData {
	return s.Data
}

func (s *QueryFailingReasonListForQualificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFailingReasonListForQualificationResponseBody) SetData(v []*QueryFailingReasonListForQualificationResponseBodyData) *QueryFailingReasonListForQualificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFailingReasonListForQualificationResponseBody) SetRequestId(v string) *QueryFailingReasonListForQualificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFailingReasonListForQualificationResponseBody) Validate() error {
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

type QueryFailingReasonListForQualificationResponseBodyData struct {
	// Review date.
	//
	// example:
	//
	// 2017-03-17 11:08:02
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Reason for domain name qualification verification failure.
	//
	// example:
	//
	// 证件审核不通过
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s QueryFailingReasonListForQualificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFailingReasonListForQualificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFailingReasonListForQualificationResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *QueryFailingReasonListForQualificationResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *QueryFailingReasonListForQualificationResponseBodyData) SetDate(v string) *QueryFailingReasonListForQualificationResponseBodyData {
	s.Date = &v
	return s
}

func (s *QueryFailingReasonListForQualificationResponseBodyData) SetFailReason(v string) *QueryFailingReasonListForQualificationResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *QueryFailingReasonListForQualificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
