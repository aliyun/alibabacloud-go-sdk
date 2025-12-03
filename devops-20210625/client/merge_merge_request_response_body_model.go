// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeMergeRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *MergeMergeRequestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *MergeMergeRequestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *MergeMergeRequestResponseBody
	GetRequestId() *string
	SetResult(v *MergeMergeRequestResponseBodyResult) *MergeMergeRequestResponseBody
	GetResult() *MergeMergeRequestResponseBodyResult
	SetSuccess(v bool) *MergeMergeRequestResponseBody
	GetSuccess() *bool
}

type MergeMergeRequestResponseBody struct {
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
	// 9ED5E382-3A58-51E4-8A81-CE25D1756025
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *MergeMergeRequestResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MergeMergeRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MergeMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *MergeMergeRequestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *MergeMergeRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MergeMergeRequestResponseBody) GetResult() *MergeMergeRequestResponseBodyResult {
	return s.Result
}

func (s *MergeMergeRequestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MergeMergeRequestResponseBody) SetErrorCode(v string) *MergeMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetErrorMessage(v string) *MergeMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetRequestId(v string) *MergeMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetResult(v *MergeMergeRequestResponseBodyResult) *MergeMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *MergeMergeRequestResponseBody) SetSuccess(v bool) *MergeMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *MergeMergeRequestResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MergeMergeRequestResponseBodyResult struct {
	BizId          *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	LocalId        *int64  `json:"localId,omitempty" xml:"localId,omitempty"`
	MergedRevision *string `json:"mergedRevision,omitempty" xml:"mergedRevision,omitempty"`
	ProjectId      *int64  `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s MergeMergeRequestResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResult) GetBizId() *string {
	return s.BizId
}

func (s *MergeMergeRequestResponseBodyResult) GetLocalId() *int64 {
	return s.LocalId
}

func (s *MergeMergeRequestResponseBodyResult) GetMergedRevision() *string {
	return s.MergedRevision
}

func (s *MergeMergeRequestResponseBodyResult) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MergeMergeRequestResponseBodyResult) GetResult() *bool {
	return s.Result
}

func (s *MergeMergeRequestResponseBodyResult) SetBizId(v string) *MergeMergeRequestResponseBodyResult {
	s.BizId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetLocalId(v int64) *MergeMergeRequestResponseBodyResult {
	s.LocalId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergedRevision(v string) *MergeMergeRequestResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetProjectId(v int64) *MergeMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetResult(v bool) *MergeMergeRequestResponseBodyResult {
	s.Result = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
