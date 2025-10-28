// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartK8sApplicationRequest
	GetAppId() *string
	SetReplicas(v int32) *StartK8sApplicationRequest
	GetReplicas() *int32
	SetTimeout(v int32) *StartK8sApplicationRequest
	GetTimeout() *int32
}

type StartK8sApplicationRequest struct {
	// The ID of the application. You can query the application ID by calling the ListApplication operation. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 93fdd228-*******-ed2ae98de18d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of instances on which you want to start the application.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The timeout period of the change process. Valid values: 1 to 1800. Default value: 600. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s StartK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *StartK8sApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartK8sApplicationRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *StartK8sApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *StartK8sApplicationRequest) SetAppId(v string) *StartK8sApplicationRequest {
	s.AppId = &v
	return s
}

func (s *StartK8sApplicationRequest) SetReplicas(v int32) *StartK8sApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *StartK8sApplicationRequest) SetTimeout(v int32) *StartK8sApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *StartK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
