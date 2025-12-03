// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemEstimateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWorkitemEstimateRequest
	GetDescription() *string
	SetRecordUserIdentifier(v string) *CreateWorkitemEstimateRequest
	GetRecordUserIdentifier() *string
	SetSpentTime(v string) *CreateWorkitemEstimateRequest
	GetSpentTime() *string
	SetType(v string) *CreateWorkitemEstimateRequest
	GetType() *string
	SetWorkitemIdentifier(v string) *CreateWorkitemEstimateRequest
	GetWorkitemIdentifier() *string
}

type CreateWorkitemEstimateRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1316458xxxxx41068
	RecordUserIdentifier *string `json:"recordUserIdentifier,omitempty" xml:"recordUserIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21
	SpentTime *string `json:"spentTime,omitempty" xml:"spentTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9144ef6b72d8exxxxx9e61a4d0
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1e9903d8b3f1xxxxxf9286ef5
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemEstimateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemEstimateRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkitemEstimateRequest) GetRecordUserIdentifier() *string {
	return s.RecordUserIdentifier
}

func (s *CreateWorkitemEstimateRequest) GetSpentTime() *string {
	return s.SpentTime
}

func (s *CreateWorkitemEstimateRequest) GetType() *string {
	return s.Type
}

func (s *CreateWorkitemEstimateRequest) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *CreateWorkitemEstimateRequest) SetDescription(v string) *CreateWorkitemEstimateRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) SetRecordUserIdentifier(v string) *CreateWorkitemEstimateRequest {
	s.RecordUserIdentifier = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) SetSpentTime(v string) *CreateWorkitemEstimateRequest {
	s.SpentTime = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) SetType(v string) *CreateWorkitemEstimateRequest {
	s.Type = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) SetWorkitemIdentifier(v string) *CreateWorkitemEstimateRequest {
	s.WorkitemIdentifier = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) Validate() error {
	return dara.Validate(s)
}
