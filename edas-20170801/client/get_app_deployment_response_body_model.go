// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppDeploymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAppDeploymentResponseBody
	GetCode() *int32
	SetData(v string) *GetAppDeploymentResponseBody
	GetData() *string
	SetMessage(v string) *GetAppDeploymentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppDeploymentResponseBody
	GetRequestId() *string
}

type GetAppDeploymentResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the Deployment of the application. The value is a JSON string.
	//
	// example:
	//
	// {"metadata":{"name":"oambuild-group*","namespace":"default","selfLink":"/apis/apps/v1/namespaces/default/deployments/oambuil*","uid":"*-afc0-436e-9949-fb01994a9b63","resourceVersion":"969614832","generation":2,"creationTimestamp":"2021-04-06T11:38:46Z","labels":{"edas-domain":"edas-admin","edas.aliyun.oam.com/rollout-name":"oambuild-group-1","edas.aliyun.oam.com/rollout-namespace":"default","edas.aliyun.oam.com/rollout-revision":"3","edas.appid":"fc5e0f54-*-4cab-91a0-b7becb1f6174","edas.controlplane":"edas-oam","edas.oam.acname":"oambuild"...
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 03FD1520-0FD6-436A-****-265318D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppDeploymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppDeploymentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAppDeploymentResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAppDeploymentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppDeploymentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppDeploymentResponseBody) SetCode(v int32) *GetAppDeploymentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppDeploymentResponseBody) SetData(v string) *GetAppDeploymentResponseBody {
	s.Data = &v
	return s
}

func (s *GetAppDeploymentResponseBody) SetMessage(v string) *GetAppDeploymentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppDeploymentResponseBody) SetRequestId(v string) *GetAppDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppDeploymentResponseBody) Validate() error {
	return dara.Validate(s)
}
