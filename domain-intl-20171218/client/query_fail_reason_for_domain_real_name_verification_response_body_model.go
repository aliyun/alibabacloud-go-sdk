// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFailReasonForDomainRealNameVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryFailReasonForDomainRealNameVerificationResponseBodyData) *QueryFailReasonForDomainRealNameVerificationResponseBody
	GetData() []*QueryFailReasonForDomainRealNameVerificationResponseBodyData
	SetRequestId(v string) *QueryFailReasonForDomainRealNameVerificationResponseBody
	GetRequestId() *string
}

type QueryFailReasonForDomainRealNameVerificationResponseBody struct {
	Data      []*QueryFailReasonForDomainRealNameVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFailReasonForDomainRealNameVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonForDomainRealNameVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBody) GetData() []*QueryFailReasonForDomainRealNameVerificationResponseBodyData {
	return s.Data
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBody) SetData(v []*QueryFailReasonForDomainRealNameVerificationResponseBodyData) *QueryFailReasonForDomainRealNameVerificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBody) SetRequestId(v string) *QueryFailReasonForDomainRealNameVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBody) Validate() error {
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

type QueryFailReasonForDomainRealNameVerificationResponseBodyData struct {
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty"`
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s QueryFailReasonForDomainRealNameVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFailReasonForDomainRealNameVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBodyData) SetDate(v string) *QueryFailReasonForDomainRealNameVerificationResponseBodyData {
	s.Date = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBodyData) SetFailReason(v string) *QueryFailReasonForDomainRealNameVerificationResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
