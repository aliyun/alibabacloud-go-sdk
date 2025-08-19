// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAppQuotaResponseBody
	GetRequestId() *string
	SetResult(v *GetAppQuotaResponseBodyResult) *GetAppQuotaResponseBody
	GetResult() *GetAppQuotaResponseBodyResult
}

type GetAppQuotaResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetAppQuotaResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAppQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppQuotaResponseBody) GetResult() *GetAppQuotaResponseBodyResult {
	return s.Result
}

func (s *GetAppQuotaResponseBody) SetRequestId(v string) *GetAppQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppQuotaResponseBody) SetResult(v *GetAppQuotaResponseBodyResult) *GetAppQuotaResponseBody {
	s.Result = v
	return s
}

func (s *GetAppQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAppQuotaResponseBodyResult struct {
	LimiterInfo *GetAppQuotaResponseBodyResultLimiterInfo `json:"limiterInfo,omitempty" xml:"limiterInfo,omitempty" type:"Struct"`
	QuotaInfo   map[string]interface{}                    `json:"quotaInfo,omitempty" xml:"quotaInfo,omitempty"`
}

func (s GetAppQuotaResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAppQuotaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponseBodyResult) GetLimiterInfo() *GetAppQuotaResponseBodyResultLimiterInfo {
	return s.LimiterInfo
}

func (s *GetAppQuotaResponseBodyResult) GetQuotaInfo() map[string]interface{} {
	return s.QuotaInfo
}

func (s *GetAppQuotaResponseBodyResult) SetLimiterInfo(v *GetAppQuotaResponseBodyResultLimiterInfo) *GetAppQuotaResponseBodyResult {
	s.LimiterInfo = v
	return s
}

func (s *GetAppQuotaResponseBodyResult) SetQuotaInfo(v map[string]interface{}) *GetAppQuotaResponseBodyResult {
	s.QuotaInfo = v
	return s
}

func (s *GetAppQuotaResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetAppQuotaResponseBodyResultLimiterInfo struct {
	Limiters []*GetAppQuotaResponseBodyResultLimiterInfoLimiters `json:"limiters,omitempty" xml:"limiters,omitempty" type:"Repeated"`
}

func (s GetAppQuotaResponseBodyResultLimiterInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAppQuotaResponseBodyResultLimiterInfo) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponseBodyResultLimiterInfo) GetLimiters() []*GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	return s.Limiters
}

func (s *GetAppQuotaResponseBodyResultLimiterInfo) SetLimiters(v []*GetAppQuotaResponseBodyResultLimiterInfoLimiters) *GetAppQuotaResponseBodyResultLimiterInfo {
	s.Limiters = v
	return s
}

func (s *GetAppQuotaResponseBodyResultLimiterInfo) Validate() error {
	return dara.Validate(s)
}

type GetAppQuotaResponseBodyResultLimiterInfoLimiters struct {
	// example:
	//
	// true
	Immutable *bool `json:"immutable,omitempty" xml:"immutable,omitempty"`
	// example:
	//
	// 10
	MaxValue *int64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	// example:
	//
	// 1
	MinValue *int64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
	// example:
	//
	// INDEX_NUMBER_OF_SHARDS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAppQuotaResponseBodyResultLimiterInfoLimiters) String() string {
	return dara.Prettify(s)
}

func (s GetAppQuotaResponseBodyResultLimiterInfoLimiters) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) GetImmutable() *bool {
	return s.Immutable
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) GetMaxValue() *int64 {
	return s.MaxValue
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) GetMinValue() *int64 {
	return s.MinValue
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) GetType() *string {
	return s.Type
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) SetImmutable(v bool) *GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	s.Immutable = &v
	return s
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) SetMaxValue(v int64) *GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	s.MaxValue = &v
	return s
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) SetMinValue(v int64) *GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	s.MinValue = &v
	return s
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) SetType(v string) *GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	s.Type = &v
	return s
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) Validate() error {
	return dara.Validate(s)
}
