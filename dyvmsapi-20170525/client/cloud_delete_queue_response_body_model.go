// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteQueueResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteQueueResponseBody
	GetCode() *string
	SetData(v *CloudDeleteQueueResponseBodyData) *CloudDeleteQueueResponseBody
	GetData() *CloudDeleteQueueResponseBodyData
	SetMessage(v string) *CloudDeleteQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteQueueResponseBody
	GetRequestId() *string
}

type CloudDeleteQueueResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteQueueResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteQueueResponseBody) GetData() *CloudDeleteQueueResponseBodyData {
	return s.Data
}

func (s *CloudDeleteQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteQueueResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteQueueResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteQueueResponseBody) SetCode(v string) *CloudDeleteQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteQueueResponseBody) SetData(v *CloudDeleteQueueResponseBodyData) *CloudDeleteQueueResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteQueueResponseBody) SetMessage(v string) *CloudDeleteQueueResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteQueueResponseBody) SetRequestId(v string) *CloudDeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteQueueResponseBodyData struct {
	// 0-成功，其它-失败
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteQueueResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteQueueResponseBodyData) SetResult(v int64) *CloudDeleteQueueResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteQueueResponseBodyData) Validate() error {
	return dara.Validate(s)
}
