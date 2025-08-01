// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenersByConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListListenersByConfigResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListListenersByConfigResponseBody
	GetHttpCode() *string
	SetListeners(v []*ListListenersByConfigResponseBodyListeners) *ListListenersByConfigResponseBody
	GetListeners() []*ListListenersByConfigResponseBodyListeners
	SetMessage(v string) *ListListenersByConfigResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListListenersByConfigResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListListenersByConfigResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListListenersByConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListListenersByConfigResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListListenersByConfigResponseBody
	GetTotalCount() *int32
}

type ListListenersByConfigResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The information about listeners.
	Listeners []*ListListenersByConfigResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 75E1442F-11EC-567A-9CF0-5A36F7904F39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersByConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListListenersByConfigResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListListenersByConfigResponseBody) GetListeners() []*ListListenersByConfigResponseBodyListeners {
	return s.Listeners
}

func (s *ListListenersByConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListListenersByConfigResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListListenersByConfigResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListListenersByConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListListenersByConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListListenersByConfigResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListListenersByConfigResponseBody) SetErrorCode(v string) *ListListenersByConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetHttpCode(v string) *ListListenersByConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetListeners(v []*ListListenersByConfigResponseBodyListeners) *ListListenersByConfigResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersByConfigResponseBody) SetMessage(v string) *ListListenersByConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetPageNumber(v int32) *ListListenersByConfigResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetPageSize(v int32) *ListListenersByConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetRequestId(v string) *ListListenersByConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetSuccess(v bool) *ListListenersByConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetTotalCount(v int32) *ListListenersByConfigResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListenersByConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListListenersByConfigResponseBodyListeners struct {
	// The IP address.
	//
	// example:
	//
	// 1.1.1.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The label of the listener.
	Labels        map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	MatchRuleName *string            `json:"MatchRuleName,omitempty" xml:"MatchRuleName,omitempty"`
	MatchRuleType *string            `json:"MatchRuleType,omitempty" xml:"MatchRuleType,omitempty"`
	// The verification string.
	//
	// example:
	//
	// 23sdfdf
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The status.
	//
	// example:
	//
	// beta
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The current version of the listener. Valid values: gray and normal.
	//
	// example:
	//
	// gray
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListListenersByConfigResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s ListListenersByConfigResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigResponseBodyListeners) GetIp() *string {
	return s.Ip
}

func (s *ListListenersByConfigResponseBodyListeners) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListListenersByConfigResponseBodyListeners) GetMatchRuleName() *string {
	return s.MatchRuleName
}

func (s *ListListenersByConfigResponseBodyListeners) GetMatchRuleType() *string {
	return s.MatchRuleType
}

func (s *ListListenersByConfigResponseBodyListeners) GetMd5() *string {
	return s.Md5
}

func (s *ListListenersByConfigResponseBodyListeners) GetStatus() *string {
	return s.Status
}

func (s *ListListenersByConfigResponseBodyListeners) GetVersion() *string {
	return s.Version
}

func (s *ListListenersByConfigResponseBodyListeners) SetIp(v string) *ListListenersByConfigResponseBodyListeners {
	s.Ip = &v
	return s
}

func (s *ListListenersByConfigResponseBodyListeners) SetLabels(v map[string]*string) *ListListenersByConfigResponseBodyListeners {
	s.Labels = v
	return s
}

func (s *ListListenersByConfigResponseBodyListeners) SetMatchRuleName(v string) *ListListenersByConfigResponseBodyListeners {
	s.MatchRuleName = &v
	return s
}

func (s *ListListenersByConfigResponseBodyListeners) SetMatchRuleType(v string) *ListListenersByConfigResponseBodyListeners {
	s.MatchRuleType = &v
	return s
}

func (s *ListListenersByConfigResponseBodyListeners) SetMd5(v string) *ListListenersByConfigResponseBodyListeners {
	s.Md5 = &v
	return s
}

func (s *ListListenersByConfigResponseBodyListeners) SetStatus(v string) *ListListenersByConfigResponseBodyListeners {
	s.Status = &v
	return s
}

func (s *ListListenersByConfigResponseBodyListeners) SetVersion(v string) *ListListenersByConfigResponseBodyListeners {
	s.Version = &v
	return s
}

func (s *ListListenersByConfigResponseBodyListeners) Validate() error {
	return dara.Validate(s)
}
