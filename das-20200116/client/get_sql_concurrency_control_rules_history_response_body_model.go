// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConcurrencyControlRulesHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetCode() *string
	SetData(v *GetSqlConcurrencyControlRulesHistoryResponseBodyData) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetData() *GetSqlConcurrencyControlRulesHistoryResponseBodyData
	SetMessage(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetSuccess() *string
}

type GetSqlConcurrencyControlRulesHistoryResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *GetSqlConcurrencyControlRulesHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetData() *GetSqlConcurrencyControlRulesHistoryResponseBodyData {
	return s.Data
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetCode(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetData(v *GetSqlConcurrencyControlRulesHistoryResponseBodyData) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetMessage(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetRequestId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetSuccess(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyData struct {
	List *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) GetList() *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList {
	return s.List
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) SetList(v *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) *GetSqlConcurrencyControlRulesHistoryResponseBodyData {
	s.List = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) SetTotal(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyDataList struct {
	Rules []*GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) GetRules() []*GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	return s.Rules
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) SetRules(v []*GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList {
	s.Rules = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules struct {
	ConcurrencyControlTime *int64  `json:"ConcurrencyControlTime,omitempty" xml:"ConcurrencyControlTime,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemId                 *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	KeywordsHash           *string `json:"KeywordsHash,omitempty" xml:"KeywordsHash,omitempty"`
	MaxConcurrency         *int64  `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	SqlKeywords            *string `json:"SqlKeywords,omitempty" xml:"SqlKeywords,omitempty"`
	SqlType                *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	StartTime              *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetConcurrencyControlTime() *int64 {
	return s.ConcurrencyControlTime
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetKeywordsHash() *string {
	return s.KeywordsHash
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetMaxConcurrency() *int64 {
	return s.MaxConcurrency
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetSqlKeywords() *string {
	return s.SqlKeywords
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetSqlType() *string {
	return s.SqlType
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetStatus() *string {
	return s.Status
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetUserId() *string {
	return s.UserId
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetConcurrencyControlTime(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.ConcurrencyControlTime = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetInstanceId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.InstanceId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetItemId(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.ItemId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetKeywordsHash(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.KeywordsHash = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetMaxConcurrency(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.MaxConcurrency = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetSqlKeywords(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.SqlKeywords = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetSqlType(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.SqlType = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetStartTime(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.StartTime = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetStatus(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.Status = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetUserId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.UserId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) Validate() error {
	return dara.Validate(s)
}
