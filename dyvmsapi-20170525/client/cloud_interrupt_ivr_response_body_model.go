// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudInterruptIvrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudInterruptIvrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudInterruptIvrResponseBody
	GetCode() *string
	SetData(v *CloudInterruptIvrResponseBodyData) *CloudInterruptIvrResponseBody
	GetData() *CloudInterruptIvrResponseBodyData
	SetMessage(v string) *CloudInterruptIvrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudInterruptIvrResponseBody
	GetRequestId() *string
}

type CloudInterruptIvrResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudInterruptIvrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudInterruptIvrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudInterruptIvrResponseBody) GoString() string {
	return s.String()
}

func (s *CloudInterruptIvrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudInterruptIvrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudInterruptIvrResponseBody) GetData() *CloudInterruptIvrResponseBodyData {
	return s.Data
}

func (s *CloudInterruptIvrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudInterruptIvrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudInterruptIvrResponseBody) SetAccessDeniedDetail(v string) *CloudInterruptIvrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudInterruptIvrResponseBody) SetCode(v string) *CloudInterruptIvrResponseBody {
	s.Code = &v
	return s
}

func (s *CloudInterruptIvrResponseBody) SetData(v *CloudInterruptIvrResponseBodyData) *CloudInterruptIvrResponseBody {
	s.Data = v
	return s
}

func (s *CloudInterruptIvrResponseBody) SetMessage(v string) *CloudInterruptIvrResponseBody {
	s.Message = &v
	return s
}

func (s *CloudInterruptIvrResponseBody) SetRequestId(v string) *CloudInterruptIvrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudInterruptIvrResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudInterruptIvrResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudInterruptIvrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudInterruptIvrResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudInterruptIvrResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudInterruptIvrResponseBodyData) SetResult(v int64) *CloudInterruptIvrResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudInterruptIvrResponseBodyData) Validate() error {
	return dara.Validate(s)
}
