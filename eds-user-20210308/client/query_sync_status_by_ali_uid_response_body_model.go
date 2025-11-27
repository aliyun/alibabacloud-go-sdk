// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySyncStatusByAliUidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySyncStatusByAliUidResponseBody
	GetCode() *string
	SetData(v *QuerySyncStatusByAliUidResponseBodyData) *QuerySyncStatusByAliUidResponseBody
	GetData() *QuerySyncStatusByAliUidResponseBodyData
	SetHttpStatusCode(v int32) *QuerySyncStatusByAliUidResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QuerySyncStatusByAliUidResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySyncStatusByAliUidResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySyncStatusByAliUidResponseBody
	GetSuccess() *bool
}

type QuerySyncStatusByAliUidResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QuerySyncStatusByAliUidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySyncStatusByAliUidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySyncStatusByAliUidResponseBody) GetData() *QuerySyncStatusByAliUidResponseBodyData {
	return s.Data
}

func (s *QuerySyncStatusByAliUidResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySyncStatusByAliUidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySyncStatusByAliUidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySyncStatusByAliUidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySyncStatusByAliUidResponseBody) SetCode(v string) *QuerySyncStatusByAliUidResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetData(v *QuerySyncStatusByAliUidResponseBodyData) *QuerySyncStatusByAliUidResponseBody {
	s.Data = v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetHttpStatusCode(v int32) *QuerySyncStatusByAliUidResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetMessage(v string) *QuerySyncStatusByAliUidResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetRequestId(v string) *QuerySyncStatusByAliUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetSuccess(v bool) *QuerySyncStatusByAliUidResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySyncStatusByAliUidResponseBodyData struct {
	// example:
	//
	// 131239236086****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// cdrs948144195608****
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 2020-06-30 07:50:42
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2022-03-02 14:27:39
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 18500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-03-02 14:31:39
	LatestBeginTime *string `json:"LatestBeginTime,omitempty" xml:"LatestBeginTime,omitempty"`
	// example:
	//
	// 2022-03-02 16:13:12
	LatestEndTime *string `json:"LatestEndTime,omitempty" xml:"LatestEndTime,omitempty"`
	// example:
	//
	// 2022-03-02 18:24:01
	LatestSuccessTime *string `json:"LatestSuccessTime,omitempty" xml:"LatestSuccessTime,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QuerySyncStatusByAliUidResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetCorpId() *string {
	return s.CorpId
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetLatestBeginTime() *string {
	return s.LatestBeginTime
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetLatestEndTime() *string {
	return s.LatestEndTime
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetLatestSuccessTime() *string {
	return s.LatestSuccessTime
}

func (s *QuerySyncStatusByAliUidResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetAliUid(v int64) *QuerySyncStatusByAliUidResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetCorpId(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetGmtCreated(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetGmtModified(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetId(v int64) *QuerySyncStatusByAliUidResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestBeginTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestBeginTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestEndTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestEndTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestSuccessTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestSuccessTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetStatus(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.Status = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) Validate() error {
	return dara.Validate(s)
}
