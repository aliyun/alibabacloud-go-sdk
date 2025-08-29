// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalculationJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCalculationJobIds(v []*string) *CreateCalculationJobsResponseBody
	GetCalculationJobIds() []*string
	SetRequestId(v string) *CreateCalculationJobsResponseBody
	GetRequestId() *string
}

type CreateCalculationJobsResponseBody struct {
	CalculationJobIds []*string `json:"CalculationJobIds,omitempty" xml:"CalculationJobIds,omitempty" type:"Repeated"`
	// example:
	//
	// 8C27790E-CCA5-56BB-BA17-646295DEC0A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCalculationJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCalculationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCalculationJobsResponseBody) GetCalculationJobIds() []*string {
	return s.CalculationJobIds
}

func (s *CreateCalculationJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCalculationJobsResponseBody) SetCalculationJobIds(v []*string) *CreateCalculationJobsResponseBody {
	s.CalculationJobIds = v
	return s
}

func (s *CreateCalculationJobsResponseBody) SetRequestId(v string) *CreateCalculationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCalculationJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
