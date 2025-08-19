// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpecReviewTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSpecReviewTaskResponseBody
	GetRequestId() *string
	SetResult(v *GetSpecReviewTaskResponseBodyResult) *GetSpecReviewTaskResponseBody
	GetResult() *GetSpecReviewTaskResponseBodyResult
}

type GetSpecReviewTaskResponseBody struct {
	// example:
	//
	// E310AC54-957A-5FD5-B85B-E972B2BDA8DE
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetSpecReviewTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSpecReviewTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSpecReviewTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpecReviewTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSpecReviewTaskResponseBody) GetResult() *GetSpecReviewTaskResponseBodyResult {
	return s.Result
}

func (s *GetSpecReviewTaskResponseBody) SetRequestId(v string) *GetSpecReviewTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpecReviewTaskResponseBody) SetResult(v *GetSpecReviewTaskResponseBodyResult) *GetSpecReviewTaskResponseBody {
	s.Result = v
	return s
}

func (s *GetSpecReviewTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSpecReviewTaskResponseBodyResult struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 339
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// {
	//
	//                "limiters": [
	//
	//                     {
	//
	//                          "type": "INDEX_QUOTA",
	//
	//                          "maxValue": 500,
	//
	//                          "immutable": false
	//
	//                     }
	//
	//                ]
	//
	//           }
	ApplyLimiter map[string]interface{} `json:"applyLimiter,omitempty" xml:"applyLimiter,omitempty"`
	// example:
	//
	// {
	//
	//                "appType": "TRIAL",
	//
	//                "cu": 4,
	//
	//                "storage": 100
	//
	//           }
	ApplyQuota  map[string]interface{} `json:"applyQuota,omitempty" xml:"applyQuota,omitempty"`
	ApplyReason *string                `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	// example:
	//
	// {
	//
	//                "limiters": [
	//
	//                     {
	//
	//                          "type": "INDEX_QUOTA",
	//
	//                          "maxValue": 500,
	//
	//                          "immutable": false
	//
	//                     }
	//
	//                ]
	//
	//           }
	EffectiveLimiter map[string]interface{} `json:"effectiveLimiter,omitempty" xml:"effectiveLimiter,omitempty"`
	// example:
	//
	// {
	//
	//                "appType": "TRIAL",
	//
	//                "cu": 4,
	//
	//                "storage": 100
	//
	//           }
	EffectiveQuota map[string]interface{} `json:"effectiveQuota,omitempty" xml:"effectiveQuota,omitempty"`
	// example:
	//
	// 2024-05-30T06:28:07.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-05-30T06:28:07.000Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// {
	//
	//                "limiters": [
	//
	//                     {
	//
	//                          "type": "INDEX_QUOTA",
	//
	//                          "maxValue": 500,
	//
	//                          "immutable": false
	//
	//                     }
	//
	//                ]
	//
	//           }
	OldLimiter map[string]interface{} `json:"oldLimiter,omitempty" xml:"oldLimiter,omitempty"`
	// example:
	//
	// {
	//
	//                "appType": "TRIAL",
	//
	//                "cu": 2,
	//
	//                "storage": 1
	//
	//           }
	OldQuota map[string]interface{} `json:"oldQuota,omitempty" xml:"oldQuota,omitempty"`
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// QUOTA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSpecReviewTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSpecReviewTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSpecReviewTaskResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetSpecReviewTaskResponseBodyResult) GetAppName() *string {
	return s.AppName
}

func (s *GetSpecReviewTaskResponseBodyResult) GetApplyLimiter() map[string]interface{} {
	return s.ApplyLimiter
}

func (s *GetSpecReviewTaskResponseBodyResult) GetApplyQuota() map[string]interface{} {
	return s.ApplyQuota
}

func (s *GetSpecReviewTaskResponseBodyResult) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *GetSpecReviewTaskResponseBodyResult) GetEffectiveLimiter() map[string]interface{} {
	return s.EffectiveLimiter
}

func (s *GetSpecReviewTaskResponseBodyResult) GetEffectiveQuota() map[string]interface{} {
	return s.EffectiveQuota
}

func (s *GetSpecReviewTaskResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetSpecReviewTaskResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetSpecReviewTaskResponseBodyResult) GetOldLimiter() map[string]interface{} {
	return s.OldLimiter
}

func (s *GetSpecReviewTaskResponseBodyResult) GetOldQuota() map[string]interface{} {
	return s.OldQuota
}

func (s *GetSpecReviewTaskResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *GetSpecReviewTaskResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetSpecReviewTaskResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetSpecReviewTaskResponseBodyResult) SetId(v string) *GetSpecReviewTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetAppName(v string) *GetSpecReviewTaskResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetApplyLimiter(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.ApplyLimiter = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetApplyQuota(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.ApplyQuota = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetApplyReason(v string) *GetSpecReviewTaskResponseBodyResult {
	s.ApplyReason = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetEffectiveLimiter(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.EffectiveLimiter = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetEffectiveQuota(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.EffectiveQuota = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetGmtCreate(v string) *GetSpecReviewTaskResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetGmtModified(v string) *GetSpecReviewTaskResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetOldLimiter(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.OldLimiter = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetOldQuota(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.OldQuota = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetSource(v string) *GetSpecReviewTaskResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetStatus(v string) *GetSpecReviewTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetType(v string) *GetSpecReviewTaskResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
