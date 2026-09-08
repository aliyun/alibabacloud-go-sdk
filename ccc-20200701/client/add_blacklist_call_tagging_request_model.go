// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlacklistCallTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddBlacklistCallTaggingRequest
	GetInstanceId() *string
	SetJobId(v string) *AddBlacklistCallTaggingRequest
	GetJobId() *string
	SetNumber(v string) *AddBlacklistCallTaggingRequest
	GetNumber() *string
}

type AddBlacklistCallTaggingRequest struct {
	// The ID of the Cloud Contact Center instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The call ID. If the phone number is encrypted and contains asterisks (\\*), specify the JobId. This parameter is not required for unencrypted numbers.
	//
	// example:
	//
	// job-6582589278232****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1764590****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s AddBlacklistCallTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBlacklistCallTaggingRequest) GoString() string {
	return s.String()
}

func (s *AddBlacklistCallTaggingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddBlacklistCallTaggingRequest) GetJobId() *string {
	return s.JobId
}

func (s *AddBlacklistCallTaggingRequest) GetNumber() *string {
	return s.Number
}

func (s *AddBlacklistCallTaggingRequest) SetInstanceId(v string) *AddBlacklistCallTaggingRequest {
	s.InstanceId = &v
	return s
}

func (s *AddBlacklistCallTaggingRequest) SetJobId(v string) *AddBlacklistCallTaggingRequest {
	s.JobId = &v
	return s
}

func (s *AddBlacklistCallTaggingRequest) SetNumber(v string) *AddBlacklistCallTaggingRequest {
	s.Number = &v
	return s
}

func (s *AddBlacklistCallTaggingRequest) Validate() error {
	return dara.Validate(s)
}
