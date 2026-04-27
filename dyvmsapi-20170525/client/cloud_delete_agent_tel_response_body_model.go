// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentTelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteAgentTelResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteAgentTelResponseBody
	GetCode() *string
	SetData(v *CloudDeleteAgentTelResponseBodyData) *CloudDeleteAgentTelResponseBody
	GetData() *CloudDeleteAgentTelResponseBodyData
	SetMessage(v string) *CloudDeleteAgentTelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteAgentTelResponseBody
	GetRequestId() *string
}

type CloudDeleteAgentTelResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteAgentTelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteAgentTelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentTelResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentTelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteAgentTelResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteAgentTelResponseBody) GetData() *CloudDeleteAgentTelResponseBodyData {
	return s.Data
}

func (s *CloudDeleteAgentTelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteAgentTelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteAgentTelResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteAgentTelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteAgentTelResponseBody) SetCode(v string) *CloudDeleteAgentTelResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteAgentTelResponseBody) SetData(v *CloudDeleteAgentTelResponseBodyData) *CloudDeleteAgentTelResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteAgentTelResponseBody) SetMessage(v string) *CloudDeleteAgentTelResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteAgentTelResponseBody) SetRequestId(v string) *CloudDeleteAgentTelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteAgentTelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteAgentTelResponseBodyData struct {
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteAgentTelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentTelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentTelResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteAgentTelResponseBodyData) SetResult(v int64) *CloudDeleteAgentTelResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteAgentTelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
