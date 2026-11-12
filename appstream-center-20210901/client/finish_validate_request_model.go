// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishValidateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *FinishValidateRequest
	GetAppId() *string
	SetAppVersionId(v string) *FinishValidateRequest
	GetAppVersionId() *string
	SetStatus(v bool) *FinishValidateRequest
	GetStatus() *bool
	SetTaskId(v string) *FinishValidateRequest
	GetTaskId() *string
}

type FinishValidateRequest struct {
	// The application ID used for exact match queries. If omitted, the results are not filtered by application ID.
	//
	// example:
	//
	// app-0aae4rxw7ql*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application version ID. **Required**. It must match the application version corresponding to the AdaptorId. This parameter is not the image version ID.
	//
	// example:
	//
	// apv-bp1example****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// The verification result.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tid-bp1example****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FinishValidateRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishValidateRequest) GoString() string {
	return s.String()
}

func (s *FinishValidateRequest) GetAppId() *string {
	return s.AppId
}

func (s *FinishValidateRequest) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *FinishValidateRequest) GetStatus() *bool {
	return s.Status
}

func (s *FinishValidateRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *FinishValidateRequest) SetAppId(v string) *FinishValidateRequest {
	s.AppId = &v
	return s
}

func (s *FinishValidateRequest) SetAppVersionId(v string) *FinishValidateRequest {
	s.AppVersionId = &v
	return s
}

func (s *FinishValidateRequest) SetStatus(v bool) *FinishValidateRequest {
	s.Status = &v
	return s
}

func (s *FinishValidateRequest) SetTaskId(v string) *FinishValidateRequest {
	s.TaskId = &v
	return s
}

func (s *FinishValidateRequest) Validate() error {
	return dara.Validate(s)
}
