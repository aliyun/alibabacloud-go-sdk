// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigQueryEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIConfigQueryEntriesResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIConfigQueryEntriesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIConfigQueryEntriesResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIConfigQueryEntriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIConfigQueryEntriesResponseBody
	GetSuccess() *bool
}

type ChatBIConfigQueryEntriesResponseBody struct {
	// example:
	//
	// JSONObject，其中"entries"对应查询表项，List<JSONObject>格式
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 具体错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 具体错误信息
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatBIConfigQueryEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigQueryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIConfigQueryEntriesResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIConfigQueryEntriesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIConfigQueryEntriesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIConfigQueryEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIConfigQueryEntriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIConfigQueryEntriesResponseBody) SetData(v interface{}) *ChatBIConfigQueryEntriesResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIConfigQueryEntriesResponseBody) SetErrCode(v string) *ChatBIConfigQueryEntriesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIConfigQueryEntriesResponseBody) SetErrMessage(v string) *ChatBIConfigQueryEntriesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIConfigQueryEntriesResponseBody) SetRequestId(v string) *ChatBIConfigQueryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIConfigQueryEntriesResponseBody) SetSuccess(v bool) *ChatBIConfigQueryEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIConfigQueryEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
