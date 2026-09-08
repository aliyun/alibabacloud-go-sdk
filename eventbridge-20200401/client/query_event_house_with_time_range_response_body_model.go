// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventHouseWithTimeRangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryEventHouseWithTimeRangeResponseBody
	GetCode() *string
	SetData(v *QueryEventHouseWithTimeRangeResponseBodyData) *QueryEventHouseWithTimeRangeResponseBody
	GetData() *QueryEventHouseWithTimeRangeResponseBodyData
	SetMessage(v string) *QueryEventHouseWithTimeRangeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryEventHouseWithTimeRangeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEventHouseWithTimeRangeResponseBody
	GetSuccess() *bool
}

type QueryEventHouseWithTimeRangeResponseBody struct {
	// The return code of the operation. Success indicates a successful call. Other values indicate specific error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The structured result data returned by the SQL query.
	Data *QueryEventHouseWithTimeRangeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// A success message if the call succeeds, or a specific error message if the call fails.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. You can use this ID for troubleshooting.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventHouseWithTimeRangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEventHouseWithTimeRangeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventHouseWithTimeRangeResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryEventHouseWithTimeRangeResponseBody) GetData() *QueryEventHouseWithTimeRangeResponseBodyData {
	return s.Data
}

func (s *QueryEventHouseWithTimeRangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEventHouseWithTimeRangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEventHouseWithTimeRangeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEventHouseWithTimeRangeResponseBody) SetCode(v string) *QueryEventHouseWithTimeRangeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponseBody) SetData(v *QueryEventHouseWithTimeRangeResponseBodyData) *QueryEventHouseWithTimeRangeResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponseBody) SetMessage(v string) *QueryEventHouseWithTimeRangeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponseBody) SetRequestId(v string) *QueryEventHouseWithTimeRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponseBody) SetSuccess(v bool) *QueryEventHouseWithTimeRangeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEventHouseWithTimeRangeResponseBodyData struct {
	// The list of result rows returned by the SQL query.
	Rows []*Row `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// The number of result rows actually returned by the query.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryEventHouseWithTimeRangeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryEventHouseWithTimeRangeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventHouseWithTimeRangeResponseBodyData) GetRows() []*Row {
	return s.Rows
}

func (s *QueryEventHouseWithTimeRangeResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryEventHouseWithTimeRangeResponseBodyData) SetRows(v []*Row) *QueryEventHouseWithTimeRangeResponseBodyData {
	s.Rows = v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponseBodyData) SetTotal(v int32) *QueryEventHouseWithTimeRangeResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeResponseBodyData) Validate() error {
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
