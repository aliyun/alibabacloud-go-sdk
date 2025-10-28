// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RestartK8sApplicationRequest
	GetAppId() *string
	SetTimeout(v int32) *RestartK8sApplicationRequest
	GetTimeout() *int32
}

type RestartK8sApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 93fdd228-********ed2ae98de18d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The timeout period of the change process. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RestartK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *RestartK8sApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RestartK8sApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *RestartK8sApplicationRequest) SetAppId(v string) *RestartK8sApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RestartK8sApplicationRequest) SetTimeout(v int32) *RestartK8sApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *RestartK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
