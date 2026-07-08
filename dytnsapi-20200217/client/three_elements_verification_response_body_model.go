// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThreeElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ThreeElementsVerificationResponseBody
	GetCode() *string
	SetData(v *ThreeElementsVerificationResponseBodyData) *ThreeElementsVerificationResponseBody
	GetData() *ThreeElementsVerificationResponseBodyData
	SetMessage(v string) *ThreeElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ThreeElementsVerificationResponseBody
	GetRequestId() *string
}

type ThreeElementsVerificationResponseBody struct {
	// The request status code.
	//
	// - **OK**: The request was successful.
	//
	// - For other error codes, see the error code table in this chapter.
	//
	// - **RequestFrequencyLimit**: Due to carrier restrictions, repeated high-frequency queries on the same number within a short period are prohibited. If this error code is returned, try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ThreeElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ThreeElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ThreeElementsVerificationResponseBody) GetData() *ThreeElementsVerificationResponseBodyData {
	return s.Data
}

func (s *ThreeElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ThreeElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ThreeElementsVerificationResponseBody) SetCode(v string) *ThreeElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetData(v *ThreeElementsVerificationResponseBodyData) *ThreeElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetMessage(v string) *ThreeElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetRequestId(v string) *ThreeElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ThreeElementsVerificationResponseBodyData struct {
	// The basic carrier. Valid values:
	//
	// - **China Mobile**.
	//
	// - **China Unicom**.
	//
	// - **China Telecom**.
	//
	// example:
	//
	// 中国移动
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// Indicates whether the verification results are consistent. Returned values:
	//
	// - **1**: Consistent
	//
	// - **0**: Inconsistent
	//
	// - **2**: Not found
	//
	// >The data update timeliness varies by carrier and city, and is typically between T+1 and T+3.
	//
	// The verification results for mobile phone numbers of different carriers in different states are as follows:
	//
	// |Carrier/Mobile Phone Number Status|Suspended|Empty Number|Deregistered|
	//
	// |--|--|--|--|
	//
	// |China Mobile|Normal verification|Not found|Not found|
	//
	// |China Unicom|Normal verification|Inconsistent|Inconsistent|
	//
	// |China Telecom|Normal verification|Not found|Not found|
	//
	// example:
	//
	// 1
	IsConsistent *int32 `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s ThreeElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponseBodyData) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *ThreeElementsVerificationResponseBodyData) GetIsConsistent() *int32 {
	return s.IsConsistent
}

func (s *ThreeElementsVerificationResponseBodyData) SetBasicCarrier(v string) *ThreeElementsVerificationResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *ThreeElementsVerificationResponseBodyData) SetIsConsistent(v int32) *ThreeElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *ThreeElementsVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
