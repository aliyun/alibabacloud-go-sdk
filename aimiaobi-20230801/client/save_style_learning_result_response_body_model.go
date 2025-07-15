// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveStyleLearningResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveStyleLearningResultResponseBody
	GetCode() *string
	SetData(v bool) *SaveStyleLearningResultResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *SaveStyleLearningResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveStyleLearningResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveStyleLearningResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveStyleLearningResultResponseBody
	GetSuccess() *bool
}

type SaveStyleLearningResultResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveStyleLearningResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveStyleLearningResultResponseBody) GoString() string {
	return s.String()
}

func (s *SaveStyleLearningResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveStyleLearningResultResponseBody) GetData() *bool {
	return s.Data
}

func (s *SaveStyleLearningResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveStyleLearningResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveStyleLearningResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveStyleLearningResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveStyleLearningResultResponseBody) SetCode(v string) *SaveStyleLearningResultResponseBody {
	s.Code = &v
	return s
}

func (s *SaveStyleLearningResultResponseBody) SetData(v bool) *SaveStyleLearningResultResponseBody {
	s.Data = &v
	return s
}

func (s *SaveStyleLearningResultResponseBody) SetHttpStatusCode(v int32) *SaveStyleLearningResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveStyleLearningResultResponseBody) SetMessage(v string) *SaveStyleLearningResultResponseBody {
	s.Message = &v
	return s
}

func (s *SaveStyleLearningResultResponseBody) SetRequestId(v string) *SaveStyleLearningResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveStyleLearningResultResponseBody) SetSuccess(v bool) *SaveStyleLearningResultResponseBody {
	s.Success = &v
	return s
}

func (s *SaveStyleLearningResultResponseBody) Validate() error {
	return dara.Validate(s)
}
