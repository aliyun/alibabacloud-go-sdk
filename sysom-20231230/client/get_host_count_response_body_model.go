// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHostCountResponseBody
	GetCode() *string
	SetData(v []*GetHostCountResponseBodyData) *GetHostCountResponseBody
	GetData() []*GetHostCountResponseBodyData
	SetMessage(v string) *GetHostCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHostCountResponseBody
	GetRequestId() *string
	SetTotal(v int64) *GetHostCountResponseBody
	GetTotal() *int64
}

type GetHostCountResponseBody struct {
	// example:
	//
	// Success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetHostCountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// “”
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 43A910E9-A739-525E-855D-A32C257F1826
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 3
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetHostCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHostCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHostCountResponseBody) GetData() []*GetHostCountResponseBodyData {
	return s.Data
}

func (s *GetHostCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHostCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHostCountResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetHostCountResponseBody) SetCode(v string) *GetHostCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetHostCountResponseBody) SetData(v []*GetHostCountResponseBodyData) *GetHostCountResponseBody {
	s.Data = v
	return s
}

func (s *GetHostCountResponseBody) SetMessage(v string) *GetHostCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetHostCountResponseBody) SetRequestId(v string) *GetHostCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostCountResponseBody) SetTotal(v int64) *GetHostCountResponseBody {
	s.Total = &v
	return s
}

func (s *GetHostCountResponseBody) Validate() error {
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

type GetHostCountResponseBodyData struct {
	// example:
	//
	// 1725797727754
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 5
	Value *int32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetHostCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHostCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHostCountResponseBodyData) GetTime() *int64 {
	return s.Time
}

func (s *GetHostCountResponseBodyData) GetValue() *int32 {
	return s.Value
}

func (s *GetHostCountResponseBodyData) SetTime(v int64) *GetHostCountResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetHostCountResponseBodyData) SetValue(v int32) *GetHostCountResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetHostCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
