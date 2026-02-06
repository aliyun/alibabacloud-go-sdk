// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScriptReviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubmitScriptReviewRequest
	GetDescription() *string
	SetFrom(v string) *SubmitScriptReviewRequest
	GetFrom() *string
	SetInstanceId(v string) *SubmitScriptReviewRequest
	GetInstanceId() *string
	SetScriptId(v string) *SubmitScriptReviewRequest
	GetScriptId() *string
}

type SubmitScriptReviewRequest struct {
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	From        *string `json:"From,omitempty" xml:"From,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b0f35dd1-0337-402e-9c4f-3a6c2426950a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s SubmitScriptReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitScriptReviewRequest) GoString() string {
	return s.String()
}

func (s *SubmitScriptReviewRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitScriptReviewRequest) GetFrom() *string {
	return s.From
}

func (s *SubmitScriptReviewRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SubmitScriptReviewRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *SubmitScriptReviewRequest) SetDescription(v string) *SubmitScriptReviewRequest {
	s.Description = &v
	return s
}

func (s *SubmitScriptReviewRequest) SetFrom(v string) *SubmitScriptReviewRequest {
	s.From = &v
	return s
}

func (s *SubmitScriptReviewRequest) SetInstanceId(v string) *SubmitScriptReviewRequest {
	s.InstanceId = &v
	return s
}

func (s *SubmitScriptReviewRequest) SetScriptId(v string) *SubmitScriptReviewRequest {
	s.ScriptId = &v
	return s
}

func (s *SubmitScriptReviewRequest) Validate() error {
	return dara.Validate(s)
}
