// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsolateLeaderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *IsolateLeaderResponseBody
	GetData() *bool
	SetErrCode(v string) *IsolateLeaderResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *IsolateLeaderResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *IsolateLeaderResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *IsolateLeaderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IsolateLeaderResponseBody
	GetSuccess() *bool
}

type IsolateLeaderResponseBody struct {
	// example:
	//
	// 24151320976****
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IsolateLeaderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IsolateLeaderResponseBody) GoString() string {
	return s.String()
}

func (s *IsolateLeaderResponseBody) GetData() *bool {
	return s.Data
}

func (s *IsolateLeaderResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *IsolateLeaderResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *IsolateLeaderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *IsolateLeaderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IsolateLeaderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IsolateLeaderResponseBody) SetData(v bool) *IsolateLeaderResponseBody {
	s.Data = &v
	return s
}

func (s *IsolateLeaderResponseBody) SetErrCode(v string) *IsolateLeaderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *IsolateLeaderResponseBody) SetErrMessage(v string) *IsolateLeaderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *IsolateLeaderResponseBody) SetHttpStatusCode(v int32) *IsolateLeaderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *IsolateLeaderResponseBody) SetRequestId(v string) *IsolateLeaderResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsolateLeaderResponseBody) SetSuccess(v bool) *IsolateLeaderResponseBody {
	s.Success = &v
	return s
}

func (s *IsolateLeaderResponseBody) Validate() error {
	return dara.Validate(s)
}
