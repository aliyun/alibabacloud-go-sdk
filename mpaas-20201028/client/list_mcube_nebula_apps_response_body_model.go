// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMcubeNebulaAppsResult(v *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) *ListMcubeNebulaAppsResponseBody
	GetListMcubeNebulaAppsResult() *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult
	SetRequestId(v string) *ListMcubeNebulaAppsResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeNebulaAppsResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeNebulaAppsResponseBody
	GetResultMessage() *string
}

type ListMcubeNebulaAppsResponseBody struct {
	ListMcubeNebulaAppsResult *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult `json:"ListMcubeNebulaAppsResult,omitempty" xml:"ListMcubeNebulaAppsResult,omitempty" type:"Struct"`
	RequestId                 *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                *string                                                   `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage             *string                                                   `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeNebulaAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaAppsResponseBody) GetListMcubeNebulaAppsResult() *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	return s.ListMcubeNebulaAppsResult
}

func (s *ListMcubeNebulaAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeNebulaAppsResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeNebulaAppsResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeNebulaAppsResponseBody) SetListMcubeNebulaAppsResult(v *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) *ListMcubeNebulaAppsResponseBody {
	s.ListMcubeNebulaAppsResult = v
	return s
}

func (s *ListMcubeNebulaAppsResponseBody) SetRequestId(v string) *ListMcubeNebulaAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBody) SetResultCode(v string) *ListMcubeNebulaAppsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBody) SetResultMessage(v string) *ListMcubeNebulaAppsResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBody) Validate() error {
	if s.ListMcubeNebulaAppsResult != nil {
		if err := s.ListMcubeNebulaAppsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult struct {
	CurrentPage    *int32                                                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ErrorCode      *string                                                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HasMore        *bool                                                                     `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	NebulaAppInfos []*ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos `json:"NebulaAppInfos,omitempty" xml:"NebulaAppInfos,omitempty" type:"Repeated"`
	PageSize       *int32                                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg      *string                                                                   `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success        *bool                                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetNebulaAppInfos() []*ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos {
	return s.NebulaAppInfos
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetCurrentPage(v int32) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetErrorCode(v string) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetHasMore(v bool) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.HasMore = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetNebulaAppInfos(v []*ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.NebulaAppInfos = v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetPageSize(v int32) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.PageSize = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetRequestId(v string) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.RequestId = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetResultMsg(v string) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetSuccess(v bool) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.Success = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) SetTotalCount(v int64) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult {
	s.TotalCount = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResult) Validate() error {
	if s.NebulaAppInfos != nil {
		for _, item := range s.NebulaAppInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos struct {
	H5Id   *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
}

func (s ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos) GetH5Id() *string {
	return s.H5Id
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos) GetH5Name() *string {
	return s.H5Name
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos) SetH5Id(v string) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos {
	s.H5Id = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos) SetH5Name(v string) *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos {
	s.H5Name = &v
	return s
}

func (s *ListMcubeNebulaAppsResponseBodyListMcubeNebulaAppsResultNebulaAppInfos) Validate() error {
	return dara.Validate(s)
}
