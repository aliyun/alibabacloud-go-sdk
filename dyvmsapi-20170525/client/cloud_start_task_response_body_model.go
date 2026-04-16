// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudStartTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudStartTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudStartTaskResponseBody
	GetCode() *string
	SetData(v *CloudStartTaskResponseBodyData) *CloudStartTaskResponseBody
	GetData() *CloudStartTaskResponseBodyData
	SetMessage(v string) *CloudStartTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudStartTaskResponseBody
	GetRequestId() *string
}

type CloudStartTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudStartTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudStartTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudStartTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloudStartTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudStartTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudStartTaskResponseBody) GetData() *CloudStartTaskResponseBodyData {
	return s.Data
}

func (s *CloudStartTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudStartTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudStartTaskResponseBody) SetAccessDeniedDetail(v string) *CloudStartTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudStartTaskResponseBody) SetCode(v string) *CloudStartTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CloudStartTaskResponseBody) SetData(v *CloudStartTaskResponseBodyData) *CloudStartTaskResponseBody {
	s.Data = v
	return s
}

func (s *CloudStartTaskResponseBody) SetMessage(v string) *CloudStartTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CloudStartTaskResponseBody) SetRequestId(v string) *CloudStartTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudStartTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudStartTaskResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudStartTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudStartTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudStartTaskResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudStartTaskResponseBodyData) SetResult(v int64) *CloudStartTaskResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudStartTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
