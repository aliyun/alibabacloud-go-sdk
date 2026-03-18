// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootECSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RebootECSResponseBody
	GetData() *bool
	SetErrCode(v string) *RebootECSResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RebootECSResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *RebootECSResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RebootECSResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RebootECSResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *RebootECSResponseBody
	GetTotal() *int32
}

type RebootECSResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RebootECSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootECSResponseBody) GoString() string {
	return s.String()
}

func (s *RebootECSResponseBody) GetData() *bool {
	return s.Data
}

func (s *RebootECSResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RebootECSResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RebootECSResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RebootECSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootECSResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RebootECSResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *RebootECSResponseBody) SetData(v bool) *RebootECSResponseBody {
	s.Data = &v
	return s
}

func (s *RebootECSResponseBody) SetErrCode(v string) *RebootECSResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RebootECSResponseBody) SetErrMessage(v string) *RebootECSResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RebootECSResponseBody) SetHttpStatusCode(v int32) *RebootECSResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RebootECSResponseBody) SetRequestId(v string) *RebootECSResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootECSResponseBody) SetSuccess(v bool) *RebootECSResponseBody {
	s.Success = &v
	return s
}

func (s *RebootECSResponseBody) SetTotal(v int32) *RebootECSResponseBody {
	s.Total = &v
	return s
}

func (s *RebootECSResponseBody) Validate() error {
	return dara.Validate(s)
}
