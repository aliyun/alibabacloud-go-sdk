// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitCampaignResponseBody
	GetCode() *string
	SetData(v bool) *SubmitCampaignResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *SubmitCampaignResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitCampaignResponseBody
	GetMessage() *string
	SetParams(v []*string) *SubmitCampaignResponseBody
	GetParams() []*string
	SetRequestId(v string) *SubmitCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitCampaignResponseBody
	GetSuccess() *bool
}

type SubmitCampaignResponseBody struct {
	// 结果码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 操作是否成功
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误信息参数列表
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// 请求ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitCampaignResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitCampaignResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitCampaignResponseBody) GetParams() []*string {
	return s.Params
}

func (s *SubmitCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitCampaignResponseBody) SetCode(v string) *SubmitCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetData(v bool) *SubmitCampaignResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetHttpStatusCode(v int32) *SubmitCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetMessage(v string) *SubmitCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetParams(v []*string) *SubmitCampaignResponseBody {
	s.Params = v
	return s
}

func (s *SubmitCampaignResponseBody) SetRequestId(v string) *SubmitCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetSuccess(v bool) *SubmitCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
