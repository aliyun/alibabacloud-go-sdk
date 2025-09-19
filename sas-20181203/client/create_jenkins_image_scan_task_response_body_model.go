// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJenkinsImageScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateJenkinsImageScanTaskResponseBodyData) *CreateJenkinsImageScanTaskResponseBody
	GetData() *CreateJenkinsImageScanTaskResponseBodyData
	SetRequestId(v string) *CreateJenkinsImageScanTaskResponseBody
	GetRequestId() *string
}

type CreateJenkinsImageScanTaskResponseBody struct {
	// The data returned.
	Data *CreateJenkinsImageScanTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 52870893-48A7-5A9E-9E05-6253E5B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJenkinsImageScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJenkinsImageScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJenkinsImageScanTaskResponseBody) GetData() *CreateJenkinsImageScanTaskResponseBodyData {
	return s.Data
}

func (s *CreateJenkinsImageScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJenkinsImageScanTaskResponseBody) SetData(v *CreateJenkinsImageScanTaskResponseBodyData) *CreateJenkinsImageScanTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateJenkinsImageScanTaskResponseBody) SetRequestId(v string) *CreateJenkinsImageScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJenkinsImageScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateJenkinsImageScanTaskResponseBodyData struct {
	// The quota for image scan.
	//
	// example:
	//
	// 100
	ImageScanCapacity *int64 `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-upze3gcopm9c****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The instance ID of the image repository.
	//
	// example:
	//
	// cri-0gkaatghnmnt****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RepoRegionId *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	// The ID of the scan task.
	//
	// example:
	//
	// fc98d58eb56f699d49bf7ebbd6d7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The UUID of the image asset.
	//
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateJenkinsImageScanTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateJenkinsImageScanTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) GetImageScanCapacity() *int64 {
	return s.ImageScanCapacity
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) SetImageScanCapacity(v int64) *CreateJenkinsImageScanTaskResponseBodyData {
	s.ImageScanCapacity = &v
	return s
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) SetRepoId(v string) *CreateJenkinsImageScanTaskResponseBodyData {
	s.RepoId = &v
	return s
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) SetRepoInstanceId(v string) *CreateJenkinsImageScanTaskResponseBodyData {
	s.RepoInstanceId = &v
	return s
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) SetRepoRegionId(v string) *CreateJenkinsImageScanTaskResponseBodyData {
	s.RepoRegionId = &v
	return s
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) SetTaskId(v string) *CreateJenkinsImageScanTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) SetUuid(v string) *CreateJenkinsImageScanTaskResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *CreateJenkinsImageScanTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
