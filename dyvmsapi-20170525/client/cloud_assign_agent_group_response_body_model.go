// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAssignAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudAssignAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudAssignAgentGroupResponseBody
	GetCode() *string
	SetData(v *CloudAssignAgentGroupResponseBodyData) *CloudAssignAgentGroupResponseBody
	GetData() *CloudAssignAgentGroupResponseBodyData
	SetMessage(v string) *CloudAssignAgentGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudAssignAgentGroupResponseBody
	GetRequestId() *string
}

type CloudAssignAgentGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudAssignAgentGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudAssignAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudAssignAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloudAssignAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudAssignAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudAssignAgentGroupResponseBody) GetData() *CloudAssignAgentGroupResponseBodyData {
	return s.Data
}

func (s *CloudAssignAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudAssignAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudAssignAgentGroupResponseBody) SetAccessDeniedDetail(v string) *CloudAssignAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudAssignAgentGroupResponseBody) SetCode(v string) *CloudAssignAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CloudAssignAgentGroupResponseBody) SetData(v *CloudAssignAgentGroupResponseBodyData) *CloudAssignAgentGroupResponseBody {
	s.Data = v
	return s
}

func (s *CloudAssignAgentGroupResponseBody) SetMessage(v string) *CloudAssignAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CloudAssignAgentGroupResponseBody) SetRequestId(v string) *CloudAssignAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudAssignAgentGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudAssignAgentGroupResponseBodyData struct {
	// 0-成功，1-失败
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudAssignAgentGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudAssignAgentGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudAssignAgentGroupResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudAssignAgentGroupResponseBodyData) SetResult(v int64) *CloudAssignAgentGroupResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudAssignAgentGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
