// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBlacklistCallTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveBlacklistCallTaggingRequest
	GetInstanceId() *string
	SetJobId(v string) *RemoveBlacklistCallTaggingRequest
	GetJobId() *string
	SetNumber(v string) *RemoveBlacklistCallTaggingRequest
	GetNumber() *string
}

type RemoveBlacklistCallTaggingRequest struct {
	// ID of the Cloud Contact Center instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Call ID. Provide this field only for masked numbers (containing \\*). Do not provide it for unmasked numbers.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1312121****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s RemoveBlacklistCallTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveBlacklistCallTaggingRequest) GoString() string {
	return s.String()
}

func (s *RemoveBlacklistCallTaggingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveBlacklistCallTaggingRequest) GetJobId() *string {
	return s.JobId
}

func (s *RemoveBlacklistCallTaggingRequest) GetNumber() *string {
	return s.Number
}

func (s *RemoveBlacklistCallTaggingRequest) SetInstanceId(v string) *RemoveBlacklistCallTaggingRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveBlacklistCallTaggingRequest) SetJobId(v string) *RemoveBlacklistCallTaggingRequest {
	s.JobId = &v
	return s
}

func (s *RemoveBlacklistCallTaggingRequest) SetNumber(v string) *RemoveBlacklistCallTaggingRequest {
	s.Number = &v
	return s
}

func (s *RemoveBlacklistCallTaggingRequest) Validate() error {
	return dara.Validate(s)
}
