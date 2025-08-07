// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadNumGroupTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadNumGroupTotalResponseBody
	GetCode() *string
	SetData(v []*ReadNumGroupTotalResponseBodyData) *ReadNumGroupTotalResponseBody
	GetData() []*ReadNumGroupTotalResponseBodyData
	SetMessage(v string) *ReadNumGroupTotalResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadNumGroupTotalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadNumGroupTotalResponseBody
	GetSuccess() *bool
}

type ReadNumGroupTotalResponseBody struct {
	// example:
	//
	// SUCCESS
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ReadNumGroupTotalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadNumGroupTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadNumGroupTotalResponseBody) GoString() string {
	return s.String()
}

func (s *ReadNumGroupTotalResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadNumGroupTotalResponseBody) GetData() []*ReadNumGroupTotalResponseBodyData {
	return s.Data
}

func (s *ReadNumGroupTotalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadNumGroupTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadNumGroupTotalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadNumGroupTotalResponseBody) SetCode(v string) *ReadNumGroupTotalResponseBody {
	s.Code = &v
	return s
}

func (s *ReadNumGroupTotalResponseBody) SetData(v []*ReadNumGroupTotalResponseBodyData) *ReadNumGroupTotalResponseBody {
	s.Data = v
	return s
}

func (s *ReadNumGroupTotalResponseBody) SetMessage(v string) *ReadNumGroupTotalResponseBody {
	s.Message = &v
	return s
}

func (s *ReadNumGroupTotalResponseBody) SetRequestId(v string) *ReadNumGroupTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadNumGroupTotalResponseBody) SetSuccess(v bool) *ReadNumGroupTotalResponseBody {
	s.Success = &v
	return s
}

func (s *ReadNumGroupTotalResponseBody) Validate() error {
	return dara.Validate(s)
}

type ReadNumGroupTotalResponseBodyData struct {
	Id          *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	ReadCount   *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UnReadCount *int64 `json:"UnReadCount,omitempty" xml:"UnReadCount,omitempty"`
}

func (s ReadNumGroupTotalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadNumGroupTotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadNumGroupTotalResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ReadNumGroupTotalResponseBodyData) GetReadCount() *int64 {
	return s.ReadCount
}

func (s *ReadNumGroupTotalResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ReadNumGroupTotalResponseBodyData) GetUnReadCount() *int64 {
	return s.UnReadCount
}

func (s *ReadNumGroupTotalResponseBodyData) SetId(v int64) *ReadNumGroupTotalResponseBodyData {
	s.Id = &v
	return s
}

func (s *ReadNumGroupTotalResponseBodyData) SetReadCount(v int64) *ReadNumGroupTotalResponseBodyData {
	s.ReadCount = &v
	return s
}

func (s *ReadNumGroupTotalResponseBodyData) SetTotalCount(v int64) *ReadNumGroupTotalResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ReadNumGroupTotalResponseBodyData) SetUnReadCount(v int64) *ReadNumGroupTotalResponseBodyData {
	s.UnReadCount = &v
	return s
}

func (s *ReadNumGroupTotalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
