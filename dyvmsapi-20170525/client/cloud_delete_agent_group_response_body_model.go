// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteAgentGroupResponseBody
	GetCode() *string
	SetData(v *CloudDeleteAgentGroupResponseBodyData) *CloudDeleteAgentGroupResponseBody
	GetData() *CloudDeleteAgentGroupResponseBodyData
	SetMessage(v string) *CloudDeleteAgentGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteAgentGroupResponseBody
	GetRequestId() *string
}

type CloudDeleteAgentGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteAgentGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9927CD06-C33A-50CC-A9BB-A3427967A66F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteAgentGroupResponseBody) GetData() *CloudDeleteAgentGroupResponseBodyData {
	return s.Data
}

func (s *CloudDeleteAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteAgentGroupResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteAgentGroupResponseBody) SetCode(v string) *CloudDeleteAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteAgentGroupResponseBody) SetData(v *CloudDeleteAgentGroupResponseBodyData) *CloudDeleteAgentGroupResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteAgentGroupResponseBody) SetMessage(v string) *CloudDeleteAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteAgentGroupResponseBody) SetRequestId(v string) *CloudDeleteAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteAgentGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteAgentGroupResponseBodyData struct {
	// 0-成功，其它-失败
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteAgentGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentGroupResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteAgentGroupResponseBodyData) SetResult(v int64) *CloudDeleteAgentGroupResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteAgentGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
