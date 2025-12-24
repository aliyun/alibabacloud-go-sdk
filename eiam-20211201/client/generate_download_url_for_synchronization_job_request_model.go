// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDownloadUrlForSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GenerateDownloadUrlForSynchronizationJobRequest
	GetInstanceId() *string
	SetSynchronizationJobId(v string) *GenerateDownloadUrlForSynchronizationJobRequest
	GetSynchronizationJobId() *string
}

type GenerateDownloadUrlForSynchronizationJobRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 同步任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// sync_000036q9p3jd5s18boeo2dvmmanu2086vxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s GenerateDownloadUrlForSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDownloadUrlForSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *GenerateDownloadUrlForSynchronizationJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateDownloadUrlForSynchronizationJobRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *GenerateDownloadUrlForSynchronizationJobRequest) SetInstanceId(v string) *GenerateDownloadUrlForSynchronizationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateDownloadUrlForSynchronizationJobRequest) SetSynchronizationJobId(v string) *GenerateDownloadUrlForSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *GenerateDownloadUrlForSynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}
