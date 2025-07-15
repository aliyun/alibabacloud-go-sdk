// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStyleLearningResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStyleLearningResultResponseBody
	GetCode() *string
	SetData(v bool) *DeleteStyleLearningResultResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteStyleLearningResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteStyleLearningResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStyleLearningResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStyleLearningResultResponseBody
	GetSuccess() *bool
}

type DeleteStyleLearningResultResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxx
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

func (s DeleteStyleLearningResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStyleLearningResultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStyleLearningResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStyleLearningResultResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteStyleLearningResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteStyleLearningResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStyleLearningResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStyleLearningResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStyleLearningResultResponseBody) SetCode(v string) *DeleteStyleLearningResultResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStyleLearningResultResponseBody) SetData(v bool) *DeleteStyleLearningResultResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteStyleLearningResultResponseBody) SetHttpStatusCode(v int32) *DeleteStyleLearningResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteStyleLearningResultResponseBody) SetMessage(v string) *DeleteStyleLearningResultResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStyleLearningResultResponseBody) SetRequestId(v string) *DeleteStyleLearningResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStyleLearningResultResponseBody) SetSuccess(v bool) *DeleteStyleLearningResultResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStyleLearningResultResponseBody) Validate() error {
	return dara.Validate(s)
}
