// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeCampaignResponseBody
	GetCode() *string
	SetData(v bool) *ResumeCampaignResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *ResumeCampaignResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResumeCampaignResponseBody
	GetMessage() *string
	SetParams(v []*string) *ResumeCampaignResponseBody
	GetParams() []*string
	SetRequestId(v string) *ResumeCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResumeCampaignResponseBody
	GetSuccess() *bool
}

type ResumeCampaignResponseBody struct {
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
	// 无
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

func (s ResumeCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeCampaignResponseBody) GetData() *bool {
	return s.Data
}

func (s *ResumeCampaignResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResumeCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeCampaignResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ResumeCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeCampaignResponseBody) SetCode(v string) *ResumeCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetData(v bool) *ResumeCampaignResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetHttpStatusCode(v int32) *ResumeCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetMessage(v string) *ResumeCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetParams(v []*string) *ResumeCampaignResponseBody {
	s.Params = v
	return s
}

func (s *ResumeCampaignResponseBody) SetRequestId(v string) *ResumeCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetSuccess(v bool) *ResumeCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
