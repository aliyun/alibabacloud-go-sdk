// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudBatchCreateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudBatchCreateAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudBatchCreateAgentResponseBody
	GetCode() *string
	SetData(v *CloudBatchCreateAgentResponseBodyData) *CloudBatchCreateAgentResponseBody
	GetData() *CloudBatchCreateAgentResponseBodyData
	SetMessage(v string) *CloudBatchCreateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudBatchCreateAgentResponseBody
	GetRequestId() *string
}

type CloudBatchCreateAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudBatchCreateAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudBatchCreateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchCreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudBatchCreateAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudBatchCreateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudBatchCreateAgentResponseBody) GetData() *CloudBatchCreateAgentResponseBodyData {
	return s.Data
}

func (s *CloudBatchCreateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudBatchCreateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudBatchCreateAgentResponseBody) SetAccessDeniedDetail(v string) *CloudBatchCreateAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudBatchCreateAgentResponseBody) SetCode(v string) *CloudBatchCreateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudBatchCreateAgentResponseBody) SetData(v *CloudBatchCreateAgentResponseBodyData) *CloudBatchCreateAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudBatchCreateAgentResponseBody) SetMessage(v string) *CloudBatchCreateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudBatchCreateAgentResponseBody) SetRequestId(v string) *CloudBatchCreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudBatchCreateAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudBatchCreateAgentResponseBodyData struct {
	// 成功创建的座席号
	//
	// example:
	//
	// 100,101,102,103
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 创建座席失败数量
	//
	// example:
	//
	// 0
	Fail *string `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// 座席创建成功,绑定技能失败的数量
	//
	// example:
	//
	// 0
	Other *string `json:"Other,omitempty" xml:"Other,omitempty"`
	// 创建座席成功数量
	//
	// example:
	//
	// 4
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloudBatchCreateAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudBatchCreateAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudBatchCreateAgentResponseBodyData) GetCnos() *string {
	return s.Cnos
}

func (s *CloudBatchCreateAgentResponseBodyData) GetFail() *string {
	return s.Fail
}

func (s *CloudBatchCreateAgentResponseBodyData) GetOther() *string {
	return s.Other
}

func (s *CloudBatchCreateAgentResponseBodyData) GetSuccess() *string {
	return s.Success
}

func (s *CloudBatchCreateAgentResponseBodyData) SetCnos(v string) *CloudBatchCreateAgentResponseBodyData {
	s.Cnos = &v
	return s
}

func (s *CloudBatchCreateAgentResponseBodyData) SetFail(v string) *CloudBatchCreateAgentResponseBodyData {
	s.Fail = &v
	return s
}

func (s *CloudBatchCreateAgentResponseBodyData) SetOther(v string) *CloudBatchCreateAgentResponseBodyData {
	s.Other = &v
	return s
}

func (s *CloudBatchCreateAgentResponseBodyData) SetSuccess(v string) *CloudBatchCreateAgentResponseBodyData {
	s.Success = &v
	return s
}

func (s *CloudBatchCreateAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
