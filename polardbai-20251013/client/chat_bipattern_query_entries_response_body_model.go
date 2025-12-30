// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternQueryEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternQueryEntriesResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternQueryEntriesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternQueryEntriesResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternQueryEntriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternQueryEntriesResponseBody
	GetSuccess() *bool
}

type ChatBIPatternQueryEntriesResponseBody struct {
	// example:
	//
	// JSONObject，其中"entries"对应查询表项，List<JSONObject>格式
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 错误信息
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

func (s ChatBIPatternQueryEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternQueryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternQueryEntriesResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternQueryEntriesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternQueryEntriesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternQueryEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternQueryEntriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternQueryEntriesResponseBody) SetData(v interface{}) *ChatBIPatternQueryEntriesResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternQueryEntriesResponseBody) SetErrCode(v string) *ChatBIPatternQueryEntriesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternQueryEntriesResponseBody) SetErrMessage(v string) *ChatBIPatternQueryEntriesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternQueryEntriesResponseBody) SetRequestId(v string) *ChatBIPatternQueryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternQueryEntriesResponseBody) SetSuccess(v bool) *ChatBIPatternQueryEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternQueryEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
