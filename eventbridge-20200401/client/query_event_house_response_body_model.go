// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventHouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryEventHouseResponseBody
	GetCode() *string
	SetData(v *QueryEventHouseResponseBodyData) *QueryEventHouseResponseBody
	GetData() *QueryEventHouseResponseBodyData
	SetMessage(v string) *QueryEventHouseResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryEventHouseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEventHouseResponseBody
	GetSuccess() *bool
}

type QueryEventHouseResponseBody struct {
	// example:
	//
	// Success
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryEventHouseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Remote error. requestId: [xxxx], error code: [xxx], message: [The target in event rule is invalid! Endpoint is xxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// f2099962-1628-45f1-9782-2bf6daad823f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventHouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEventHouseResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventHouseResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryEventHouseResponseBody) GetData() *QueryEventHouseResponseBodyData {
	return s.Data
}

func (s *QueryEventHouseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEventHouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEventHouseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEventHouseResponseBody) SetCode(v string) *QueryEventHouseResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventHouseResponseBody) SetData(v *QueryEventHouseResponseBodyData) *QueryEventHouseResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventHouseResponseBody) SetMessage(v string) *QueryEventHouseResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEventHouseResponseBody) SetRequestId(v string) *QueryEventHouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventHouseResponseBody) SetSuccess(v bool) *QueryEventHouseResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEventHouseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEventHouseResponseBodyData struct {
	Rows []*Row `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// example:
	//
	// 18
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryEventHouseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryEventHouseResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventHouseResponseBodyData) GetRows() []*Row {
	return s.Rows
}

func (s *QueryEventHouseResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryEventHouseResponseBodyData) SetRows(v []*Row) *QueryEventHouseResponseBodyData {
	s.Rows = v
	return s
}

func (s *QueryEventHouseResponseBodyData) SetTotal(v int32) *QueryEventHouseResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryEventHouseResponseBodyData) Validate() error {
	if s.Rows != nil {
		for _, item := range s.Rows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
