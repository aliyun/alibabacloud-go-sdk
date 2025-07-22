// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanShareAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetSavingPlanShareAccountsResponseBodyData) *GetSavingPlanShareAccountsResponseBody
	GetData() []*GetSavingPlanShareAccountsResponseBodyData
	SetRequestId(v string) *GetSavingPlanShareAccountsResponseBody
	GetRequestId() *string
}

type GetSavingPlanShareAccountsResponseBody struct {
	Data      []*GetSavingPlanShareAccountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSavingPlanShareAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanShareAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsResponseBody) GetData() []*GetSavingPlanShareAccountsResponseBodyData {
	return s.Data
}

func (s *GetSavingPlanShareAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSavingPlanShareAccountsResponseBody) SetData(v []*GetSavingPlanShareAccountsResponseBodyData) *GetSavingPlanShareAccountsResponseBody {
	s.Data = v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBody) SetRequestId(v string) *GetSavingPlanShareAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanShareAccountsResponseBodyData struct {
	AccountId     *string                                                    `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AliUid        *int64                                                     `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ShareTimeList []*GetSavingPlanShareAccountsResponseBodyDataShareTimeList `json:"ShareTimeList,omitempty" xml:"ShareTimeList,omitempty" type:"Repeated"`
}

func (s GetSavingPlanShareAccountsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanShareAccountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *GetSavingPlanShareAccountsResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetSavingPlanShareAccountsResponseBodyData) GetShareTimeList() []*GetSavingPlanShareAccountsResponseBodyDataShareTimeList {
	return s.ShareTimeList
}

func (s *GetSavingPlanShareAccountsResponseBodyData) SetAccountId(v string) *GetSavingPlanShareAccountsResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBodyData) SetAliUid(v int64) *GetSavingPlanShareAccountsResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBodyData) SetShareTimeList(v []*GetSavingPlanShareAccountsResponseBodyDataShareTimeList) *GetSavingPlanShareAccountsResponseBodyData {
	s.ShareTimeList = v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanShareAccountsResponseBodyDataShareTimeList struct {
	ShareEndTime   *string `json:"ShareEndTime,omitempty" xml:"ShareEndTime,omitempty"`
	ShareStartTime *string `json:"ShareStartTime,omitempty" xml:"ShareStartTime,omitempty"`
}

func (s GetSavingPlanShareAccountsResponseBodyDataShareTimeList) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanShareAccountsResponseBodyDataShareTimeList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsResponseBodyDataShareTimeList) GetShareEndTime() *string {
	return s.ShareEndTime
}

func (s *GetSavingPlanShareAccountsResponseBodyDataShareTimeList) GetShareStartTime() *string {
	return s.ShareStartTime
}

func (s *GetSavingPlanShareAccountsResponseBodyDataShareTimeList) SetShareEndTime(v string) *GetSavingPlanShareAccountsResponseBodyDataShareTimeList {
	s.ShareEndTime = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBodyDataShareTimeList) SetShareStartTime(v string) *GetSavingPlanShareAccountsResponseBodyDataShareTimeList {
	s.ShareStartTime = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBodyDataShareTimeList) Validate() error {
	return dara.Validate(s)
}
