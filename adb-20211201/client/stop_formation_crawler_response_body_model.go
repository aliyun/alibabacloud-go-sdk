// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopFormationCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopFormationCrawlerResponseBody
	GetCode() *string
	SetData(v string) *StopFormationCrawlerResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *StopFormationCrawlerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopFormationCrawlerResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopFormationCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopFormationCrawlerResponseBody
	GetSuccess() *bool
}

type StopFormationCrawlerResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the successfully created task.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9CCFAAB4-97B7-5800-B9F2-685EB596E3EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopFormationCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopFormationCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *StopFormationCrawlerResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopFormationCrawlerResponseBody) GetData() *string {
	return s.Data
}

func (s *StopFormationCrawlerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopFormationCrawlerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopFormationCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopFormationCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopFormationCrawlerResponseBody) SetCode(v string) *StopFormationCrawlerResponseBody {
	s.Code = &v
	return s
}

func (s *StopFormationCrawlerResponseBody) SetData(v string) *StopFormationCrawlerResponseBody {
	s.Data = &v
	return s
}

func (s *StopFormationCrawlerResponseBody) SetHttpStatusCode(v int32) *StopFormationCrawlerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopFormationCrawlerResponseBody) SetMessage(v string) *StopFormationCrawlerResponseBody {
	s.Message = &v
	return s
}

func (s *StopFormationCrawlerResponseBody) SetRequestId(v string) *StopFormationCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopFormationCrawlerResponseBody) SetSuccess(v bool) *StopFormationCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *StopFormationCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
