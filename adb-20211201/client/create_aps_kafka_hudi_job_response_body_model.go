// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsKafkaHudiJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApsKafkaHudiJobResponseBody
	GetCode() *string
	SetData(v string) *CreateApsKafkaHudiJobResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateApsKafkaHudiJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApsKafkaHudiJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApsKafkaHudiJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApsKafkaHudiJobResponseBody
	GetSuccess() *bool
}

type CreateApsKafkaHudiJobResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1A943417-5B0E-1DB9-A8**-A566****C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateApsKafkaHudiJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApsKafkaHudiJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApsKafkaHudiJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApsKafkaHudiJobResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateApsKafkaHudiJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApsKafkaHudiJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApsKafkaHudiJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApsKafkaHudiJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApsKafkaHudiJobResponseBody) SetCode(v string) *CreateApsKafkaHudiJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApsKafkaHudiJobResponseBody) SetData(v string) *CreateApsKafkaHudiJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApsKafkaHudiJobResponseBody) SetHttpStatusCode(v int32) *CreateApsKafkaHudiJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApsKafkaHudiJobResponseBody) SetMessage(v string) *CreateApsKafkaHudiJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApsKafkaHudiJobResponseBody) SetRequestId(v string) *CreateApsKafkaHudiJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApsKafkaHudiJobResponseBody) SetSuccess(v bool) *CreateApsKafkaHudiJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApsKafkaHudiJobResponseBody) Validate() error {
	return dara.Validate(s)
}
