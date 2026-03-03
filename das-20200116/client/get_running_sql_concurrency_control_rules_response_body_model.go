// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningSqlConcurrencyControlRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetCode() *string
	SetData(v *GetRunningSqlConcurrencyControlRulesResponseBodyData) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetData() *GetRunningSqlConcurrencyControlRulesResponseBodyData
	SetMessage(v string) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetSuccess() *string
}

type GetRunningSqlConcurrencyControlRulesResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *GetRunningSqlConcurrencyControlRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetRunningSqlConcurrencyControlRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetData() *GetRunningSqlConcurrencyControlRulesResponseBodyData {
	return s.Data
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetCode(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Code = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetData(v *GetRunningSqlConcurrencyControlRulesResponseBodyData) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Data = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetMessage(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Message = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetRequestId(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetSuccess(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Success = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRunningSqlConcurrencyControlRulesResponseBodyData struct {
	List *GetRunningSqlConcurrencyControlRulesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) GetList() *GetRunningSqlConcurrencyControlRulesResponseBodyDataList {
	return s.List
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) SetList(v *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) *GetRunningSqlConcurrencyControlRulesResponseBodyData {
	s.List = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) SetTotal(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRunningSqlConcurrencyControlRulesResponseBodyDataList struct {
	RunningRules []*GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules `json:"runningRules,omitempty" xml:"runningRules,omitempty" type:"Repeated"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) GetRunningRules() []*GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	return s.RunningRules
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) SetRunningRules(v []*GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) *GetRunningSqlConcurrencyControlRulesResponseBodyDataList {
	s.RunningRules = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) Validate() error {
	if s.RunningRules != nil {
		for _, item := range s.RunningRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules struct {
	ConcurrencyControlTime *int64  `json:"ConcurrencyControlTime,omitempty" xml:"ConcurrencyControlTime,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemId                 *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	KeywordsHash           *string `json:"KeywordsHash,omitempty" xml:"KeywordsHash,omitempty"`
	MaxConcurrency         *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	SqlKeywords            *string `json:"SqlKeywords,omitempty" xml:"SqlKeywords,omitempty"`
	SqlType                *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	StartTime              *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetConcurrencyControlTime() *int64 {
	return s.ConcurrencyControlTime
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetKeywordsHash() *string {
	return s.KeywordsHash
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetMaxConcurrency() *string {
	return s.MaxConcurrency
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetSqlKeywords() *string {
	return s.SqlKeywords
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetSqlType() *string {
	return s.SqlType
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetStatus() *string {
	return s.Status
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetUserId() *string {
	return s.UserId
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetConcurrencyControlTime(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.ConcurrencyControlTime = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetInstanceId(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.InstanceId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetItemId(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.ItemId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetKeywordsHash(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.KeywordsHash = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetMaxConcurrency(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.MaxConcurrency = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetSqlKeywords(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.SqlKeywords = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetSqlType(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.SqlType = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetStartTime(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.StartTime = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetStatus(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.Status = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetUserId(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.UserId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) Validate() error {
	return dara.Validate(s)
}
