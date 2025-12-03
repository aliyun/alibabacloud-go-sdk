// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestPatchSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListMergeRequestPatchSetsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMergeRequestPatchSetsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListMergeRequestPatchSetsResponseBody
	GetRequestId() *string
	SetResult(v []*ListMergeRequestPatchSetsResponseBodyResult) *ListMergeRequestPatchSetsResponseBody
	GetResult() []*ListMergeRequestPatchSetsResponseBodyResult
	SetSuccess(v bool) *ListMergeRequestPatchSetsResponseBody
	GetSuccess() *bool
}

type ListMergeRequestPatchSetsResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListMergeRequestPatchSetsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMergeRequestPatchSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestPatchSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMergeRequestPatchSetsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMergeRequestPatchSetsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMergeRequestPatchSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMergeRequestPatchSetsResponseBody) GetResult() []*ListMergeRequestPatchSetsResponseBodyResult {
	return s.Result
}

func (s *ListMergeRequestPatchSetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMergeRequestPatchSetsResponseBody) SetErrorCode(v string) *ListMergeRequestPatchSetsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBody) SetErrorMessage(v string) *ListMergeRequestPatchSetsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBody) SetRequestId(v string) *ListMergeRequestPatchSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBody) SetResult(v []*ListMergeRequestPatchSetsResponseBodyResult) *ListMergeRequestPatchSetsResponseBody {
	s.Result = v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBody) SetSuccess(v bool) *ListMergeRequestPatchSetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMergeRequestPatchSetsResponseBodyResult struct {
	// example:
	//
	// 1a072f5367c21f9de3464b8c0ee8546e47764d2d
	CommitId *string `json:"commitId,omitempty" xml:"commitId,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 513fcfd81a9142d2bb0db4f72c0aa15b
	PatchSetBizId *string `json:"patchSetBizId,omitempty" xml:"patchSetBizId,omitempty"`
	PatchSetName  *string `json:"patchSetName,omitempty" xml:"patchSetName,omitempty"`
	// example:
	//
	// 1
	PatchSetNo *int64 `json:"patchSetNo,omitempty" xml:"patchSetNo,omitempty"`
	// example:
	//
	// MERGE_SOURCE
	RelatedMergeItemType *string `json:"relatedMergeItemType,omitempty" xml:"relatedMergeItemType,omitempty"`
	// example:
	//
	// 1a072f53
	ShortCommitId *string `json:"shortCommitId,omitempty" xml:"shortCommitId,omitempty"`
}

func (s ListMergeRequestPatchSetsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestPatchSetsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) GetCommitId() *string {
	return s.CommitId
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) GetPatchSetBizId() *string {
	return s.PatchSetBizId
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) GetPatchSetName() *string {
	return s.PatchSetName
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) GetPatchSetNo() *int64 {
	return s.PatchSetNo
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) GetRelatedMergeItemType() *string {
	return s.RelatedMergeItemType
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) GetShortCommitId() *string {
	return s.ShortCommitId
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) SetCommitId(v string) *ListMergeRequestPatchSetsResponseBodyResult {
	s.CommitId = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) SetCreatedAt(v string) *ListMergeRequestPatchSetsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) SetPatchSetBizId(v string) *ListMergeRequestPatchSetsResponseBodyResult {
	s.PatchSetBizId = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) SetPatchSetName(v string) *ListMergeRequestPatchSetsResponseBodyResult {
	s.PatchSetName = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) SetPatchSetNo(v int64) *ListMergeRequestPatchSetsResponseBodyResult {
	s.PatchSetNo = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) SetRelatedMergeItemType(v string) *ListMergeRequestPatchSetsResponseBodyResult {
	s.RelatedMergeItemType = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) SetShortCommitId(v string) *ListMergeRequestPatchSetsResponseBodyResult {
	s.ShortCommitId = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
