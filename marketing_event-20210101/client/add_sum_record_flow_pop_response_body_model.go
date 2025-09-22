// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSumRecordFlowPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddSumRecordFlowPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *AddSumRecordFlowPopResponseBody
	GetData() *bool
	SetErrCode(v string) *AddSumRecordFlowPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AddSumRecordFlowPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AddSumRecordFlowPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AddSumRecordFlowPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddSumRecordFlowPopResponseBody
	GetSuccess() *bool
}

type AddSumRecordFlowPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSumRecordFlowPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSumRecordFlowPopResponseBody) GoString() string {
	return s.String()
}

func (s *AddSumRecordFlowPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddSumRecordFlowPopResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddSumRecordFlowPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AddSumRecordFlowPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AddSumRecordFlowPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddSumRecordFlowPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSumRecordFlowPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddSumRecordFlowPopResponseBody) SetAccessDeniedDetail(v string) *AddSumRecordFlowPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetData(v bool) *AddSumRecordFlowPopResponseBody {
	s.Data = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetErrCode(v string) *AddSumRecordFlowPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetErrMessage(v string) *AddSumRecordFlowPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetHttpStatusCode(v int32) *AddSumRecordFlowPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetRequestId(v string) *AddSumRecordFlowPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) SetSuccess(v bool) *AddSumRecordFlowPopResponseBody {
	s.Success = &v
	return s
}

func (s *AddSumRecordFlowPopResponseBody) Validate() error {
	return dara.Validate(s)
}
