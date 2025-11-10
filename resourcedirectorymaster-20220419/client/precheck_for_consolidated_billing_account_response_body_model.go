// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckForConsolidatedBillingAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReasons(v []*PrecheckForConsolidatedBillingAccountResponseBodyReasons) *PrecheckForConsolidatedBillingAccountResponseBody
	GetReasons() []*PrecheckForConsolidatedBillingAccountResponseBodyReasons
	SetRequestId(v string) *PrecheckForConsolidatedBillingAccountResponseBody
	GetRequestId() *string
	SetResult(v bool) *PrecheckForConsolidatedBillingAccountResponseBody
	GetResult() *bool
}

type PrecheckForConsolidatedBillingAccountResponseBody struct {
	// The cause of the check failure.
	Reasons []*PrecheckForConsolidatedBillingAccountResponseBodyReasons `json:"Reasons,omitempty" xml:"Reasons,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9E6B6CA8-9E7A-521F-A743-AA582714727E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the check was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PrecheckForConsolidatedBillingAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrecheckForConsolidatedBillingAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) GetReasons() []*PrecheckForConsolidatedBillingAccountResponseBodyReasons {
	return s.Reasons
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) SetReasons(v []*PrecheckForConsolidatedBillingAccountResponseBodyReasons) *PrecheckForConsolidatedBillingAccountResponseBody {
	s.Reasons = v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) SetRequestId(v string) *PrecheckForConsolidatedBillingAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) SetResult(v bool) *PrecheckForConsolidatedBillingAccountResponseBody {
	s.Result = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponseBody) Validate() error {
	if s.Reasons != nil {
		for _, item := range s.Reasons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PrecheckForConsolidatedBillingAccountResponseBodyReasons struct {
	// The error code.
	//
	// example:
	//
	// PaymentAccountEnterpriseInvoiceError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// No enterprise invoice header information is set for the payment account.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s PrecheckForConsolidatedBillingAccountResponseBodyReasons) String() string {
	return dara.Prettify(s)
}

func (s PrecheckForConsolidatedBillingAccountResponseBodyReasons) GoString() string {
	return s.String()
}

func (s *PrecheckForConsolidatedBillingAccountResponseBodyReasons) GetCode() *string {
	return s.Code
}

func (s *PrecheckForConsolidatedBillingAccountResponseBodyReasons) GetMessage() *string {
	return s.Message
}

func (s *PrecheckForConsolidatedBillingAccountResponseBodyReasons) SetCode(v string) *PrecheckForConsolidatedBillingAccountResponseBodyReasons {
	s.Code = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponseBodyReasons) SetMessage(v string) *PrecheckForConsolidatedBillingAccountResponseBodyReasons {
	s.Message = &v
	return s
}

func (s *PrecheckForConsolidatedBillingAccountResponseBodyReasons) Validate() error {
	return dara.Validate(s)
}
