// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudUnassignAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudUnassignAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudUnassignAgentGroupResponseBody
	GetCode() *string
	SetData(v *CloudUnassignAgentGroupResponseBodyData) *CloudUnassignAgentGroupResponseBody
	GetData() *CloudUnassignAgentGroupResponseBodyData
	SetMessage(v string) *CloudUnassignAgentGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudUnassignAgentGroupResponseBody
	GetRequestId() *string
}

type CloudUnassignAgentGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudUnassignAgentGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudUnassignAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudUnassignAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloudUnassignAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudUnassignAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudUnassignAgentGroupResponseBody) GetData() *CloudUnassignAgentGroupResponseBodyData {
	return s.Data
}

func (s *CloudUnassignAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudUnassignAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudUnassignAgentGroupResponseBody) SetAccessDeniedDetail(v string) *CloudUnassignAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudUnassignAgentGroupResponseBody) SetCode(v string) *CloudUnassignAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CloudUnassignAgentGroupResponseBody) SetData(v *CloudUnassignAgentGroupResponseBodyData) *CloudUnassignAgentGroupResponseBody {
	s.Data = v
	return s
}

func (s *CloudUnassignAgentGroupResponseBody) SetMessage(v string) *CloudUnassignAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CloudUnassignAgentGroupResponseBody) SetRequestId(v string) *CloudUnassignAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudUnassignAgentGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudUnassignAgentGroupResponseBodyData struct {
	// 0-成功，其它-失败
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudUnassignAgentGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudUnassignAgentGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudUnassignAgentGroupResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudUnassignAgentGroupResponseBodyData) SetResult(v int64) *CloudUnassignAgentGroupResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudUnassignAgentGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
