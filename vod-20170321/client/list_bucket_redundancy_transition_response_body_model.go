// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketRedundancyTransitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRedundancyTransitionInfo(v *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) *ListBucketRedundancyTransitionResponseBody
	GetRedundancyTransitionInfo() *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo
	SetRequestId(v string) *ListBucketRedundancyTransitionResponseBody
	GetRequestId() *string
	SetStorageRedundancyType(v string) *ListBucketRedundancyTransitionResponseBody
	GetStorageRedundancyType() *string
}

type ListBucketRedundancyTransitionResponseBody struct {
	RedundancyTransitionInfo *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo `json:"RedundancyTransitionInfo,omitempty" xml:"RedundancyTransitionInfo,omitempty" type:"Struct"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageRedundancyType    *string                                                             `json:"StorageRedundancyType,omitempty" xml:"StorageRedundancyType,omitempty"`
}

func (s ListBucketRedundancyTransitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBucketRedundancyTransitionResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketRedundancyTransitionResponseBody) GetRedundancyTransitionInfo() *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo {
	return s.RedundancyTransitionInfo
}

func (s *ListBucketRedundancyTransitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBucketRedundancyTransitionResponseBody) GetStorageRedundancyType() *string {
	return s.StorageRedundancyType
}

func (s *ListBucketRedundancyTransitionResponseBody) SetRedundancyTransitionInfo(v *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) *ListBucketRedundancyTransitionResponseBody {
	s.RedundancyTransitionInfo = v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBody) SetRequestId(v string) *ListBucketRedundancyTransitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBody) SetStorageRedundancyType(v string) *ListBucketRedundancyTransitionResponseBody {
	s.StorageRedundancyType = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime                *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EstimatedRemainingTime *string `json:"EstimatedRemainingTime,omitempty" xml:"EstimatedRemainingTime,omitempty"`
	ProcessPercentage      *string `json:"ProcessPercentage,omitempty" xml:"ProcessPercentage,omitempty"`
	StartTime              *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId                 *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) String() string {
	return dara.Prettify(s)
}

func (s ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) GoString() string {
	return s.String()
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) GetEstimatedRemainingTime() *string {
	return s.EstimatedRemainingTime
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) GetProcessPercentage() *string {
	return s.ProcessPercentage
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) GetStatus() *string {
	return s.Status
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) SetCreateTime(v string) *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo {
	s.CreateTime = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) SetEndTime(v string) *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo {
	s.EndTime = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) SetEstimatedRemainingTime(v string) *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo {
	s.EstimatedRemainingTime = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) SetProcessPercentage(v string) *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo {
	s.ProcessPercentage = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) SetStartTime(v string) *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo {
	s.StartTime = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) SetStatus(v string) *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo {
	s.Status = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) SetTaskId(v string) *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo {
	s.TaskId = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponseBodyRedundancyTransitionInfo) Validate() error {
	return dara.Validate(s)
}
