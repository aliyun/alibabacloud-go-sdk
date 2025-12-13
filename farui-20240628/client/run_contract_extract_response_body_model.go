// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractExtractResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunContractExtractResponseBody
	GetCode() *string
	SetData(v *RunContractExtractResponseBodyData) *RunContractExtractResponseBody
	GetData() *RunContractExtractResponseBodyData
	SetHttpStatusCode(v int64) *RunContractExtractResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *RunContractExtractResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunContractExtractResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunContractExtractResponseBody
	GetSuccess() *bool
}

type RunContractExtractResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *RunContractExtractResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C844BE6B-33A9-5AC4-A1AE-97B131849E0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RunContractExtractResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunContractExtractResponseBody) GoString() string {
	return s.String()
}

func (s *RunContractExtractResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunContractExtractResponseBody) GetData() *RunContractExtractResponseBodyData {
	return s.Data
}

func (s *RunContractExtractResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *RunContractExtractResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunContractExtractResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunContractExtractResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunContractExtractResponseBody) SetCode(v string) *RunContractExtractResponseBody {
	s.Code = &v
	return s
}

func (s *RunContractExtractResponseBody) SetData(v *RunContractExtractResponseBodyData) *RunContractExtractResponseBody {
	s.Data = v
	return s
}

func (s *RunContractExtractResponseBody) SetHttpStatusCode(v int64) *RunContractExtractResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunContractExtractResponseBody) SetMessage(v string) *RunContractExtractResponseBody {
	s.Message = &v
	return s
}

func (s *RunContractExtractResponseBody) SetRequestId(v string) *RunContractExtractResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunContractExtractResponseBody) SetSuccess(v bool) *RunContractExtractResponseBody {
	s.Success = &v
	return s
}

func (s *RunContractExtractResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunContractExtractResponseBodyData struct {
	ContractText  *string                                            `json:"contractText,omitempty" xml:"contractText,omitempty"`
	ExtractResult []*RunContractExtractResponseBodyDataExtractResult `json:"extractResult,omitempty" xml:"extractResult,omitempty" type:"Repeated"`
}

func (s RunContractExtractResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunContractExtractResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunContractExtractResponseBodyData) GetContractText() *string {
	return s.ContractText
}

func (s *RunContractExtractResponseBodyData) GetExtractResult() []*RunContractExtractResponseBodyDataExtractResult {
	return s.ExtractResult
}

func (s *RunContractExtractResponseBodyData) SetContractText(v string) *RunContractExtractResponseBodyData {
	s.ContractText = &v
	return s
}

func (s *RunContractExtractResponseBodyData) SetExtractResult(v []*RunContractExtractResponseBodyDataExtractResult) *RunContractExtractResponseBodyData {
	s.ExtractResult = v
	return s
}

func (s *RunContractExtractResponseBodyData) Validate() error {
	if s.ExtractResult != nil {
		for _, item := range s.ExtractResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunContractExtractResponseBodyDataExtractResult struct {
	Desc        *string `json:"desc,omitempty" xml:"desc,omitempty"`
	ExtractItem *string `json:"extractItem,omitempty" xml:"extractItem,omitempty"`
	// example:
	//
	// null
	Option *string                                                 `json:"option,omitempty" xml:"option,omitempty"`
	Value  []*RunContractExtractResponseBodyDataExtractResultValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s RunContractExtractResponseBodyDataExtractResult) String() string {
	return dara.Prettify(s)
}

func (s RunContractExtractResponseBodyDataExtractResult) GoString() string {
	return s.String()
}

func (s *RunContractExtractResponseBodyDataExtractResult) GetDesc() *string {
	return s.Desc
}

func (s *RunContractExtractResponseBodyDataExtractResult) GetExtractItem() *string {
	return s.ExtractItem
}

func (s *RunContractExtractResponseBodyDataExtractResult) GetOption() *string {
	return s.Option
}

func (s *RunContractExtractResponseBodyDataExtractResult) GetValue() []*RunContractExtractResponseBodyDataExtractResultValue {
	return s.Value
}

func (s *RunContractExtractResponseBodyDataExtractResult) SetDesc(v string) *RunContractExtractResponseBodyDataExtractResult {
	s.Desc = &v
	return s
}

func (s *RunContractExtractResponseBodyDataExtractResult) SetExtractItem(v string) *RunContractExtractResponseBodyDataExtractResult {
	s.ExtractItem = &v
	return s
}

func (s *RunContractExtractResponseBodyDataExtractResult) SetOption(v string) *RunContractExtractResponseBodyDataExtractResult {
	s.Option = &v
	return s
}

func (s *RunContractExtractResponseBodyDataExtractResult) SetValue(v []*RunContractExtractResponseBodyDataExtractResultValue) *RunContractExtractResponseBodyDataExtractResult {
	s.Value = v
	return s
}

func (s *RunContractExtractResponseBodyDataExtractResult) Validate() error {
	if s.Value != nil {
		for _, item := range s.Value {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunContractExtractResponseBodyDataExtractResultValue struct {
	// example:
	//
	// HT-2022-0001
	Data         *string `json:"data,omitempty" xml:"data,omitempty"`
	OriginalText *string `json:"originalText,omitempty" xml:"originalText,omitempty"`
}

func (s RunContractExtractResponseBodyDataExtractResultValue) String() string {
	return dara.Prettify(s)
}

func (s RunContractExtractResponseBodyDataExtractResultValue) GoString() string {
	return s.String()
}

func (s *RunContractExtractResponseBodyDataExtractResultValue) GetData() *string {
	return s.Data
}

func (s *RunContractExtractResponseBodyDataExtractResultValue) GetOriginalText() *string {
	return s.OriginalText
}

func (s *RunContractExtractResponseBodyDataExtractResultValue) SetData(v string) *RunContractExtractResponseBodyDataExtractResultValue {
	s.Data = &v
	return s
}

func (s *RunContractExtractResponseBodyDataExtractResultValue) SetOriginalText(v string) *RunContractExtractResponseBodyDataExtractResultValue {
	s.OriginalText = &v
	return s
}

func (s *RunContractExtractResponseBodyDataExtractResultValue) Validate() error {
	return dara.Validate(s)
}
