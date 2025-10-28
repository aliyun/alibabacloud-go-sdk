// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutonomousNotifyEventsInRangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAutonomousNotifyEventsInRangeResponseBody
	GetCode() *string
	SetData(v *GetAutonomousNotifyEventsInRangeResponseBodyData) *GetAutonomousNotifyEventsInRangeResponseBody
	GetData() *GetAutonomousNotifyEventsInRangeResponseBodyData
	SetMessage(v string) *GetAutonomousNotifyEventsInRangeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutonomousNotifyEventsInRangeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetAutonomousNotifyEventsInRangeResponseBody
	GetSuccess() *string
}

type GetAutonomousNotifyEventsInRangeResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *GetAutonomousNotifyEventsInRangeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutonomousNotifyEventsInRangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) GetData() *GetAutonomousNotifyEventsInRangeResponseBodyData {
	return s.Data
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetCode(v string) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetData(v *GetAutonomousNotifyEventsInRangeResponseBodyData) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.Data = v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetMessage(v string) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetRequestId(v string) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) SetSuccess(v string) *GetAutonomousNotifyEventsInRangeResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutonomousNotifyEventsInRangeResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	List *GetAutonomousNotifyEventsInRangeResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAutonomousNotifyEventsInRangeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) GetList() *GetAutonomousNotifyEventsInRangeResponseBodyDataList {
	return s.List
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetExtra(v string) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetList(v *GetAutonomousNotifyEventsInRangeResponseBodyDataList) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.List = v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetPageNo(v int64) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetPageSize(v int64) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) SetTotal(v int64) *GetAutonomousNotifyEventsInRangeResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyData) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutonomousNotifyEventsInRangeResponseBodyDataList struct {
	T []*string `json:"T,omitempty" xml:"T,omitempty" type:"Repeated"`
}

func (s GetAutonomousNotifyEventsInRangeResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetAutonomousNotifyEventsInRangeResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyDataList) GetT() []*string {
	return s.T
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyDataList) SetT(v []*string) *GetAutonomousNotifyEventsInRangeResponseBodyDataList {
	s.T = v
	return s
}

func (s *GetAutonomousNotifyEventsInRangeResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
