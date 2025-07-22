// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPfsMetricTrendsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetPfsMetricTrendsResponseBody
	GetCode() *int64
	SetData(v map[string][]*DataValue) *GetPfsMetricTrendsResponseBody
	GetData() map[string][]*DataValue
	SetMessage(v string) *GetPfsMetricTrendsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPfsMetricTrendsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPfsMetricTrendsResponseBody
	GetSuccess() *bool
}

type GetPfsMetricTrendsResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data map[string][]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F4C27966-284E-51C4-9407-DB50CAB8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPfsMetricTrendsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPfsMetricTrendsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPfsMetricTrendsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetPfsMetricTrendsResponseBody) GetData() map[string][]*DataValue {
	return s.Data
}

func (s *GetPfsMetricTrendsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPfsMetricTrendsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPfsMetricTrendsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPfsMetricTrendsResponseBody) SetCode(v int64) *GetPfsMetricTrendsResponseBody {
	s.Code = &v
	return s
}

func (s *GetPfsMetricTrendsResponseBody) SetData(v map[string][]*DataValue) *GetPfsMetricTrendsResponseBody {
	s.Data = v
	return s
}

func (s *GetPfsMetricTrendsResponseBody) SetMessage(v string) *GetPfsMetricTrendsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPfsMetricTrendsResponseBody) SetRequestId(v string) *GetPfsMetricTrendsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPfsMetricTrendsResponseBody) SetSuccess(v bool) *GetPfsMetricTrendsResponseBody {
	s.Success = &v
	return s
}

func (s *GetPfsMetricTrendsResponseBody) Validate() error {
	return dara.Validate(s)
}
