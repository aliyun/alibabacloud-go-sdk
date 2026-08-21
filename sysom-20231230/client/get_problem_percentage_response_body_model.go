// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProblemPercentageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProblemPercentageResponseBody
	GetCode() *string
	SetData(v []*GetProblemPercentageResponseBodyData) *GetProblemPercentageResponseBody
	GetData() []*GetProblemPercentageResponseBodyData
	SetMessage(v string) *GetProblemPercentageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProblemPercentageResponseBody
	GetRequestId() *string
	SetTotal(v int64) *GetProblemPercentageResponseBody
	GetTotal() *int64
}

type GetProblemPercentageResponseBody struct {
	// The status code.
	//
	// - `code == Success` indicates that the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. Check the `message` field for the detailed fault information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*GetProblemPercentageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error information.
	//
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The total amount of data.
	//
	// example:
	//
	// 19
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetProblemPercentageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProblemPercentageResponseBody) GoString() string {
	return s.String()
}

func (s *GetProblemPercentageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProblemPercentageResponseBody) GetData() []*GetProblemPercentageResponseBodyData {
	return s.Data
}

func (s *GetProblemPercentageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProblemPercentageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProblemPercentageResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetProblemPercentageResponseBody) SetCode(v string) *GetProblemPercentageResponseBody {
	s.Code = &v
	return s
}

func (s *GetProblemPercentageResponseBody) SetData(v []*GetProblemPercentageResponseBodyData) *GetProblemPercentageResponseBody {
	s.Data = v
	return s
}

func (s *GetProblemPercentageResponseBody) SetMessage(v string) *GetProblemPercentageResponseBody {
	s.Message = &v
	return s
}

func (s *GetProblemPercentageResponseBody) SetRequestId(v string) *GetProblemPercentageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProblemPercentageResponseBody) SetTotal(v int64) *GetProblemPercentageResponseBody {
	s.Total = &v
	return s
}

func (s *GetProblemPercentageResponseBody) Validate() error {
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

type GetProblemPercentageResponseBodyData struct {
	// The type of the abnormal item.
	//
	// example:
	//
	// saturation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The number of abnormal items of this type.
	//
	// example:
	//
	// 5
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetProblemPercentageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProblemPercentageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProblemPercentageResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetProblemPercentageResponseBodyData) GetValue() *int64 {
	return s.Value
}

func (s *GetProblemPercentageResponseBodyData) SetType(v string) *GetProblemPercentageResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetProblemPercentageResponseBodyData) SetValue(v int64) *GetProblemPercentageResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetProblemPercentageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
