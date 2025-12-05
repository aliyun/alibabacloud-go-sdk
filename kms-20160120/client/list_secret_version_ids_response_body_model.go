// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretVersionIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSecretVersionIdsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSecretVersionIdsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSecretVersionIdsResponseBody
	GetRequestId() *string
	SetSecretName(v string) *ListSecretVersionIdsResponseBody
	GetSecretName() *string
	SetTotalCount(v int32) *ListSecretVersionIdsResponseBody
	GetTotalCount() *int32
	SetVersionIds(v *ListSecretVersionIdsResponseBodyVersionIds) *ListSecretVersionIdsResponseBody
	GetVersionIds() *ListSecretVersionIdsResponseBodyVersionIds
}

type ListSecretVersionIdsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5b75d8b1-5b6a-4ec0-8e0c-c08befdfad47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of secret versions.
	VersionIds *ListSecretVersionIdsResponseBodyVersionIds `json:"VersionIds,omitempty" xml:"VersionIds,omitempty" type:"Struct"`
}

func (s ListSecretVersionIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretVersionIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecretVersionIdsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecretVersionIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretVersionIdsResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *ListSecretVersionIdsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecretVersionIdsResponseBody) GetVersionIds() *ListSecretVersionIdsResponseBodyVersionIds {
	return s.VersionIds
}

func (s *ListSecretVersionIdsResponseBody) SetPageNumber(v int32) *ListSecretVersionIdsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetPageSize(v int32) *ListSecretVersionIdsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetRequestId(v string) *ListSecretVersionIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetSecretName(v string) *ListSecretVersionIdsResponseBody {
	s.SecretName = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetTotalCount(v int32) *ListSecretVersionIdsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecretVersionIdsResponseBody) SetVersionIds(v *ListSecretVersionIdsResponseBodyVersionIds) *ListSecretVersionIdsResponseBody {
	s.VersionIds = v
	return s
}

func (s *ListSecretVersionIdsResponseBody) Validate() error {
	if s.VersionIds != nil {
		if err := s.VersionIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecretVersionIdsResponseBodyVersionIds struct {
	VersionId []*ListSecretVersionIdsResponseBodyVersionIdsVersionId `json:"VersionId,omitempty" xml:"VersionId,omitempty" type:"Repeated"`
}

func (s ListSecretVersionIdsResponseBodyVersionIds) String() string {
	return dara.Prettify(s)
}

func (s ListSecretVersionIdsResponseBodyVersionIds) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBodyVersionIds) GetVersionId() []*ListSecretVersionIdsResponseBodyVersionIdsVersionId {
	return s.VersionId
}

func (s *ListSecretVersionIdsResponseBodyVersionIds) SetVersionId(v []*ListSecretVersionIdsResponseBodyVersionIdsVersionId) *ListSecretVersionIdsResponseBodyVersionIds {
	s.VersionId = v
	return s
}

func (s *ListSecretVersionIdsResponseBodyVersionIds) Validate() error {
	if s.VersionId != nil {
		for _, item := range s.VersionId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecretVersionIdsResponseBodyVersionIdsVersionId struct {
	// The time when the secret version was created.
	//
	// example:
	//
	// 2020-02-21T15:39:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// 00000000000000000000000000000000203
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that mark the secret version.
	VersionStages *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
}

func (s ListSecretVersionIdsResponseBodyVersionIdsVersionId) String() string {
	return dara.Prettify(s)
}

func (s ListSecretVersionIdsResponseBodyVersionIdsVersionId) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) GetVersionId() *string {
	return s.VersionId
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) GetVersionStages() *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages {
	return s.VersionStages
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) SetCreateTime(v string) *ListSecretVersionIdsResponseBodyVersionIdsVersionId {
	s.CreateTime = &v
	return s
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) SetVersionId(v string) *ListSecretVersionIdsResponseBodyVersionIdsVersionId {
	s.VersionId = &v
	return s
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) SetVersionStages(v *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) *ListSecretVersionIdsResponseBodyVersionIdsVersionId {
	s.VersionStages = v
	return s
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionId) Validate() error {
	if s.VersionStages != nil {
		if err := s.VersionStages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" type:"Repeated"`
}

func (s ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) String() string {
	return dara.Prettify(s)
}

func (s ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) GetVersionStage() []*string {
	return s.VersionStage
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) SetVersionStage(v []*string) *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages {
	s.VersionStage = v
	return s
}

func (s *ListSecretVersionIdsResponseBodyVersionIdsVersionIdVersionStages) Validate() error {
	return dara.Validate(s)
}
