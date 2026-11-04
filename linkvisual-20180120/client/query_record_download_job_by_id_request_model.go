// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordDownloadJobByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIotInstanceId(v string) *QueryRecordDownloadJobByIdRequest
	GetIotInstanceId() *string
	SetJobId(v string) *QueryRecordDownloadJobByIdRequest
	GetJobId() *string
}

type QueryRecordDownloadJobByIdRequest struct {
	// The instance ID. You can view the ID of the current instance on the **Instance Overview*	- tab in the IoT Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// Iot-r******
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ID of the cloud record download task.
	//
	// This parameter is required.
	//
	// example:
	//
	// w2s******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryRecordDownloadJobByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordDownloadJobByIdRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordDownloadJobByIdRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRecordDownloadJobByIdRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryRecordDownloadJobByIdRequest) SetIotInstanceId(v string) *QueryRecordDownloadJobByIdRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordDownloadJobByIdRequest) SetJobId(v string) *QueryRecordDownloadJobByIdRequest {
	s.JobId = &v
	return s
}

func (s *QueryRecordDownloadJobByIdRequest) Validate() error {
	return dara.Validate(s)
}
