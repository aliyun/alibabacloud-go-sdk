// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DescribeSyncStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeSyncStatusResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeSyncStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSyncStatusResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeSyncStatusResponseBody
	GetHttpStatusCode() *int32
	SetPageNumber(v int32) *DescribeSyncStatusResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeSyncStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSyncStatusResponseBody
	GetSuccess() *bool
	SetSyncStatusList(v []*DescribeSyncStatusResponseBodySyncStatusList) *DescribeSyncStatusResponseBody
	GetSyncStatusList() []*DescribeSyncStatusResponseBodySyncStatusList
}

type DescribeSyncStatusResponseBody struct {
	DynamicCode    *string                                         `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrCode        *string                                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                                         `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	PageNumber     *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	SyncStatusList []*DescribeSyncStatusResponseBodySyncStatusList `json:"SyncStatusList,omitempty" xml:"SyncStatusList,omitempty" type:"Repeated"`
}

func (s DescribeSyncStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeSyncStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeSyncStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSyncStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSyncStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeSyncStatusResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSyncStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSyncStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSyncStatusResponseBody) GetSyncStatusList() []*DescribeSyncStatusResponseBodySyncStatusList {
	return s.SyncStatusList
}

func (s *DescribeSyncStatusResponseBody) SetDynamicCode(v string) *DescribeSyncStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeSyncStatusResponseBody) SetDynamicMessage(v string) *DescribeSyncStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeSyncStatusResponseBody) SetErrCode(v string) *DescribeSyncStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSyncStatusResponseBody) SetErrMessage(v string) *DescribeSyncStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSyncStatusResponseBody) SetHttpStatusCode(v int32) *DescribeSyncStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSyncStatusResponseBody) SetPageNumber(v int32) *DescribeSyncStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSyncStatusResponseBody) SetRequestId(v string) *DescribeSyncStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncStatusResponseBody) SetSuccess(v bool) *DescribeSyncStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSyncStatusResponseBody) SetSyncStatusList(v []*DescribeSyncStatusResponseBodySyncStatusList) *DescribeSyncStatusResponseBody {
	s.SyncStatusList = v
	return s
}

func (s *DescribeSyncStatusResponseBody) Validate() error {
	if s.SyncStatusList != nil {
		for _, item := range s.SyncStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSyncStatusResponseBodySyncStatusList struct {
	Checkpoint *int64  `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Delay      *int64  `json:"Delay,omitempty" xml:"Delay,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Rate       *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSyncStatusResponseBodySyncStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncStatusResponseBodySyncStatusList) GoString() string {
	return s.String()
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) GetCheckpoint() *int64 {
	return s.Checkpoint
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) GetCode() *string {
	return s.Code
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) GetJobId() *string {
	return s.JobId
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) GetRate() *string {
	return s.Rate
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) GetStatus() *string {
	return s.Status
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) SetCheckpoint(v int64) *DescribeSyncStatusResponseBodySyncStatusList {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) SetCode(v string) *DescribeSyncStatusResponseBodySyncStatusList {
	s.Code = &v
	return s
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) SetDelay(v int64) *DescribeSyncStatusResponseBodySyncStatusList {
	s.Delay = &v
	return s
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) SetJobId(v string) *DescribeSyncStatusResponseBodySyncStatusList {
	s.JobId = &v
	return s
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) SetRate(v string) *DescribeSyncStatusResponseBodySyncStatusList {
	s.Rate = &v
	return s
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) SetStatus(v string) *DescribeSyncStatusResponseBodySyncStatusList {
	s.Status = &v
	return s
}

func (s *DescribeSyncStatusResponseBodySyncStatusList) Validate() error {
	return dara.Validate(s)
}
