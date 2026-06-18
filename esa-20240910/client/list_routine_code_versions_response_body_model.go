// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineCodeVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersions(v []*ListRoutineCodeVersionsResponseBodyCodeVersions) *ListRoutineCodeVersionsResponseBody
	GetCodeVersions() []*ListRoutineCodeVersionsResponseBodyCodeVersions
	SetPageNumber(v int64) *ListRoutineCodeVersionsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRoutineCodeVersionsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListRoutineCodeVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRoutineCodeVersionsResponseBody
	GetTotalCount() *int64
}

type ListRoutineCodeVersionsResponseBody struct {
	// The list of Edge Routine code versions.
	CodeVersions []*ListRoutineCodeVersionsResponseBodyCodeVersions `json:"CodeVersions,omitempty" xml:"CodeVersions,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoutineCodeVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCodeVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutineCodeVersionsResponseBody) GetCodeVersions() []*ListRoutineCodeVersionsResponseBodyCodeVersions {
	return s.CodeVersions
}

func (s *ListRoutineCodeVersionsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRoutineCodeVersionsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRoutineCodeVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutineCodeVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRoutineCodeVersionsResponseBody) SetCodeVersions(v []*ListRoutineCodeVersionsResponseBodyCodeVersions) *ListRoutineCodeVersionsResponseBody {
	s.CodeVersions = v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) SetPageNumber(v int64) *ListRoutineCodeVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) SetPageSize(v int64) *ListRoutineCodeVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) SetRequestId(v string) *ListRoutineCodeVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) SetTotalCount(v int64) *ListRoutineCodeVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBody) Validate() error {
	if s.CodeVersions != nil {
		for _, item := range s.CodeVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoutineCodeVersionsResponseBodyCodeVersions struct {
	// The build ID of the code version.
	//
	// example:
	//
	// 25801233
	BuildId *int64 `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// The description of the code version.
	//
	// example:
	//
	// test desc
	CodeDescription *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	// The code version number.
	//
	// example:
	//
	// 1723599747213377175
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The configuration items of the code version.
	ConfOptions *ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions `json:"ConfOptions,omitempty" xml:"ConfOptions,omitempty" type:"Struct"`
	// The time when the code version was created.
	//
	// example:
	//
	// 2024-04-16T09:42:47Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The deployment environment. Valid values:
	//
	// - **staging**: staging environment.
	//
	// - **production**: production environment.
	//
	// example:
	//
	// staging
	DeployEnv *string `json:"DeployEnv,omitempty" xml:"DeployEnv,omitempty"`
	// The additional information about the code version.
	//
	// example:
	//
	// {\\"approver\\":[\\"348678\\",\\"111133\\",\\"411544\\"]}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// Indicates whether environment variables exist.
	HasEnvVars *bool `json:"HasEnvVars,omitempty" xml:"HasEnvVars,omitempty"`
	// The status of the code version.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRoutineCodeVersionsResponseBodyCodeVersions) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCodeVersionsResponseBodyCodeVersions) GoString() string {
	return s.String()
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetBuildId() *int64 {
	return s.BuildId
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetConfOptions() *ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions {
	return s.ConfOptions
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetDeployEnv() *string {
	return s.DeployEnv
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetHasEnvVars() *bool {
	return s.HasEnvVars
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) GetStatus() *string {
	return s.Status
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetBuildId(v int64) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.BuildId = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetCodeDescription(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.CodeDescription = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetCodeVersion(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.CodeVersion = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetConfOptions(v *ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.ConfOptions = v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetCreateTime(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.CreateTime = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetDeployEnv(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.DeployEnv = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetExtraInfo(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.ExtraInfo = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetHasEnvVars(v bool) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.HasEnvVars = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) SetStatus(v string) *ListRoutineCodeVersionsResponseBodyCodeVersions {
	s.Status = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersions) Validate() error {
	if s.ConfOptions != nil {
		if err := s.ConfOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions struct {
	// The NotFoundStrategy configuration item of the code version.
	//
	// example:
	//
	// SinglePageApplication
	NotFoundStrategy *string `json:"NotFoundStrategy,omitempty" xml:"NotFoundStrategy,omitempty"`
}

func (s ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions) GoString() string {
	return s.String()
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions) GetNotFoundStrategy() *string {
	return s.NotFoundStrategy
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions) SetNotFoundStrategy(v string) *ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions {
	s.NotFoundStrategy = &v
	return s
}

func (s *ListRoutineCodeVersionsResponseBodyCodeVersionsConfOptions) Validate() error {
	return dara.Validate(s)
}
