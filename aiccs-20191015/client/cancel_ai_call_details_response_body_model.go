// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAiCallDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CancelAiCallDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CancelAiCallDetailsResponseBody
	GetCode() *string
	SetData(v *CancelAiCallDetailsResponseBodyData) *CancelAiCallDetailsResponseBody
	GetData() *CancelAiCallDetailsResponseBodyData
	SetMessage(v string) *CancelAiCallDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelAiCallDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelAiCallDetailsResponseBody
	GetSuccess() *bool
}

type CancelAiCallDetailsResponseBody struct {
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CancelAiCallDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 46C98E28-9239-5D95-AC76-648B8FD4889A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAiCallDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAiCallDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAiCallDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CancelAiCallDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelAiCallDetailsResponseBody) GetData() *CancelAiCallDetailsResponseBodyData {
	return s.Data
}

func (s *CancelAiCallDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelAiCallDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAiCallDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelAiCallDetailsResponseBody) SetAccessDeniedDetail(v string) *CancelAiCallDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CancelAiCallDetailsResponseBody) SetCode(v string) *CancelAiCallDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAiCallDetailsResponseBody) SetData(v *CancelAiCallDetailsResponseBodyData) *CancelAiCallDetailsResponseBody {
	s.Data = v
	return s
}

func (s *CancelAiCallDetailsResponseBody) SetMessage(v string) *CancelAiCallDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAiCallDetailsResponseBody) SetRequestId(v string) *CancelAiCallDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAiCallDetailsResponseBody) SetSuccess(v bool) *CancelAiCallDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *CancelAiCallDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelAiCallDetailsResponseBodyData struct {
	// example:
	//
	// 75
	FailedCount   *int64                 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	FailedDetails map[string]interface{} `json:"FailedDetails,omitempty" xml:"FailedDetails,omitempty"`
	// example:
	//
	// ALL_SUCCEED
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// 81
	SucceedCount *int64 `json:"SucceedCount,omitempty" xml:"SucceedCount,omitempty"`
	// example:
	//
	// 50
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CancelAiCallDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelAiCallDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelAiCallDetailsResponseBodyData) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *CancelAiCallDetailsResponseBodyData) GetFailedDetails() map[string]interface{} {
	return s.FailedDetails
}

func (s *CancelAiCallDetailsResponseBodyData) GetResultCode() *string {
	return s.ResultCode
}

func (s *CancelAiCallDetailsResponseBodyData) GetSucceedCount() *int64 {
	return s.SucceedCount
}

func (s *CancelAiCallDetailsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *CancelAiCallDetailsResponseBodyData) SetFailedCount(v int64) *CancelAiCallDetailsResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *CancelAiCallDetailsResponseBodyData) SetFailedDetails(v map[string]interface{}) *CancelAiCallDetailsResponseBodyData {
	s.FailedDetails = v
	return s
}

func (s *CancelAiCallDetailsResponseBodyData) SetResultCode(v string) *CancelAiCallDetailsResponseBodyData {
	s.ResultCode = &v
	return s
}

func (s *CancelAiCallDetailsResponseBodyData) SetSucceedCount(v int64) *CancelAiCallDetailsResponseBodyData {
	s.SucceedCount = &v
	return s
}

func (s *CancelAiCallDetailsResponseBodyData) SetTotalCount(v int64) *CancelAiCallDetailsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *CancelAiCallDetailsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
