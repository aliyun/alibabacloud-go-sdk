// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPauseTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudPauseTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudPauseTaskResponseBody
	GetCode() *string
	SetData(v *CloudPauseTaskResponseBodyData) *CloudPauseTaskResponseBody
	GetData() *CloudPauseTaskResponseBodyData
	SetMessage(v string) *CloudPauseTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudPauseTaskResponseBody
	GetRequestId() *string
}

type CloudPauseTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudPauseTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudPauseTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudPauseTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloudPauseTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudPauseTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudPauseTaskResponseBody) GetData() *CloudPauseTaskResponseBodyData {
	return s.Data
}

func (s *CloudPauseTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudPauseTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudPauseTaskResponseBody) SetAccessDeniedDetail(v string) *CloudPauseTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudPauseTaskResponseBody) SetCode(v string) *CloudPauseTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CloudPauseTaskResponseBody) SetData(v *CloudPauseTaskResponseBodyData) *CloudPauseTaskResponseBody {
	s.Data = v
	return s
}

func (s *CloudPauseTaskResponseBody) SetMessage(v string) *CloudPauseTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CloudPauseTaskResponseBody) SetRequestId(v string) *CloudPauseTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudPauseTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudPauseTaskResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudPauseTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudPauseTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudPauseTaskResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudPauseTaskResponseBodyData) SetResult(v int64) *CloudPauseTaskResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudPauseTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
