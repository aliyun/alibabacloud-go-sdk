// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFotaTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFotaTasksResponseBody
	GetCode() *string
	SetFotaTasks(v []*DescribeFotaTasksResponseBodyFotaTasks) *DescribeFotaTasksResponseBody
	GetFotaTasks() []*DescribeFotaTasksResponseBodyFotaTasks
	SetMessage(v string) *DescribeFotaTasksResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeFotaTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeFotaTasksResponseBody
	GetRequestId() *string
}

type DescribeFotaTasksResponseBody struct {
	// The returned message. If the request was successful, a `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details about the image update task.
	FotaTasks []*DescribeFotaTasksResponseBodyFotaTasks `json:"FotaTasks,omitempty" xml:"FotaTasks,omitempty" type:"Repeated"`
	// The returned error message. This parameter is not returned if the Code value is a `success` message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFotaTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFotaTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFotaTasksResponseBody) GetFotaTasks() []*DescribeFotaTasksResponseBodyFotaTasks {
	return s.FotaTasks
}

func (s *DescribeFotaTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFotaTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFotaTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFotaTasksResponseBody) SetCode(v string) *DescribeFotaTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFotaTasksResponseBody) SetFotaTasks(v []*DescribeFotaTasksResponseBodyFotaTasks) *DescribeFotaTasksResponseBody {
	s.FotaTasks = v
	return s
}

func (s *DescribeFotaTasksResponseBody) SetMessage(v string) *DescribeFotaTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFotaTasksResponseBody) SetNextToken(v string) *DescribeFotaTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFotaTasksResponseBody) SetRequestId(v string) *DescribeFotaTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFotaTasksResponseBody) Validate() error {
	if s.FotaTasks != nil {
		for _, item := range s.FotaTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFotaTasksResponseBodyFotaTasks struct {
	// The image version. You can call the [DescribeImages](https://help.aliyun.com/document_detail/188895.html) operation to obtain the value of this parameter.
	//
	// example:
	//
	// 0.0.1-D-20220513.143129
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// To be hidden.
	FotaProject *string `json:"FotaProject,omitempty" xml:"FotaProject,omitempty"`
	// The number of custom images that can be updated to this version.
	//
	// example:
	//
	// 1
	PendingCustomImageCount *int32 `json:"PendingCustomImageCount,omitempty" xml:"PendingCustomImageCount,omitempty"`
	// The number of cloud computers whose images can be updated to this version.
	//
	// example:
	//
	// 1
	PendingDesktopCount *int32 `json:"PendingDesktopCount,omitempty" xml:"PendingDesktopCount,omitempty"`
	// The time when the image version available for update was published.
	//
	// example:
	//
	// 2022-05-31T04:28:48Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The description of the image version available for update.
	//
	// example:
	//
	// test
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The size of the update package. Unit: KB.
	//
	// example:
	//
	// 568533470
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Indicates whether the image update task is automatically pushed.
	//
	// Valid values:
	//
	// 	- Running: automatically pushes the image update task.
	//
	// 	- Pending: does not automatically push the image update task.
	//
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the image upgrade task.
	//
	// example:
	//
	// aot-c4khwrp9ocml4****
	TaskUid *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s DescribeFotaTasksResponseBodyFotaTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaTasksResponseBodyFotaTasks) GoString() string {
	return s.String()
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetAppVersion() *string {
	return s.AppVersion
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetFotaProject() *string {
	return s.FotaProject
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetPendingCustomImageCount() *int32 {
	return s.PendingCustomImageCount
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetPendingDesktopCount() *int32 {
	return s.PendingDesktopCount
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetPublishTime() *string {
	return s.PublishTime
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetSize() *int32 {
	return s.Size
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) GetTaskUid() *string {
	return s.TaskUid
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetAppVersion(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.AppVersion = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetFotaProject(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.FotaProject = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetPendingCustomImageCount(v int32) *DescribeFotaTasksResponseBodyFotaTasks {
	s.PendingCustomImageCount = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetPendingDesktopCount(v int32) *DescribeFotaTasksResponseBodyFotaTasks {
	s.PendingDesktopCount = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetPublishTime(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.PublishTime = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetReleaseNote(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetSize(v int32) *DescribeFotaTasksResponseBodyFotaTasks {
	s.Size = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetStatus(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.Status = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) SetTaskUid(v string) *DescribeFotaTasksResponseBodyFotaTasks {
	s.TaskUid = &v
	return s
}

func (s *DescribeFotaTasksResponseBodyFotaTasks) Validate() error {
	return dara.Validate(s)
}
