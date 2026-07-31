// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFormationCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFormationCrawlerResponseBody
	GetCode() *string
	SetData(v bool) *DeleteFormationCrawlerResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteFormationCrawlerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteFormationCrawlerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFormationCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFormationCrawlerResponseBody
	GetSuccess() *bool
}

type DeleteFormationCrawlerResponseBody struct {
	// The status code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the task is deleted. Valid values:
	//
	// - true: The task is deleted.
	//
	// - false: The task failed to be deleted.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E8DA77FA-FF0F-5516-A551-86C6E2D4BE92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFormationCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormationCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFormationCrawlerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFormationCrawlerResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteFormationCrawlerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFormationCrawlerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFormationCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFormationCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFormationCrawlerResponseBody) SetCode(v string) *DeleteFormationCrawlerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFormationCrawlerResponseBody) SetData(v bool) *DeleteFormationCrawlerResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFormationCrawlerResponseBody) SetHttpStatusCode(v int32) *DeleteFormationCrawlerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFormationCrawlerResponseBody) SetMessage(v string) *DeleteFormationCrawlerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFormationCrawlerResponseBody) SetRequestId(v string) *DeleteFormationCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFormationCrawlerResponseBody) SetSuccess(v bool) *DeleteFormationCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFormationCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
