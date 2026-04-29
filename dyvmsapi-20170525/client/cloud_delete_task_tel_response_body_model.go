// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteTaskTelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteTaskTelResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteTaskTelResponseBody
	GetCode() *string
	SetData(v *CloudDeleteTaskTelResponseBodyData) *CloudDeleteTaskTelResponseBody
	GetData() *CloudDeleteTaskTelResponseBodyData
	SetMessage(v string) *CloudDeleteTaskTelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteTaskTelResponseBody
	GetRequestId() *string
}

type CloudDeleteTaskTelResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteTaskTelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteTaskTelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteTaskTelResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteTaskTelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteTaskTelResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteTaskTelResponseBody) GetData() *CloudDeleteTaskTelResponseBodyData {
	return s.Data
}

func (s *CloudDeleteTaskTelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteTaskTelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteTaskTelResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteTaskTelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteTaskTelResponseBody) SetCode(v string) *CloudDeleteTaskTelResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteTaskTelResponseBody) SetData(v *CloudDeleteTaskTelResponseBodyData) *CloudDeleteTaskTelResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteTaskTelResponseBody) SetMessage(v string) *CloudDeleteTaskTelResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteTaskTelResponseBody) SetRequestId(v string) *CloudDeleteTaskTelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteTaskTelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteTaskTelResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteTaskTelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteTaskTelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteTaskTelResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteTaskTelResponseBodyData) SetResult(v int64) *CloudDeleteTaskTelResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteTaskTelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
