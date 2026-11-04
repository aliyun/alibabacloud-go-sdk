// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicePictureLifeCycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIotId(v string) *QueryDevicePictureLifeCycleRequest
	GetIotId() *string
}

type QueryDevicePictureLifeCycleRequest struct {
	// The ID of the IPC device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the IoT Platform [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWCg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryDevicePictureLifeCycleRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicePictureLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureLifeCycleRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryDevicePictureLifeCycleRequest) SetIotId(v string) *QueryDevicePictureLifeCycleRequest {
	s.IotId = &v
	return s
}

func (s *QueryDevicePictureLifeCycleRequest) Validate() error {
	return dara.Validate(s)
}
