// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrgentDemandItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteUrgentDemandItemResponseBody
	GetCode() *int64
	SetData(v int64) *DeleteUrgentDemandItemResponseBody
	GetData() *int64
	SetMessage(v string) *DeleteUrgentDemandItemResponseBody
	GetMessage() *string
	SetSuccess(v bool) *DeleteUrgentDemandItemResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteUrgentDemandItemResponseBody
	GetTraceId() *string
}

type DeleteUrgentDemandItemResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
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
	// 212cf01016405759151137225e83cd
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteUrgentDemandItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrgentDemandItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandItemResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteUrgentDemandItemResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeleteUrgentDemandItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUrgentDemandItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUrgentDemandItemResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteUrgentDemandItemResponseBody) SetCode(v int64) *DeleteUrgentDemandItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) SetData(v int64) *DeleteUrgentDemandItemResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) SetMessage(v string) *DeleteUrgentDemandItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) SetSuccess(v bool) *DeleteUrgentDemandItemResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) SetTraceId(v string) *DeleteUrgentDemandItemResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) Validate() error {
	return dara.Validate(s)
}
