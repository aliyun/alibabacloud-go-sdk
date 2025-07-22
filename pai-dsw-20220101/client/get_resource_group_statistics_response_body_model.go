// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourceGroupStatisticsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetResourceGroupStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetResourceGroupStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourceGroupStatisticsResponseBody
	GetRequestId() *string
	SetStatistics(v map[string]map[string]interface{}) *GetResourceGroupStatisticsResponseBody
	GetStatistics() map[string]map[string]interface{}
	SetSuccess(v bool) *GetResourceGroupStatisticsResponseBody
	GetSuccess() *bool
}

type GetResourceGroupStatisticsResponseBody struct {
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

func (s GetResourceGroupStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourceGroupStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetResourceGroupStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourceGroupStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupStatisticsResponseBody) GetStatistics() map[string]map[string]interface{} {
	return s.Statistics
}

func (s *GetResourceGroupStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetResourceGroupStatisticsResponseBody) SetCode(v string) *GetResourceGroupStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetHttpStatusCode(v int32) *GetResourceGroupStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetMessage(v string) *GetResourceGroupStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetRequestId(v string) *GetResourceGroupStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetStatistics(v map[string]map[string]interface{}) *GetResourceGroupStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetSuccess(v bool) *GetResourceGroupStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
