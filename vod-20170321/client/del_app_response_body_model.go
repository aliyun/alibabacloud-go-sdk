// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DelAppResponseBody
	GetCode() *string
	SetData(v *DelAppResponseBodyData) *DelAppResponseBody
	GetData() *DelAppResponseBodyData
	SetHttpStatusCode(v int32) *DelAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DelAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *DelAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DelAppResponseBody
	GetSuccess() *bool
}

type DelAppResponseBody struct {
	Code           *string                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DelAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelAppResponseBody) GoString() string {
	return s.String()
}

func (s *DelAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *DelAppResponseBody) GetData() *DelAppResponseBodyData {
	return s.Data
}

func (s *DelAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DelAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DelAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DelAppResponseBody) SetCode(v string) *DelAppResponseBody {
	s.Code = &v
	return s
}

func (s *DelAppResponseBody) SetData(v *DelAppResponseBodyData) *DelAppResponseBody {
	s.Data = v
	return s
}

func (s *DelAppResponseBody) SetHttpStatusCode(v int32) *DelAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DelAppResponseBody) SetMessage(v string) *DelAppResponseBody {
	s.Message = &v
	return s
}

func (s *DelAppResponseBody) SetRequestId(v string) *DelAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelAppResponseBody) SetSuccess(v bool) *DelAppResponseBody {
	s.Success = &v
	return s
}

func (s *DelAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type DelAppResponseBodyData struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DelAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DelAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *DelAppResponseBodyData) GetResult() *bool {
	return s.Result
}

func (s *DelAppResponseBodyData) SetResult(v bool) *DelAppResponseBodyData {
	s.Result = &v
	return s
}

func (s *DelAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
