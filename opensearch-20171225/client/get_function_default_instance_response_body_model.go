// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionDefaultInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFunctionDefaultInstanceResponseBody
	GetCode() *string
	SetFunctionName(v string) *GetFunctionDefaultInstanceResponseBody
	GetFunctionName() *string
	SetHttpCode(v int64) *GetFunctionDefaultInstanceResponseBody
	GetHttpCode() *int64
	SetInstanceName(v string) *GetFunctionDefaultInstanceResponseBody
	GetInstanceName() *string
	SetLatency(v int64) *GetFunctionDefaultInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *GetFunctionDefaultInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFunctionDefaultInstanceResponseBody
	GetRequestId() *string
	SetResult(v *GetFunctionDefaultInstanceResponseBodyResult) *GetFunctionDefaultInstanceResponseBody
	GetResult() *GetFunctionDefaultInstanceResponseBodyResult
	SetStatus(v string) *GetFunctionDefaultInstanceResponseBody
	GetStatus() *string
}

type GetFunctionDefaultInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// DefaultInstance.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// cdn_waf
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// sh-bp1oi31w1jn4ctdyv
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The default running time.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// default instance not set.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 062BA91F-A568-5779-8A5B-9E62C9BB3DC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// {\\"Pagination\\": {\\"TotalCount\\": 0, \\"PageNumber\\": 1, \\"PageSize\\": 10}, \\"AntConsortiums\\": []}
	Result *GetFunctionDefaultInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionDefaultInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFunctionDefaultInstanceResponseBody) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetFunctionDefaultInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *GetFunctionDefaultInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetFunctionDefaultInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *GetFunctionDefaultInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFunctionDefaultInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunctionDefaultInstanceResponseBody) GetResult() *GetFunctionDefaultInstanceResponseBodyResult {
	return s.Result
}

func (s *GetFunctionDefaultInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionDefaultInstanceResponseBody) SetCode(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetFunctionName(v string) *GetFunctionDefaultInstanceResponseBody {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetHttpCode(v int64) *GetFunctionDefaultInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetInstanceName(v string) *GetFunctionDefaultInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetLatency(v int64) *GetFunctionDefaultInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetMessage(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetRequestId(v string) *GetFunctionDefaultInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetResult(v *GetFunctionDefaultInstanceResponseBodyResult) *GetFunctionDefaultInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetStatus(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionDefaultInstanceResponseBodyResult struct {
	// The default instance name.
	//
	// example:
	//
	// model1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s GetFunctionDefaultInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponseBodyResult) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetFunctionDefaultInstanceResponseBodyResult) SetInstanceName(v string) *GetFunctionDefaultInstanceResponseBodyResult {
	s.InstanceName = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
