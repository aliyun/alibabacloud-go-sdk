// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTwoElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TwoElementsVerificationResponseBody
	GetCode() *string
	SetData(v *TwoElementsVerificationResponseBodyData) *TwoElementsVerificationResponseBody
	GetData() *TwoElementsVerificationResponseBodyData
	SetMessage(v string) *TwoElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *TwoElementsVerificationResponseBody
	GetRequestId() *string
}

type TwoElementsVerificationResponseBody struct {
	// The request status code.
	//
	// - **OK**: The request was successful.
	//
	// - For other error codes, see the error code table in this chapter.
	//
	// - **RequestFrequencyLimit**: Due to operator restrictions, repeated high-frequency queries against the same number or name in a short period are prohibited. If this error code is returned, try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The structure.
	Data *TwoElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TwoElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TwoElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *TwoElementsVerificationResponseBody) GetData() *TwoElementsVerificationResponseBodyData {
	return s.Data
}

func (s *TwoElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TwoElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TwoElementsVerificationResponseBody) SetCode(v string) *TwoElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetData(v *TwoElementsVerificationResponseBodyData) *TwoElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetMessage(v string) *TwoElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetRequestId(v string) *TwoElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TwoElementsVerificationResponseBodyData struct {
	// The basic operator. Valid values:
	//
	// - **China Mobile**.
	//
	// - **China Unicom**.
	//
	// - **China Telecom**.
	//
	// 	Notice: China Broadcasting Network numbers are not currently supported.
	//
	// example:
	//
	// 中国移动
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// Indicates whether the verification result is consistent. Returns:
	//
	// - **1**: Consistent.
	//
	// - **0**: Inconsistent.
	//
	// - **2**: Not found.
	//
	// The data update timeliness for different operators and cities is typically T+1 to T+3.
	//
	// The verification results for different operator phone numbers in different states are as follows:
	//
	// |Operator/Phone Number Status|Suspended|Empty Number|Cancelled|
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

func (s TwoElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TwoElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponseBodyData) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *TwoElementsVerificationResponseBodyData) GetIsConsistent() *int32 {
	return s.IsConsistent
}

func (s *TwoElementsVerificationResponseBodyData) SetBasicCarrier(v string) *TwoElementsVerificationResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *TwoElementsVerificationResponseBodyData) SetIsConsistent(v int32) *TwoElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *TwoElementsVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
