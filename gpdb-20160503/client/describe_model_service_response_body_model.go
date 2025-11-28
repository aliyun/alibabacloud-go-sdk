// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiNodes(v []*string) *DescribeModelServiceResponseBody
	GetAiNodes() []*string
	SetApiKey(v string) *DescribeModelServiceResponseBody
	GetApiKey() *string
	SetCreateTime(v string) *DescribeModelServiceResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeModelServiceResponseBody
	GetDescription() *string
	SetModelName(v string) *DescribeModelServiceResponseBody
	GetModelName() *string
	SetModelParams(v map[string]interface{}) *DescribeModelServiceResponseBody
	GetModelParams() map[string]interface{}
	SetModelServiceId(v string) *DescribeModelServiceResponseBody
	GetModelServiceId() *string
	SetPrivateConnectUrl(v string) *DescribeModelServiceResponseBody
	GetPrivateConnectUrl() *string
	SetPublicConnectUrl(v string) *DescribeModelServiceResponseBody
	GetPublicConnectUrl() *string
	SetRequestId(v string) *DescribeModelServiceResponseBody
	GetRequestId() *string
	SetSecurityIpList(v string) *DescribeModelServiceResponseBody
	GetSecurityIpList() *string
	SetStatus(v string) *DescribeModelServiceResponseBody
	GetStatus() *string
}

type DescribeModelServiceResponseBody struct {
	// The list of AI nodes.
	AiNodes []*string `json:"AiNodes,omitempty" xml:"AiNodes,omitempty" type:"Repeated"`
	// The API key.
	//
	// example:
	//
	// cx/Y57lTNf*********
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The model name.
	//
	// example:
	//
	// Qwen3-Embedding-8B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Model parameters.
	ModelParams map[string]interface{} `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	// The model service ID.
	//
	// example:
	//
	// ms-xxxxxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// http://ms-xxxxxxx.xxxx.rds.aliyuncs.com
	PrivateConnectUrl *string `json:"PrivateConnectUrl,omitempty" xml:"PrivateConnectUrl,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// http://ms-xxxxxxx-o.xxxx.rds.aliyuncs.com
	PublicConnectUrl *string `json:"PublicConnectUrl,omitempty" xml:"PublicConnectUrl,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A comma-separated list of IP addresses and CIDR blocks allowed to connect.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	// The status of the model service.
	//
	// example:
	//
	// - deployming
	//
	// - active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeModelServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelServiceResponseBody) GetAiNodes() []*string {
	return s.AiNodes
}

func (s *DescribeModelServiceResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeModelServiceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeModelServiceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeModelServiceResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeModelServiceResponseBody) GetModelParams() map[string]interface{} {
	return s.ModelParams
}

func (s *DescribeModelServiceResponseBody) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *DescribeModelServiceResponseBody) GetPrivateConnectUrl() *string {
	return s.PrivateConnectUrl
}

func (s *DescribeModelServiceResponseBody) GetPublicConnectUrl() *string {
	return s.PublicConnectUrl
}

func (s *DescribeModelServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelServiceResponseBody) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *DescribeModelServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeModelServiceResponseBody) SetAiNodes(v []*string) *DescribeModelServiceResponseBody {
	s.AiNodes = v
	return s
}

func (s *DescribeModelServiceResponseBody) SetApiKey(v string) *DescribeModelServiceResponseBody {
	s.ApiKey = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetCreateTime(v string) *DescribeModelServiceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetDescription(v string) *DescribeModelServiceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetModelName(v string) *DescribeModelServiceResponseBody {
	s.ModelName = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetModelParams(v map[string]interface{}) *DescribeModelServiceResponseBody {
	s.ModelParams = v
	return s
}

func (s *DescribeModelServiceResponseBody) SetModelServiceId(v string) *DescribeModelServiceResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetPrivateConnectUrl(v string) *DescribeModelServiceResponseBody {
	s.PrivateConnectUrl = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetPublicConnectUrl(v string) *DescribeModelServiceResponseBody {
	s.PublicConnectUrl = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetRequestId(v string) *DescribeModelServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetSecurityIpList(v string) *DescribeModelServiceResponseBody {
	s.SecurityIpList = &v
	return s
}

func (s *DescribeModelServiceResponseBody) SetStatus(v string) *DescribeModelServiceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeModelServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
