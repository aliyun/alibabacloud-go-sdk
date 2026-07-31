// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormationCrawlerScheduleStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFormationCrawlerScheduleStateResponseBody
	GetCode() *string
	SetData(v bool) *UpdateFormationCrawlerScheduleStateResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateFormationCrawlerScheduleStateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateFormationCrawlerScheduleStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFormationCrawlerScheduleStateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFormationCrawlerScheduleStateResponseBody
	GetSuccess() *bool
}

type UpdateFormationCrawlerScheduleStateResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 642F3512-C628-5D0C-8815-F6670CEA00D4
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

func (s UpdateFormationCrawlerScheduleStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormationCrawlerScheduleStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) SetCode(v string) *UpdateFormationCrawlerScheduleStateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) SetData(v bool) *UpdateFormationCrawlerScheduleStateResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) SetHttpStatusCode(v int32) *UpdateFormationCrawlerScheduleStateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) SetMessage(v string) *UpdateFormationCrawlerScheduleStateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) SetRequestId(v string) *UpdateFormationCrawlerScheduleStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) SetSuccess(v bool) *UpdateFormationCrawlerScheduleStateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFormationCrawlerScheduleStateResponseBody) Validate() error {
	return dara.Validate(s)
}
