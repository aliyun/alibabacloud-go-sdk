// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopK8sApplicationRequest
	GetAppId() *string
	SetTimeout(v int32) *StopK8sApplicationRequest
	GetTimeout() *int32
}

type StopK8sApplicationRequest struct {
	// The ID of the application. You can query the application ID by calling the ListApplication operation. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 93fdd228-d237-*******8de18d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The timeout period of the change process. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s StopK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *StopK8sApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopK8sApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *StopK8sApplicationRequest) SetAppId(v string) *StopK8sApplicationRequest {
	s.AppId = &v
	return s
}

func (s *StopK8sApplicationRequest) SetTimeout(v int32) *StopK8sApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *StopK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
