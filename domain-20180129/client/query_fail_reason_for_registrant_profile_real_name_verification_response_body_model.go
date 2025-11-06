// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonForRegistrantProfileRealNameVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody
	GetData() []*QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData
	SetRequestId(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody
	GetRequestId() *string
}

type QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody struct {
	Data []*QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 548C407F-AEA2-4B5D-90DF-EC11EBB1D76F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) GetData() []*QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData {
	return s.Data
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) SetData(v []*QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) SetRequestId(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) Validate() error {
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

type QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData struct {
	// example:
	//
	// 2017-03-17 11:08:02
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty"`
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) SetDate(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData {
	s.Date = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) SetFailReason(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
