// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceSnapshotResponseBody
	GetCode() *string
	SetExcludePaths(v []*string) *GetInstanceSnapshotResponseBody
	GetExcludePaths() []*string
	SetGmtCreateTime(v string) *GetInstanceSnapshotResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetInstanceSnapshotResponseBody
	GetGmtModifiedTime() *string
	SetHttpStatusCode(v int32) *GetInstanceSnapshotResponseBody
	GetHttpStatusCode() *int32
	SetImageId(v string) *GetInstanceSnapshotResponseBody
	GetImageId() *string
	SetImageUrl(v string) *GetInstanceSnapshotResponseBody
	GetImageUrl() *string
	SetInstanceId(v string) *GetInstanceSnapshotResponseBody
	GetInstanceId() *string
	SetLabels(v []*GetInstanceSnapshotResponseBodyLabels) *GetInstanceSnapshotResponseBody
	GetLabels() []*GetInstanceSnapshotResponseBodyLabels
	SetMessage(v string) *GetInstanceSnapshotResponseBody
	GetMessage() *string
	SetReasonCode(v string) *GetInstanceSnapshotResponseBody
	GetReasonCode() *string
	SetReasonMessage(v string) *GetInstanceSnapshotResponseBody
	GetReasonMessage() *string
	SetRequestId(v string) *GetInstanceSnapshotResponseBody
	GetRequestId() *string
	SetSnapshotId(v string) *GetInstanceSnapshotResponseBody
	GetSnapshotId() *string
	SetSnapshotName(v string) *GetInstanceSnapshotResponseBody
	GetSnapshotName() *string
	SetStatus(v string) *GetInstanceSnapshotResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetInstanceSnapshotResponseBody
	GetSuccess() *bool
}

type GetInstanceSnapshotResponseBody struct {
	// example:
	//
	// null
	Code         *string   `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// {\"foo\": \"bar\"}
	Labels []*GetInstanceSnapshotResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceSnapshotResponseBody) GetExcludePaths() []*string {
	return s.ExcludePaths
}

func (s *GetInstanceSnapshotResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceSnapshotResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceSnapshotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceSnapshotResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *GetInstanceSnapshotResponseBody) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetInstanceSnapshotResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceSnapshotResponseBody) GetLabels() []*GetInstanceSnapshotResponseBodyLabels {
	return s.Labels
}

func (s *GetInstanceSnapshotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceSnapshotResponseBody) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetInstanceSnapshotResponseBody) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetInstanceSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceSnapshotResponseBody) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *GetInstanceSnapshotResponseBody) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *GetInstanceSnapshotResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceSnapshotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceSnapshotResponseBody) SetCode(v string) *GetInstanceSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetExcludePaths(v []*string) *GetInstanceSnapshotResponseBody {
	s.ExcludePaths = v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetGmtCreateTime(v string) *GetInstanceSnapshotResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetGmtModifiedTime(v string) *GetInstanceSnapshotResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetHttpStatusCode(v int32) *GetInstanceSnapshotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetImageId(v string) *GetInstanceSnapshotResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetImageUrl(v string) *GetInstanceSnapshotResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceId(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetLabels(v []*GetInstanceSnapshotResponseBodyLabels) *GetInstanceSnapshotResponseBody {
	s.Labels = v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetMessage(v string) *GetInstanceSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetReasonCode(v string) *GetInstanceSnapshotResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetReasonMessage(v string) *GetInstanceSnapshotResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetRequestId(v string) *GetInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetSnapshotId(v string) *GetInstanceSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetSnapshotName(v string) *GetInstanceSnapshotResponseBody {
	s.SnapshotName = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetStatus(v string) *GetInstanceSnapshotResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetSuccess(v bool) *GetInstanceSnapshotResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) Validate() error {
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

type GetInstanceSnapshotResponseBodyLabels struct {
	// example:
	//
	// stsTokenOwner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceSnapshotResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSnapshotResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotResponseBodyLabels) GetKey() *string {
	return s.Key
}

func (s *GetInstanceSnapshotResponseBodyLabels) GetValue() *string {
	return s.Value
}

func (s *GetInstanceSnapshotResponseBodyLabels) SetKey(v string) *GetInstanceSnapshotResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetInstanceSnapshotResponseBodyLabels) SetValue(v string) *GetInstanceSnapshotResponseBodyLabels {
	s.Value = &v
	return s
}

func (s *GetInstanceSnapshotResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}
