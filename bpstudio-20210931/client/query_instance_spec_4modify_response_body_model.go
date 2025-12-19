// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceSpec4ModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInstanceSpec4ModifyResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryInstanceSpec4ModifyResponseBody
	GetCode() *string
	SetData(v *QueryInstanceSpec4ModifyResponseBodyData) *QueryInstanceSpec4ModifyResponseBody
	GetData() *QueryInstanceSpec4ModifyResponseBodyData
	SetMessage(v string) *QueryInstanceSpec4ModifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInstanceSpec4ModifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInstanceSpec4ModifyResponseBody
	GetSuccess() *bool
}

type QueryInstanceSpec4ModifyResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryInstanceSpec4ModifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 847C9D0A-BABD-589C-8A9C-6464409EDED9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstanceSpec4ModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceSpec4ModifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInstanceSpec4ModifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryInstanceSpec4ModifyResponseBody) GetData() *QueryInstanceSpec4ModifyResponseBodyData {
	return s.Data
}

func (s *QueryInstanceSpec4ModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInstanceSpec4ModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstanceSpec4ModifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetAccessDeniedDetail(v string) *QueryInstanceSpec4ModifyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetCode(v string) *QueryInstanceSpec4ModifyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetData(v *QueryInstanceSpec4ModifyResponseBodyData) *QueryInstanceSpec4ModifyResponseBody {
	s.Data = v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetMessage(v string) *QueryInstanceSpec4ModifyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetRequestId(v string) *QueryInstanceSpec4ModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetSuccess(v bool) *QueryInstanceSpec4ModifyResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInstanceSpec4ModifyResponseBodyData struct {
	OptionalValues []*QueryInstanceSpec4ModifyResponseBodyDataOptionalValues `json:"OptionalValues,omitempty" xml:"OptionalValues,omitempty" type:"Repeated"`
}

func (s QueryInstanceSpec4ModifyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceSpec4ModifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyResponseBodyData) GetOptionalValues() []*QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	return s.OptionalValues
}

func (s *QueryInstanceSpec4ModifyResponseBodyData) SetOptionalValues(v []*QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) *QueryInstanceSpec4ModifyResponseBodyData {
	s.OptionalValues = v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyData) Validate() error {
	if s.OptionalValues != nil {
		for _, item := range s.OptionalValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryInstanceSpec4ModifyResponseBodyDataOptionalValues struct {
	Label *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	Max   *float64 `json:"Max,omitempty" xml:"Max,omitempty"`
	Min   *float64 `json:"Min,omitempty" xml:"Min,omitempty"`
	Step  *float64 `json:"Step,omitempty" xml:"Step,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) GetLabel() *string {
	return s.Label
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) GetMax() *float64 {
	return s.Max
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) GetMin() *float64 {
	return s.Min
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) GetStep() *float64 {
	return s.Step
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) GetValue() *string {
	return s.Value
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetLabel(v string) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Label = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetMax(v float64) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Max = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetMin(v float64) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Min = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetStep(v float64) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Step = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetValue(v string) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Value = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) Validate() error {
	return dara.Validate(s)
}
