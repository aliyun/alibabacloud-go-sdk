// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDesensStatusListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DsgQueryDesensStatusListResponseBodyData) *DsgQueryDesensStatusListResponseBody
	GetData() *DsgQueryDesensStatusListResponseBodyData
	SetErrorCode(v string) *DsgQueryDesensStatusListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgQueryDesensStatusListResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgQueryDesensStatusListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgQueryDesensStatusListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgQueryDesensStatusListResponseBody
	GetSuccess() *bool
}

type DsgQueryDesensStatusListResponseBody struct {
	Data *DsgQueryDesensStatusListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400010
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// AASFDFSDFG-DFSDF-DFSDFD-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgQueryDesensStatusListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDesensStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *DsgQueryDesensStatusListResponseBody) GetData() *DsgQueryDesensStatusListResponseBodyData {
	return s.Data
}

func (s *DsgQueryDesensStatusListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgQueryDesensStatusListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgQueryDesensStatusListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgQueryDesensStatusListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgQueryDesensStatusListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgQueryDesensStatusListResponseBody) SetData(v *DsgQueryDesensStatusListResponseBodyData) *DsgQueryDesensStatusListResponseBody {
	s.Data = v
	return s
}

func (s *DsgQueryDesensStatusListResponseBody) SetErrorCode(v string) *DsgQueryDesensStatusListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBody) SetErrorMessage(v string) *DsgQueryDesensStatusListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBody) SetHttpStatusCode(v int32) *DsgQueryDesensStatusListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBody) SetRequestId(v string) *DsgQueryDesensStatusListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBody) SetSuccess(v bool) *DsgQueryDesensStatusListResponseBody {
	s.Success = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DsgQueryDesensStatusListResponseBodyData struct {
	PageData []*DsgQueryDesensStatusListResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DsgQueryDesensStatusListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDesensStatusListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DsgQueryDesensStatusListResponseBodyData) GetPageData() []*DsgQueryDesensStatusListResponseBodyDataPageData {
	return s.PageData
}

func (s *DsgQueryDesensStatusListResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgQueryDesensStatusListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgQueryDesensStatusListResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DsgQueryDesensStatusListResponseBodyData) SetPageData(v []*DsgQueryDesensStatusListResponseBodyDataPageData) *DsgQueryDesensStatusListResponseBodyData {
	s.PageData = v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyData) SetPageNumber(v int32) *DsgQueryDesensStatusListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyData) SetPageSize(v int32) *DsgQueryDesensStatusListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyData) SetTotalCount(v int32) *DsgQueryDesensStatusListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DsgQueryDesensStatusListResponseBodyDataPageData struct {
	// example:
	//
	// 1
	DesensStatus *int32 `json:"DesensStatus,omitempty" xml:"DesensStatus,omitempty"`
	// example:
	//
	// 1
	HandleSpace *string `json:"HandleSpace,omitempty" xml:"HandleSpace,omitempty"`
	// example:
	//
	// 56207
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test_space
	WorkspaceIdentifier *string `json:"WorkspaceIdentifier,omitempty" xml:"WorkspaceIdentifier,omitempty"`
	// example:
	//
	// test_space
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s DsgQueryDesensStatusListResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDesensStatusListResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) GetDesensStatus() *int32 {
	return s.DesensStatus
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) GetHandleSpace() *string {
	return s.HandleSpace
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) GetId() *int64 {
	return s.Id
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) GetWorkspaceIdentifier() *string {
	return s.WorkspaceIdentifier
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) SetDesensStatus(v int32) *DsgQueryDesensStatusListResponseBodyDataPageData {
	s.DesensStatus = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) SetHandleSpace(v string) *DsgQueryDesensStatusListResponseBodyDataPageData {
	s.HandleSpace = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) SetId(v int64) *DsgQueryDesensStatusListResponseBodyDataPageData {
	s.Id = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) SetWorkspaceIdentifier(v string) *DsgQueryDesensStatusListResponseBodyDataPageData {
	s.WorkspaceIdentifier = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) SetWorkspaceName(v string) *DsgQueryDesensStatusListResponseBodyDataPageData {
	s.WorkspaceName = &v
	return s
}

func (s *DsgQueryDesensStatusListResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
