// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDownloadUrlV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAgentDownloadUrlV2ResponseBody
	GetCode() *int64
	SetData(v *GetAgentDownloadUrlV2ResponseBodyData) *GetAgentDownloadUrlV2ResponseBody
	GetData() *GetAgentDownloadUrlV2ResponseBodyData
	SetMessage(v string) *GetAgentDownloadUrlV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentDownloadUrlV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentDownloadUrlV2ResponseBody
	GetSuccess() *bool
}

type GetAgentDownloadUrlV2ResponseBody struct {
	// The HTTP status code.\\
	//
	// **Valid values:**
	//
	// 	- 2xx: The request was successful.
	//
	// 	- 3xx: The request was redirected.
	//
	// 	- 4xx: The request was invalid.
	//
	// 	- 5xx: A server error occurred.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The version number and download URL of the agent.
	Data *GetAgentDownloadUrlV2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.\\
	//
	// **Valid values:**
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentDownloadUrlV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDownloadUrlV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlV2ResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAgentDownloadUrlV2ResponseBody) GetData() *GetAgentDownloadUrlV2ResponseBodyData {
	return s.Data
}

func (s *GetAgentDownloadUrlV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentDownloadUrlV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentDownloadUrlV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentDownloadUrlV2ResponseBody) SetCode(v int64) *GetAgentDownloadUrlV2ResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentDownloadUrlV2ResponseBody) SetData(v *GetAgentDownloadUrlV2ResponseBodyData) *GetAgentDownloadUrlV2ResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentDownloadUrlV2ResponseBody) SetMessage(v string) *GetAgentDownloadUrlV2ResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentDownloadUrlV2ResponseBody) SetRequestId(v string) *GetAgentDownloadUrlV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentDownloadUrlV2ResponseBody) SetSuccess(v bool) *GetAgentDownloadUrlV2ResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentDownloadUrlV2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentDownloadUrlV2ResponseBodyData struct {
	// The download URL of the agent.
	//
	// example:
	//
	// http://arms-apm-cn-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/3.2.9/ArmsAgent.zip
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The version number of the agent.
	//
	// example:
	//
	// 3.2.9
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAgentDownloadUrlV2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDownloadUrlV2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlV2ResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetAgentDownloadUrlV2ResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetAgentDownloadUrlV2ResponseBodyData) SetUrl(v string) *GetAgentDownloadUrlV2ResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetAgentDownloadUrlV2ResponseBodyData) SetVersion(v string) *GetAgentDownloadUrlV2ResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetAgentDownloadUrlV2ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
