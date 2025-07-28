// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveUrgentDemandItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SaveUrgentDemandItemResponseBody
	GetCode() *int64
	SetData(v []*int64) *SaveUrgentDemandItemResponseBody
	GetData() []*int64
	SetMessage(v string) *SaveUrgentDemandItemResponseBody
	GetMessage() *string
	SetSuccess(v bool) *SaveUrgentDemandItemResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *SaveUrgentDemandItemResponseBody
	GetTraceId() *string
}

type SaveUrgentDemandItemResponseBody struct {
	// example:
	//
	// 0
	Code *int64   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*int64 `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2107d95616405752026995105e83b0
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s SaveUrgentDemandItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveUrgentDemandItemResponseBody) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SaveUrgentDemandItemResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *SaveUrgentDemandItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveUrgentDemandItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveUrgentDemandItemResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *SaveUrgentDemandItemResponseBody) SetCode(v int64) *SaveUrgentDemandItemResponseBody {
	s.Code = &v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) SetData(v []*int64) *SaveUrgentDemandItemResponseBody {
	s.Data = v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) SetMessage(v string) *SaveUrgentDemandItemResponseBody {
	s.Message = &v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) SetSuccess(v bool) *SaveUrgentDemandItemResponseBody {
	s.Success = &v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) SetTraceId(v string) *SaveUrgentDemandItemResponseBody {
	s.TraceId = &v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) Validate() error {
	return dara.Validate(s)
}
