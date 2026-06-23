// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineCodeVersionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildId(v int64) *GetRoutineCodeVersionInfoResponseBody
	GetBuildId() *int64
	SetCodeDescription(v string) *GetRoutineCodeVersionInfoResponseBody
	GetCodeDescription() *string
	SetCodeVersion(v string) *GetRoutineCodeVersionInfoResponseBody
	GetCodeVersion() *string
	SetConfOptions(v *GetRoutineCodeVersionInfoResponseBodyConfOptions) *GetRoutineCodeVersionInfoResponseBody
	GetConfOptions() *GetRoutineCodeVersionInfoResponseBodyConfOptions
	SetCreateTime(v string) *GetRoutineCodeVersionInfoResponseBody
	GetCreateTime() *string
	SetExtraInfo(v string) *GetRoutineCodeVersionInfoResponseBody
	GetExtraInfo() *string
	SetHasAssets(v bool) *GetRoutineCodeVersionInfoResponseBody
	GetHasAssets() *bool
	SetRequestId(v string) *GetRoutineCodeVersionInfoResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetRoutineCodeVersionInfoResponseBody
	GetStatus() *string
}

type GetRoutineCodeVersionInfoResponseBody struct {
	// The build ID of the code.
	//
	// example:
	//
	// 26035169
	BuildId *int64 `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// The description of the code version.
	//
	// example:
	//
	// code desc version unstable
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	// The code version number.
	//
	// example:
	//
	// 1710120201067203242
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The list of configuration items for the code version.
	ConfOptions *GetRoutineCodeVersionInfoResponseBodyConfOptions `json:"ConfOptions,omitempty" xml:"ConfOptions,omitempty" type:"Struct"`
	// The time when the code version was created.
	//
	// example:
	//
	// 2025-08-04T01:54:19Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The additional information about the code version. The value is in JSON string format.
	//
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// Indicates whether the code version contains asset files.
	//
	// example:
	//
	// true
	HasAssets *bool `json:"HasAssets,omitempty" xml:"HasAssets,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the code version.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRoutineCodeVersionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineCodeVersionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetBuildId() *int64 {
	return s.BuildId
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetConfOptions() *GetRoutineCodeVersionInfoResponseBodyConfOptions {
	return s.ConfOptions
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetHasAssets() *bool {
	return s.HasAssets
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineCodeVersionInfoResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetBuildId(v int64) *GetRoutineCodeVersionInfoResponseBody {
	s.BuildId = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetCodeDescription(v string) *GetRoutineCodeVersionInfoResponseBody {
	s.CodeDescription = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetCodeVersion(v string) *GetRoutineCodeVersionInfoResponseBody {
	s.CodeVersion = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetConfOptions(v *GetRoutineCodeVersionInfoResponseBodyConfOptions) *GetRoutineCodeVersionInfoResponseBody {
	s.ConfOptions = v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetCreateTime(v string) *GetRoutineCodeVersionInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetExtraInfo(v string) *GetRoutineCodeVersionInfoResponseBody {
	s.ExtraInfo = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetHasAssets(v bool) *GetRoutineCodeVersionInfoResponseBody {
	s.HasAssets = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetRequestId(v string) *GetRoutineCodeVersionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) SetStatus(v string) *GetRoutineCodeVersionInfoResponseBody {
	s.Status = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBody) Validate() error {
	if s.ConfOptions != nil {
		if err := s.ConfOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRoutineCodeVersionInfoResponseBodyConfOptions struct {
	// The NotFoundStrategy configuration item.
	//
	// example:
	//
	// SinglePageApplication
	NotFoundStrategy *string `json:"NotFoundStrategy,omitempty" xml:"NotFoundStrategy,omitempty"`
}

func (s GetRoutineCodeVersionInfoResponseBodyConfOptions) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineCodeVersionInfoResponseBodyConfOptions) GoString() string {
	return s.String()
}

func (s *GetRoutineCodeVersionInfoResponseBodyConfOptions) GetNotFoundStrategy() *string {
	return s.NotFoundStrategy
}

func (s *GetRoutineCodeVersionInfoResponseBodyConfOptions) SetNotFoundStrategy(v string) *GetRoutineCodeVersionInfoResponseBodyConfOptions {
	s.NotFoundStrategy = &v
	return s
}

func (s *GetRoutineCodeVersionInfoResponseBodyConfOptions) Validate() error {
	return dara.Validate(s)
}
