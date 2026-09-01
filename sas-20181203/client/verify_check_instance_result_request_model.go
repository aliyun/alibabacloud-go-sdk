// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckInstanceResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *VerifyCheckInstanceResultRequest
	GetCheckId() *int64
	SetCheckIds(v []*int64) *VerifyCheckInstanceResultRequest
	GetCheckIds() []*int64
	SetClientToken(v string) *VerifyCheckInstanceResultRequest
	GetClientToken() *string
	SetInstanceIds(v []*string) *VerifyCheckInstanceResultRequest
	GetInstanceIds() []*string
	SetTaskSource(v string) *VerifyCheckInstanceResultRequest
	GetTaskSource() *string
}

type VerifyCheckInstanceResultRequest struct {
	// The ID of the check item.
	//
	// > You can call the [ListCheckResult](~~ListCheckResult~~) operation to obtain the check item ID.
	//
	// example:
	//
	// 16
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The list of check item IDs.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of instance IDs of the affected assets under the check item.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The task source. Valid values:
	//
	// - **YAO_CHI**: ApsaraDB console.
	//
	// example:
	//
	// YAO_CHI
	TaskSource *string `json:"TaskSource,omitempty" xml:"TaskSource,omitempty"`
}

func (s VerifyCheckInstanceResultRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckInstanceResultRequest) GoString() string {
	return s.String()
}

func (s *VerifyCheckInstanceResultRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *VerifyCheckInstanceResultRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *VerifyCheckInstanceResultRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *VerifyCheckInstanceResultRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *VerifyCheckInstanceResultRequest) GetTaskSource() *string {
	return s.TaskSource
}

func (s *VerifyCheckInstanceResultRequest) SetCheckId(v int64) *VerifyCheckInstanceResultRequest {
	s.CheckId = &v
	return s
}

func (s *VerifyCheckInstanceResultRequest) SetCheckIds(v []*int64) *VerifyCheckInstanceResultRequest {
	s.CheckIds = v
	return s
}

func (s *VerifyCheckInstanceResultRequest) SetClientToken(v string) *VerifyCheckInstanceResultRequest {
	s.ClientToken = &v
	return s
}

func (s *VerifyCheckInstanceResultRequest) SetInstanceIds(v []*string) *VerifyCheckInstanceResultRequest {
	s.InstanceIds = v
	return s
}

func (s *VerifyCheckInstanceResultRequest) SetTaskSource(v string) *VerifyCheckInstanceResultRequest {
	s.TaskSource = &v
	return s
}

func (s *VerifyCheckInstanceResultRequest) Validate() error {
	return dara.Validate(s)
}
