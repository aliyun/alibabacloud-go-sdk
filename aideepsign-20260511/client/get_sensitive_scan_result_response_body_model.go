// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveScanResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSensitiveScanResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetSensitiveScanResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSensitiveScanResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSensitiveScanResultResponseBody
	GetRequestId() *string
	SetResult(v *GetSensitiveScanResultResponseBodyResult) *GetSensitiveScanResultResponseBody
	GetResult() *GetSensitiveScanResultResponseBodyResult
	SetStatus(v string) *GetSensitiveScanResultResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetSensitiveScanResultResponseBody
	GetSuccess() *bool
}

type GetSensitiveScanResultResponseBody struct {
	// The business error code. The value "OK" is returned when the request succeeds.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. The value 200 is returned when the request succeeds.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. The value "success" is returned when the request succeeds.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scan result. This parameter is returned only when the status is completed.
	Result *GetSensitiveScanResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The task status. Valid values:
	//
	// - running: The task is running.
	//
	// - completed: The task is completed.
	//
	// - terminated: The task is terminated or failed.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSensitiveScanResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSensitiveScanResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSensitiveScanResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSensitiveScanResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSensitiveScanResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSensitiveScanResultResponseBody) GetResult() *GetSensitiveScanResultResponseBodyResult {
	return s.Result
}

func (s *GetSensitiveScanResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSensitiveScanResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSensitiveScanResultResponseBody) SetCode(v string) *GetSensitiveScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSensitiveScanResultResponseBody) SetHttpStatusCode(v int32) *GetSensitiveScanResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSensitiveScanResultResponseBody) SetMessage(v string) *GetSensitiveScanResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSensitiveScanResultResponseBody) SetRequestId(v string) *GetSensitiveScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSensitiveScanResultResponseBody) SetResult(v *GetSensitiveScanResultResponseBodyResult) *GetSensitiveScanResultResponseBody {
	s.Result = v
	return s
}

func (s *GetSensitiveScanResultResponseBody) SetStatus(v string) *GetSensitiveScanResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetSensitiveScanResultResponseBody) SetSuccess(v bool) *GetSensitiveScanResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetSensitiveScanResultResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSensitiveScanResultResponseBodyResult struct {
	// The name of the scanned object.
	OssObjectDetail *GetSensitiveScanResultResponseBodyResultOssObjectDetail `json:"OssObjectDetail,omitempty" xml:"OssObjectDetail,omitempty" type:"Struct"`
}

func (s GetSensitiveScanResultResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveScanResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSensitiveScanResultResponseBodyResult) GetOssObjectDetail() *GetSensitiveScanResultResponseBodyResultOssObjectDetail {
	return s.OssObjectDetail
}

func (s *GetSensitiveScanResultResponseBodyResult) SetOssObjectDetail(v *GetSensitiveScanResultResponseBodyResultOssObjectDetail) *GetSensitiveScanResultResponseBodyResult {
	s.OssObjectDetail = v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResult) Validate() error {
	if s.OssObjectDetail != nil {
		if err := s.OssObjectDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSensitiveScanResultResponseBodyResultOssObjectDetail struct {
	// The name of the bucket to which the object belongs.
	//
	// example:
	//
	// aideepsign-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The sensitive data category name.
	//
	// example:
	//
	// 个人信息
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The name of the scanned object.
	//
	// example:
	//
	// abc12345.jpg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The overall risk level name.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The list of sensitive data rules that are hit.
	RuleList []*GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s GetSensitiveScanResultResponseBodyResultOssObjectDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveScanResultResponseBodyResultOssObjectDetail) GoString() string {
	return s.String()
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) GetBucketName() *string {
	return s.BucketName
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) GetName() *string {
	return s.Name
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) GetRuleList() []*GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList {
	return s.RuleList
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) SetBucketName(v string) *GetSensitiveScanResultResponseBodyResultOssObjectDetail {
	s.BucketName = &v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) SetCategoryName(v string) *GetSensitiveScanResultResponseBodyResultOssObjectDetail {
	s.CategoryName = &v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) SetName(v string) *GetSensitiveScanResultResponseBodyResultOssObjectDetail {
	s.Name = &v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) SetRiskLevelName(v string) *GetSensitiveScanResultResponseBodyResultOssObjectDetail {
	s.RiskLevelName = &v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) SetRuleList(v []*GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) *GetSensitiveScanResultResponseBodyResultOssObjectDetail {
	s.RuleList = v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetail) Validate() error {
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList struct {
	// The category name of the rule.
	//
	// example:
	//
	// 个人信息
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The number of hits.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The risk level name of the rule.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The rule name.
	//
	// example:
	//
	// 身份证号
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) GoString() string {
	return s.String()
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) GetCount() *int32 {
	return s.Count
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) SetCategoryName(v string) *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList {
	s.CategoryName = &v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) SetCount(v int32) *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList {
	s.Count = &v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) SetRiskLevelName(v string) *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList {
	s.RiskLevelName = &v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) SetRuleName(v string) *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList {
	s.RuleName = &v
	return s
}

func (s *GetSensitiveScanResultResponseBodyResultOssObjectDetailRuleList) Validate() error {
	return dara.Validate(s)
}
