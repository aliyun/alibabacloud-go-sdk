// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgendaSumRecordFlowPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddAgendaSumRecordFlowPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v bool) *AddAgendaSumRecordFlowPopResponseBody
	GetData() *bool
	SetErrCode(v string) *AddAgendaSumRecordFlowPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AddAgendaSumRecordFlowPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AddAgendaSumRecordFlowPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AddAgendaSumRecordFlowPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAgendaSumRecordFlowPopResponseBody
	GetSuccess() *bool
}

type AddAgendaSumRecordFlowPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 502
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// param error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// request-id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAgendaSumRecordFlowPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAgendaSumRecordFlowPopResponseBody) GoString() string {
	return s.String()
}

func (s *AddAgendaSumRecordFlowPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddAgendaSumRecordFlowPopResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddAgendaSumRecordFlowPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AddAgendaSumRecordFlowPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AddAgendaSumRecordFlowPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddAgendaSumRecordFlowPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAgendaSumRecordFlowPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAgendaSumRecordFlowPopResponseBody) SetAccessDeniedDetail(v string) *AddAgendaSumRecordFlowPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponseBody) SetData(v bool) *AddAgendaSumRecordFlowPopResponseBody {
	s.Data = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponseBody) SetErrCode(v string) *AddAgendaSumRecordFlowPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponseBody) SetErrMessage(v string) *AddAgendaSumRecordFlowPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponseBody) SetHttpStatusCode(v int32) *AddAgendaSumRecordFlowPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponseBody) SetRequestId(v string) *AddAgendaSumRecordFlowPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponseBody) SetSuccess(v bool) *AddAgendaSumRecordFlowPopResponseBody {
	s.Success = &v
	return s
}

func (s *AddAgendaSumRecordFlowPopResponseBody) Validate() error {
	return dara.Validate(s)
}
