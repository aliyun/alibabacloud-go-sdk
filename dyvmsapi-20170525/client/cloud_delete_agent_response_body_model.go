// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteAgentResponseBody
	GetCode() *string
	SetData(v *CloudDeleteAgentResponseBodyData) *CloudDeleteAgentResponseBody
	GetData() *CloudDeleteAgentResponseBodyData
	SetMessage(v string) *CloudDeleteAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteAgentResponseBody
	GetRequestId() *string
}

type CloudDeleteAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteAgentResponseBody) GetData() *CloudDeleteAgentResponseBodyData {
	return s.Data
}

func (s *CloudDeleteAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteAgentResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteAgentResponseBody) SetCode(v string) *CloudDeleteAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteAgentResponseBody) SetData(v *CloudDeleteAgentResponseBodyData) *CloudDeleteAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteAgentResponseBody) SetMessage(v string) *CloudDeleteAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteAgentResponseBody) SetRequestId(v string) *CloudDeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteAgentResponseBodyData struct {
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteAgentResponseBodyData) SetResult(v int64) *CloudDeleteAgentResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
