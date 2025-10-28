// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ScaleK8sApplicationRequest
	GetAppId() *string
	SetReplicas(v int32) *ScaleK8sApplicationRequest
	GetReplicas() *int32
	SetTimeout(v int32) *ScaleK8sApplicationRequest
	GetTimeout() *int32
}

type ScaleK8sApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 23bf94d9-****-4994-****-616a827aa777
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The expected number of application instances after the scale-out or scale-in. The minimum number is 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The timeout period of the change process. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ScaleK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *ScaleK8sApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *ScaleK8sApplicationRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *ScaleK8sApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ScaleK8sApplicationRequest) SetAppId(v string) *ScaleK8sApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ScaleK8sApplicationRequest) SetReplicas(v int32) *ScaleK8sApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *ScaleK8sApplicationRequest) SetTimeout(v int32) *ScaleK8sApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *ScaleK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
