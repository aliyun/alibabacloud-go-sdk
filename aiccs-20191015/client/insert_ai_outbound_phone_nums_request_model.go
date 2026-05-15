// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertAiOutboundPhoneNumsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchVersion(v int32) *InsertAiOutboundPhoneNumsRequest
	GetBatchVersion() *int32
	SetDetails(v []*InsertAiOutboundPhoneNumsRequestDetails) *InsertAiOutboundPhoneNumsRequest
	GetDetails() []*InsertAiOutboundPhoneNumsRequestDetails
	SetInstanceId(v string) *InsertAiOutboundPhoneNumsRequest
	GetInstanceId() *string
	SetTaskId(v int64) *InsertAiOutboundPhoneNumsRequest
	GetTaskId() *int64
}

type InsertAiOutboundPhoneNumsRequest struct {
	// example:
	//
	// 2
	BatchVersion *int32 `json:"BatchVersion,omitempty" xml:"BatchVersion,omitempty"`
	// This parameter is required.
	Details []*InsertAiOutboundPhoneNumsRequestDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InsertAiOutboundPhoneNumsRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertAiOutboundPhoneNumsRequest) GoString() string {
	return s.String()
}

func (s *InsertAiOutboundPhoneNumsRequest) GetBatchVersion() *int32 {
	return s.BatchVersion
}

func (s *InsertAiOutboundPhoneNumsRequest) GetDetails() []*InsertAiOutboundPhoneNumsRequestDetails {
	return s.Details
}

func (s *InsertAiOutboundPhoneNumsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InsertAiOutboundPhoneNumsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *InsertAiOutboundPhoneNumsRequest) SetBatchVersion(v int32) *InsertAiOutboundPhoneNumsRequest {
	s.BatchVersion = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsRequest) SetDetails(v []*InsertAiOutboundPhoneNumsRequestDetails) *InsertAiOutboundPhoneNumsRequest {
	s.Details = v
	return s
}

func (s *InsertAiOutboundPhoneNumsRequest) SetInstanceId(v string) *InsertAiOutboundPhoneNumsRequest {
	s.InstanceId = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsRequest) SetTaskId(v int64) *InsertAiOutboundPhoneNumsRequest {
	s.TaskId = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsRequest) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InsertAiOutboundPhoneNumsRequestDetails struct {
	// example:
	//
	// xxxx
	BizData *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	// example:
	//
	// 150****0000
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
}

func (s InsertAiOutboundPhoneNumsRequestDetails) String() string {
	return dara.Prettify(s)
}

func (s InsertAiOutboundPhoneNumsRequestDetails) GoString() string {
	return s.String()
}

func (s *InsertAiOutboundPhoneNumsRequestDetails) GetBizData() *string {
	return s.BizData
}

func (s *InsertAiOutboundPhoneNumsRequestDetails) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *InsertAiOutboundPhoneNumsRequestDetails) SetBizData(v string) *InsertAiOutboundPhoneNumsRequestDetails {
	s.BizData = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsRequestDetails) SetPhoneNum(v string) *InsertAiOutboundPhoneNumsRequestDetails {
	s.PhoneNum = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsRequestDetails) Validate() error {
	return dara.Validate(s)
}
