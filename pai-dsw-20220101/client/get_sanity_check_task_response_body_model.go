// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSanityCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckDetails(v []*GetSanityCheckTaskResponseBodyCheckDetails) *GetSanityCheckTaskResponseBody
	GetCheckDetails() []*GetSanityCheckTaskResponseBodyCheckDetails
	SetCheckType(v string) *GetSanityCheckTaskResponseBody
	GetCheckType() *string
	SetEndTime(v string) *GetSanityCheckTaskResponseBody
	GetEndTime() *string
	SetInstanceId(v string) *GetSanityCheckTaskResponseBody
	GetInstanceId() *string
	SetIssues(v []*string) *GetSanityCheckTaskResponseBody
	GetIssues() []*string
	SetStartTime(v string) *GetSanityCheckTaskResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetSanityCheckTaskResponseBody
	GetStatus() *string
	SetTaskId(v string) *GetSanityCheckTaskResponseBody
	GetTaskId() *string
	SetRequestId(v string) *GetSanityCheckTaskResponseBody
	GetRequestId() *string
}

type GetSanityCheckTaskResponseBody struct {
	CheckDetails []*GetSanityCheckTaskResponseBodyCheckDetails `json:"CheckDetails,omitempty" xml:"CheckDetails,omitempty" type:"Repeated"`
	// example:
	//
	// SSH
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// 2020-11-08T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Issues     []*string `json:"Issues,omitempty" xml:"Issues,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 917479ff-c869-49ea-908e-ae85bd987bc0
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 44FB0869-AA85-599D-A09D-C42F7467618A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSanityCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSanityCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSanityCheckTaskResponseBody) GetCheckDetails() []*GetSanityCheckTaskResponseBodyCheckDetails {
	return s.CheckDetails
}

func (s *GetSanityCheckTaskResponseBody) GetCheckType() *string {
	return s.CheckType
}

func (s *GetSanityCheckTaskResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSanityCheckTaskResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSanityCheckTaskResponseBody) GetIssues() []*string {
	return s.Issues
}

func (s *GetSanityCheckTaskResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSanityCheckTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSanityCheckTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSanityCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSanityCheckTaskResponseBody) SetCheckDetails(v []*GetSanityCheckTaskResponseBodyCheckDetails) *GetSanityCheckTaskResponseBody {
	s.CheckDetails = v
	return s
}

func (s *GetSanityCheckTaskResponseBody) SetCheckType(v string) *GetSanityCheckTaskResponseBody {
	s.CheckType = &v
	return s
}

func (s *GetSanityCheckTaskResponseBody) SetEndTime(v string) *GetSanityCheckTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetSanityCheckTaskResponseBody) SetInstanceId(v string) *GetSanityCheckTaskResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetSanityCheckTaskResponseBody) SetIssues(v []*string) *GetSanityCheckTaskResponseBody {
	s.Issues = v
	return s
}

func (s *GetSanityCheckTaskResponseBody) SetStartTime(v string) *GetSanityCheckTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetSanityCheckTaskResponseBody) SetStatus(v string) *GetSanityCheckTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetSanityCheckTaskResponseBody) SetTaskId(v string) *GetSanityCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetSanityCheckTaskResponseBody) SetRequestId(v string) *GetSanityCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSanityCheckTaskResponseBody) Validate() error {
	if s.CheckDetails != nil {
		for _, item := range s.CheckDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSanityCheckTaskResponseBodyCheckDetails struct {
	// example:
	//
	// Check whether the security group allows traffic on port 22
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// secureGroupCheck
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Port 22 is blocked by the security group
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// passed
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetSanityCheckTaskResponseBodyCheckDetails) String() string {
	return dara.Prettify(s)
}

func (s GetSanityCheckTaskResponseBodyCheckDetails) GoString() string {
	return s.String()
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) GetDescription() *string {
	return s.Description
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) GetName() *string {
	return s.Name
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) GetReason() *string {
	return s.Reason
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) GetResult() *string {
	return s.Result
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) SetDescription(v string) *GetSanityCheckTaskResponseBodyCheckDetails {
	s.Description = &v
	return s
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) SetName(v string) *GetSanityCheckTaskResponseBodyCheckDetails {
	s.Name = &v
	return s
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) SetReason(v string) *GetSanityCheckTaskResponseBodyCheckDetails {
	s.Reason = &v
	return s
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) SetResult(v string) *GetSanityCheckTaskResponseBodyCheckDetails {
	s.Result = &v
	return s
}

func (s *GetSanityCheckTaskResponseBodyCheckDetails) Validate() error {
	return dara.Validate(s)
}
