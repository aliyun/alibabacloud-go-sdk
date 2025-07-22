// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceStatisticsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListInstanceStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstanceStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceStatisticsResponseBody
	GetRequestId() *string
	SetStatistics(v map[string]map[string]interface{}) *ListInstanceStatisticsResponseBody
	GetStatistics() map[string]map[string]interface{}
	SetSuccess(v bool) *ListInstanceStatisticsResponseBody
	GetSuccess() *bool
}

type ListInstanceStatisticsResponseBody struct {
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics map[string]map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstanceStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceStatisticsResponseBody) GetStatistics() map[string]map[string]interface{} {
	return s.Statistics
}

func (s *ListInstanceStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceStatisticsResponseBody) SetCode(v string) *ListInstanceStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetHttpStatusCode(v int32) *ListInstanceStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetMessage(v string) *ListInstanceStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetRequestId(v string) *ListInstanceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetStatistics(v map[string]map[string]interface{}) *ListInstanceStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetSuccess(v bool) *ListInstanceStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
