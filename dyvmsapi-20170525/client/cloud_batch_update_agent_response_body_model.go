// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchUpdateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudBatchUpdateAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudBatchUpdateAgentResponseBody
	GetCode() *string
	SetData(v *CloudBatchUpdateAgentResponseBodyData) *CloudBatchUpdateAgentResponseBody
	GetData() *CloudBatchUpdateAgentResponseBodyData
	SetMessage(v string) *CloudBatchUpdateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudBatchUpdateAgentResponseBody
	GetRequestId() *string
}

type CloudBatchUpdateAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudBatchUpdateAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudBatchUpdateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchUpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudBatchUpdateAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudBatchUpdateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudBatchUpdateAgentResponseBody) GetData() *CloudBatchUpdateAgentResponseBodyData {
	return s.Data
}

func (s *CloudBatchUpdateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudBatchUpdateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudBatchUpdateAgentResponseBody) SetAccessDeniedDetail(v string) *CloudBatchUpdateAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudBatchUpdateAgentResponseBody) SetCode(v string) *CloudBatchUpdateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudBatchUpdateAgentResponseBody) SetData(v *CloudBatchUpdateAgentResponseBodyData) *CloudBatchUpdateAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudBatchUpdateAgentResponseBody) SetMessage(v string) *CloudBatchUpdateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudBatchUpdateAgentResponseBody) SetRequestId(v string) *CloudBatchUpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudBatchUpdateAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudBatchUpdateAgentResponseBodyData struct {
	// 座席配置信息修改失败的座席工号列表
	//
	// example:
	//
	// null
	Fail *string `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// 座席配置信息修改成功数量
	//
	// example:
	//
	// 3
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloudBatchUpdateAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchUpdateAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudBatchUpdateAgentResponseBodyData) GetFail() *string {
	return s.Fail
}

func (s *CloudBatchUpdateAgentResponseBodyData) GetSuccess() *string {
	return s.Success
}

func (s *CloudBatchUpdateAgentResponseBodyData) SetFail(v string) *CloudBatchUpdateAgentResponseBodyData {
	s.Fail = &v
	return s
}

func (s *CloudBatchUpdateAgentResponseBodyData) SetSuccess(v string) *CloudBatchUpdateAgentResponseBodyData {
	s.Success = &v
	return s
}

func (s *CloudBatchUpdateAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
