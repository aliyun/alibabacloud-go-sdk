// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPictureSearchJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *GetPictureSearchJobStatusRequest
	GetAppInstanceId() *string
	SetJobId(v string) *GetPictureSearchJobStatusRequest
	GetJobId() *string
}

type GetPictureSearchJobStatusRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](https://help.aliyun.com/document_detail/206608.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 892f97c9518548469d319d955d98****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the image search task. You can call the [QueryPictureSearchJob](https://help.aliyun.com/document_detail/206778.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b3786935a71442efae014c7b2212****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetPictureSearchJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPictureSearchJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPictureSearchJobStatusRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *GetPictureSearchJobStatusRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetPictureSearchJobStatusRequest) SetAppInstanceId(v string) *GetPictureSearchJobStatusRequest {
	s.AppInstanceId = &v
	return s
}

func (s *GetPictureSearchJobStatusRequest) SetJobId(v string) *GetPictureSearchJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetPictureSearchJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
