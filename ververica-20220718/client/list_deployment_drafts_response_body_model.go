// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentDraftsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DeploymentDraft) *ListDeploymentDraftsResponseBody
	GetData() []*DeploymentDraft
	SetErrorCode(v string) *ListDeploymentDraftsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDeploymentDraftsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListDeploymentDraftsResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListDeploymentDraftsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDeploymentDraftsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDeploymentDraftsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDeploymentDraftsResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListDeploymentDraftsResponseBody
	GetTotalSize() *int32
}

type ListDeploymentDraftsResponseBody struct {
	Data []*DeploymentDraft `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 69
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListDeploymentDraftsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentDraftsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentDraftsResponseBody) GetData() []*DeploymentDraft {
	return s.Data
}

func (s *ListDeploymentDraftsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDeploymentDraftsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDeploymentDraftsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListDeploymentDraftsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDeploymentDraftsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentDraftsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentDraftsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDeploymentDraftsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListDeploymentDraftsResponseBody) SetData(v []*DeploymentDraft) *ListDeploymentDraftsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetErrorCode(v string) *ListDeploymentDraftsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetErrorMessage(v string) *ListDeploymentDraftsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetHttpCode(v int32) *ListDeploymentDraftsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetPageIndex(v int32) *ListDeploymentDraftsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetPageSize(v int32) *ListDeploymentDraftsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetRequestId(v string) *ListDeploymentDraftsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetSuccess(v bool) *ListDeploymentDraftsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) SetTotalSize(v int32) *ListDeploymentDraftsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListDeploymentDraftsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
