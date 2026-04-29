// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateAsrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateAsrResponseBody
	GetCode() *string
	SetData(v *CloudCreateAsrResponseBodyData) *CloudCreateAsrResponseBody
	GetData() *CloudCreateAsrResponseBodyData
	SetMessage(v string) *CloudCreateAsrResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateAsrResponseBody
	GetRequestId() *string
}

type CloudCreateAsrResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateAsrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateAsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAsrResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateAsrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateAsrResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateAsrResponseBody) GetData() *CloudCreateAsrResponseBodyData {
	return s.Data
}

func (s *CloudCreateAsrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateAsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateAsrResponseBody) SetAccessDeniedDetail(v string) *CloudCreateAsrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateAsrResponseBody) SetCode(v string) *CloudCreateAsrResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateAsrResponseBody) SetData(v *CloudCreateAsrResponseBodyData) *CloudCreateAsrResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateAsrResponseBody) SetMessage(v string) *CloudCreateAsrResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateAsrResponseBody) SetRequestId(v string) *CloudCreateAsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateAsrResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateAsrResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudCreateAsrResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAsrResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateAsrResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudCreateAsrResponseBodyData) SetResult(v int64) *CloudCreateAsrResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudCreateAsrResponseBodyData) Validate() error {
	return dara.Validate(s)
}
