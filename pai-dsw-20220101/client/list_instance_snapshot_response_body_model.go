// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceSnapshotResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListInstanceSnapshotResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstanceSnapshotResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceSnapshotResponseBody
	GetRequestId() *string
	SetSnapshots(v []*ListInstanceSnapshotResponseBodySnapshots) *ListInstanceSnapshotResponseBody
	GetSnapshots() []*ListInstanceSnapshotResponseBodySnapshots
	SetSuccess(v bool) *ListInstanceSnapshotResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListInstanceSnapshotResponseBody
	GetTotalCount() *int64
}

type ListInstanceSnapshotResponseBody struct {
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots []*ListInstanceSnapshotResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 35
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceSnapshotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstanceSnapshotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceSnapshotResponseBody) GetSnapshots() []*ListInstanceSnapshotResponseBodySnapshots {
	return s.Snapshots
}

func (s *ListInstanceSnapshotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceSnapshotResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceSnapshotResponseBody) SetCode(v string) *ListInstanceSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetHttpStatusCode(v int32) *ListInstanceSnapshotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetMessage(v string) *ListInstanceSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetRequestId(v string) *ListInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetSnapshots(v []*ListInstanceSnapshotResponseBodySnapshots) *ListInstanceSnapshotResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetSuccess(v bool) *ListInstanceSnapshotResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetTotalCount(v int64) *ListInstanceSnapshotResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) Validate() error {
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceSnapshotResponseBodySnapshots struct {
	// example:
	//
	// ["/path1","/path2"]
	ExcludePaths []*string `json:"ExcludePaths,omitempty" xml:"ExcludePaths,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// image-05cefd0be2exxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/pai_product/tensorflow:py36_cpu_tf1.12_ubuntu
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// {\\"foo\\": \\"bar\\"}
	Labels []*ListInstanceSnapshotResponseBodySnapshotsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// Internal Error
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// ImagePullBackOff
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// snp-05cexxxxxxxxx
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// example:
	//
	// training_data_env
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// example:
	//
	// Pushing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceSnapshotResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSnapshotResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetExcludePaths() []*string {
	return s.ExcludePaths
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetImageId() *string {
	return s.ImageId
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetLabels() []*ListInstanceSnapshotResponseBodySnapshotsLabels {
	return s.Labels
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *ListInstanceSnapshotResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetExcludePaths(v []*string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ExcludePaths = v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetGmtCreateTime(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetGmtModifiedTime(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetImageId(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ImageId = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetImageUrl(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ImageUrl = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetInstanceId(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetLabels(v []*ListInstanceSnapshotResponseBodySnapshotsLabels) *ListInstanceSnapshotResponseBodySnapshots {
	s.Labels = v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetReasonCode(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ReasonCode = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetReasonMessage(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetSnapshotId(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetSnapshotName(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetStatus(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceSnapshotResponseBodySnapshotsLabels struct {
	// example:
	//
	// stsTokenOwner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstanceSnapshotResponseBodySnapshotsLabels) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSnapshotResponseBodySnapshotsLabels) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotResponseBodySnapshotsLabels) GetKey() *string {
	return s.Key
}

func (s *ListInstanceSnapshotResponseBodySnapshotsLabels) GetValue() *string {
	return s.Value
}

func (s *ListInstanceSnapshotResponseBodySnapshotsLabels) SetKey(v string) *ListInstanceSnapshotResponseBodySnapshotsLabels {
	s.Key = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshotsLabels) SetValue(v string) *ListInstanceSnapshotResponseBodySnapshotsLabels {
	s.Value = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshotsLabels) Validate() error {
	return dara.Validate(s)
}
