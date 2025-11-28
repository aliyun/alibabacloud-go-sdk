// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackgeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePackgeInfoResponseBody
	GetCode() *string
	SetData(v *DescribePackgeInfoResponseBodyData) *DescribePackgeInfoResponseBody
	GetData() *DescribePackgeInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribePackgeInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribePackgeInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePackgeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePackgeInfoResponseBody
	GetSuccess() *bool
}

type DescribePackgeInfoResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribePackgeInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePackgeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePackgeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackgeInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePackgeInfoResponseBody) GetData() *DescribePackgeInfoResponseBodyData {
	return s.Data
}

func (s *DescribePackgeInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribePackgeInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePackgeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePackgeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePackgeInfoResponseBody) SetCode(v string) *DescribePackgeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePackgeInfoResponseBody) SetData(v *DescribePackgeInfoResponseBodyData) *DescribePackgeInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribePackgeInfoResponseBody) SetHttpStatusCode(v int32) *DescribePackgeInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribePackgeInfoResponseBody) SetMessage(v string) *DescribePackgeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePackgeInfoResponseBody) SetRequestId(v string) *DescribePackgeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackgeInfoResponseBody) SetSuccess(v bool) *DescribePackgeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePackgeInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePackgeInfoResponseBodyData struct {
	EnableTrace *bool  `json:"EnableTrace,omitempty" xml:"EnableTrace,omitempty"`
	EndTime     *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime   *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePackgeInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePackgeInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePackgeInfoResponseBodyData) GetEnableTrace() *bool {
	return s.EnableTrace
}

func (s *DescribePackgeInfoResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePackgeInfoResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePackgeInfoResponseBodyData) SetEnableTrace(v bool) *DescribePackgeInfoResponseBodyData {
	s.EnableTrace = &v
	return s
}

func (s *DescribePackgeInfoResponseBodyData) SetEndTime(v int64) *DescribePackgeInfoResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribePackgeInfoResponseBodyData) SetStartTime(v int64) *DescribePackgeInfoResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribePackgeInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
