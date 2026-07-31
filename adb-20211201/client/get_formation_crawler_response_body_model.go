// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormationCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFormationCrawlerResponseBody
	GetCode() *string
	SetData(v string) *GetFormationCrawlerResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetFormationCrawlerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetFormationCrawlerResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFormationCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFormationCrawlerResponseBody
	GetSuccess() *bool
}

type GetFormationCrawlerResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message. OK is returned if the call was successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFormationCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFormationCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormationCrawlerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFormationCrawlerResponseBody) GetData() *string {
	return s.Data
}

func (s *GetFormationCrawlerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFormationCrawlerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFormationCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFormationCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFormationCrawlerResponseBody) SetCode(v string) *GetFormationCrawlerResponseBody {
	s.Code = &v
	return s
}

func (s *GetFormationCrawlerResponseBody) SetData(v string) *GetFormationCrawlerResponseBody {
	s.Data = &v
	return s
}

func (s *GetFormationCrawlerResponseBody) SetHttpStatusCode(v int32) *GetFormationCrawlerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFormationCrawlerResponseBody) SetMessage(v string) *GetFormationCrawlerResponseBody {
	s.Message = &v
	return s
}

func (s *GetFormationCrawlerResponseBody) SetRequestId(v string) *GetFormationCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFormationCrawlerResponseBody) SetSuccess(v bool) *GetFormationCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *GetFormationCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
