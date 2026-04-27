// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPreviewoutcallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudPreviewoutcallResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudPreviewoutcallResponseBody
	GetCode() *string
	SetData(v *CloudPreviewoutcallResponseBodyData) *CloudPreviewoutcallResponseBody
	GetData() *CloudPreviewoutcallResponseBodyData
	SetMessage(v string) *CloudPreviewoutcallResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudPreviewoutcallResponseBody
	GetRequestId() *string
}

type CloudPreviewoutcallResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudPreviewoutcallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudPreviewoutcallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudPreviewoutcallResponseBody) GoString() string {
	return s.String()
}

func (s *CloudPreviewoutcallResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudPreviewoutcallResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudPreviewoutcallResponseBody) GetData() *CloudPreviewoutcallResponseBodyData {
	return s.Data
}

func (s *CloudPreviewoutcallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudPreviewoutcallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudPreviewoutcallResponseBody) SetAccessDeniedDetail(v string) *CloudPreviewoutcallResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudPreviewoutcallResponseBody) SetCode(v string) *CloudPreviewoutcallResponseBody {
	s.Code = &v
	return s
}

func (s *CloudPreviewoutcallResponseBody) SetData(v *CloudPreviewoutcallResponseBodyData) *CloudPreviewoutcallResponseBody {
	s.Data = v
	return s
}

func (s *CloudPreviewoutcallResponseBody) SetMessage(v string) *CloudPreviewoutcallResponseBody {
	s.Message = &v
	return s
}

func (s *CloudPreviewoutcallResponseBody) SetRequestId(v string) *CloudPreviewoutcallResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudPreviewoutcallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudPreviewoutcallResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudPreviewoutcallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudPreviewoutcallResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudPreviewoutcallResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudPreviewoutcallResponseBodyData) SetResult(v int64) *CloudPreviewoutcallResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudPreviewoutcallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
