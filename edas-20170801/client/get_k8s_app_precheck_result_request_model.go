// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sAppPrecheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetK8sAppPrecheckResultRequest
	GetAppName() *string
	SetClusterId(v string) *GetK8sAppPrecheckResultRequest
	GetClusterId() *string
	SetNamespace(v string) *GetK8sAppPrecheckResultRequest
	GetNamespace() *string
}

type GetK8sAppPrecheckResultRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// testapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the cluster in Enterprise Distributed Application Service (EDAS).
	//
	// This parameter is required.
	//
	// example:
	//
	// c37aec2a-bcca-4ec1-****-************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetK8sAppPrecheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetK8sAppPrecheckResultRequest) GoString() string {
	return s.String()
}

func (s *GetK8sAppPrecheckResultRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetK8sAppPrecheckResultRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetK8sAppPrecheckResultRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetK8sAppPrecheckResultRequest) SetAppName(v string) *GetK8sAppPrecheckResultRequest {
	s.AppName = &v
	return s
}

func (s *GetK8sAppPrecheckResultRequest) SetClusterId(v string) *GetK8sAppPrecheckResultRequest {
	s.ClusterId = &v
	return s
}

func (s *GetK8sAppPrecheckResultRequest) SetNamespace(v string) *GetK8sAppPrecheckResultRequest {
	s.Namespace = &v
	return s
}

func (s *GetK8sAppPrecheckResultRequest) Validate() error {
	return dara.Validate(s)
}
