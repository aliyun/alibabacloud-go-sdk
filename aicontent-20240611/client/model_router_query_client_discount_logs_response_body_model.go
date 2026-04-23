// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientDiscountLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ModelRouterQueryClientDiscountLogsResponseBodyData) *ModelRouterQueryClientDiscountLogsResponseBody
	GetData() []*ModelRouterQueryClientDiscountLogsResponseBodyData
	SetErrCode(v string) *ModelRouterQueryClientDiscountLogsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryClientDiscountLogsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryClientDiscountLogsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryClientDiscountLogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryClientDiscountLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryClientDiscountLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryClientDiscountLogsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryClientDiscountLogsResponseBody struct {
	Data []*ModelRouterQueryClientDiscountLogsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryClientDiscountLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientDiscountLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) GetData() []*ModelRouterQueryClientDiscountLogsResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) SetData(v []*ModelRouterQueryClientDiscountLogsResponseBodyData) *ModelRouterQueryClientDiscountLogsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) SetErrCode(v string) *ModelRouterQueryClientDiscountLogsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) SetErrMessage(v string) *ModelRouterQueryClientDiscountLogsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryClientDiscountLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) SetMaxResults(v int32) *ModelRouterQueryClientDiscountLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) SetNextToken(v string) *ModelRouterQueryClientDiscountLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) SetRequestId(v string) *ModelRouterQueryClientDiscountLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) SetSuccess(v bool) *ModelRouterQueryClientDiscountLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBody) Validate() error {
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

type ModelRouterQueryClientDiscountLogsResponseBodyData struct {
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int64 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 0.5
	Discount *float32 `json:"discount,omitempty" xml:"discount,omitempty"`
	// example:
	//
	// 2025-09-01 00:00:00
	EffectiveTime *string `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// example:
	//
	// 2025-09-10 00:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id     *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s ModelRouterQueryClientDiscountLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientDiscountLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetDeleteTag() *int64 {
	return s.DeleteTag
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetDiscount() *float32 {
	return s.Discount
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetClientId(v int64) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetDeleteTag(v int64) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.DeleteTag = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetDiscount(v float32) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.Discount = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetEffectiveTime(v string) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetExpireTime(v string) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetGmtCreate(v string) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetGmtModified(v string) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetId(v int64) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) SetRemark(v string) *ModelRouterQueryClientDiscountLogsResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
