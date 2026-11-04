// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocalFileUploadJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIotInstanceId(v string) *QueryLocalFileUploadJobRequest
	GetIotInstanceId() *string
	SetJobId(v string) *QueryLocalFileUploadJobRequest
	GetJobId() *string
}

type QueryLocalFileUploadJobRequest struct {
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify this ID. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID, see [Obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ID of the local video recording upload task. The task is created by calling [CreateLocalFileUploadJob](https://help.aliyun.com/document_detail/363002.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ec1a****5d0435fbc0****34f6f139d
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryLocalFileUploadJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalFileUploadJobRequest) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryLocalFileUploadJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryLocalFileUploadJobRequest) SetIotInstanceId(v string) *QueryLocalFileUploadJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryLocalFileUploadJobRequest) SetJobId(v string) *QueryLocalFileUploadJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryLocalFileUploadJobRequest) Validate() error {
	return dara.Validate(s)
}
