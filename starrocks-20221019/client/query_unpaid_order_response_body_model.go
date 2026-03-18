// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnpaidOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *QueryUnpaidOrderResponseBody
	GetData() *int64
	SetErrCode(v string) *QueryUnpaidOrderResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryUnpaidOrderResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryUnpaidOrderResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryUnpaidOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryUnpaidOrderResponseBody
	GetSuccess() *bool
}

type QueryUnpaidOrderResponseBody struct {
	// example:
	//
	// true
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
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
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUnpaidOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUnpaidOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnpaidOrderResponseBody) GetData() *int64 {
	return s.Data
}

func (s *QueryUnpaidOrderResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryUnpaidOrderResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryUnpaidOrderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryUnpaidOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUnpaidOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUnpaidOrderResponseBody) SetData(v int64) *QueryUnpaidOrderResponseBody {
	s.Data = &v
	return s
}

func (s *QueryUnpaidOrderResponseBody) SetErrCode(v string) *QueryUnpaidOrderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryUnpaidOrderResponseBody) SetErrMessage(v string) *QueryUnpaidOrderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryUnpaidOrderResponseBody) SetHttpStatusCode(v int32) *QueryUnpaidOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryUnpaidOrderResponseBody) SetRequestId(v string) *QueryUnpaidOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUnpaidOrderResponseBody) SetSuccess(v bool) *QueryUnpaidOrderResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUnpaidOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
