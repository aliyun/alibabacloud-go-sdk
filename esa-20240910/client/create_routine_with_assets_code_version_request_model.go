// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineWithAssetsCodeVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildId(v int64) *CreateRoutineWithAssetsCodeVersionRequest
	GetBuildId() *int64
	SetCodeDescription(v string) *CreateRoutineWithAssetsCodeVersionRequest
	GetCodeDescription() *string
	SetConfOptions(v *CreateRoutineWithAssetsCodeVersionRequestConfOptions) *CreateRoutineWithAssetsCodeVersionRequest
	GetConfOptions() *CreateRoutineWithAssetsCodeVersionRequestConfOptions
	SetExtraInfo(v string) *CreateRoutineWithAssetsCodeVersionRequest
	GetExtraInfo() *string
	SetName(v string) *CreateRoutineWithAssetsCodeVersionRequest
	GetName() *string
}

type CreateRoutineWithAssetsCodeVersionRequest struct {
	BuildId         *int64                                                `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	CodeDescription *string                                               `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	ConfOptions     *CreateRoutineWithAssetsCodeVersionRequestConfOptions `json:"ConfOptions,omitempty" xml:"ConfOptions,omitempty" type:"Struct"`
	ExtraInfo       *string                                               `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRoutineWithAssetsCodeVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineWithAssetsCodeVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) GetBuildId() *int64 {
	return s.BuildId
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) GetConfOptions() *CreateRoutineWithAssetsCodeVersionRequestConfOptions {
	return s.ConfOptions
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) GetName() *string {
	return s.Name
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) SetBuildId(v int64) *CreateRoutineWithAssetsCodeVersionRequest {
	s.BuildId = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) SetCodeDescription(v string) *CreateRoutineWithAssetsCodeVersionRequest {
	s.CodeDescription = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) SetConfOptions(v *CreateRoutineWithAssetsCodeVersionRequestConfOptions) *CreateRoutineWithAssetsCodeVersionRequest {
	s.ConfOptions = v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) SetExtraInfo(v string) *CreateRoutineWithAssetsCodeVersionRequest {
	s.ExtraInfo = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) SetName(v string) *CreateRoutineWithAssetsCodeVersionRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionRequest) Validate() error {
	return dara.Validate(s)
}

type CreateRoutineWithAssetsCodeVersionRequestConfOptions struct {
	NotFoundStrategy *string `json:"NotFoundStrategy,omitempty" xml:"NotFoundStrategy,omitempty"`
}

func (s CreateRoutineWithAssetsCodeVersionRequestConfOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineWithAssetsCodeVersionRequestConfOptions) GoString() string {
	return s.String()
}

func (s *CreateRoutineWithAssetsCodeVersionRequestConfOptions) GetNotFoundStrategy() *string {
	return s.NotFoundStrategy
}

func (s *CreateRoutineWithAssetsCodeVersionRequestConfOptions) SetNotFoundStrategy(v string) *CreateRoutineWithAssetsCodeVersionRequestConfOptions {
	s.NotFoundStrategy = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionRequestConfOptions) Validate() error {
	return dara.Validate(s)
}
