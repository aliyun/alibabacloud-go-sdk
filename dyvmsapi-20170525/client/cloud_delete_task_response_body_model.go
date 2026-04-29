// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteTaskResponseBody
	GetCode() *string
	SetData(v *CloudDeleteTaskResponseBodyData) *CloudDeleteTaskResponseBody
	GetData() *CloudDeleteTaskResponseBodyData
	SetMessage(v string) *CloudDeleteTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteTaskResponseBody
	GetRequestId() *string
}

type CloudDeleteTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteTaskResponseBody) GetData() *CloudDeleteTaskResponseBodyData {
	return s.Data
}

func (s *CloudDeleteTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteTaskResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteTaskResponseBody) SetCode(v string) *CloudDeleteTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteTaskResponseBody) SetData(v *CloudDeleteTaskResponseBodyData) *CloudDeleteTaskResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteTaskResponseBody) SetMessage(v string) *CloudDeleteTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteTaskResponseBody) SetRequestId(v string) *CloudDeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteTaskResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteTaskResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteTaskResponseBodyData) SetResult(v int64) *CloudDeleteTaskResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
