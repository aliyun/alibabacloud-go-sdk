// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMiniGameOnlineStatusByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SyncMiniGameOnlineStatusByAppResponseBody
	GetAccessDeniedDetail() *string
	SetContent(v *SyncMiniGameOnlineStatusByAppResponseBodyContent) *SyncMiniGameOnlineStatusByAppResponseBody
	GetContent() *SyncMiniGameOnlineStatusByAppResponseBodyContent
	SetErrorCode(v string) *SyncMiniGameOnlineStatusByAppResponseBody
	GetErrorCode() *string
	SetRequestId(v string) *SyncMiniGameOnlineStatusByAppResponseBody
	GetRequestId() *string
	SetResultMsg(v string) *SyncMiniGameOnlineStatusByAppResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *SyncMiniGameOnlineStatusByAppResponseBody
	GetSuccess() *bool
}

type SyncMiniGameOnlineStatusByAppResponseBody struct {
	AccessDeniedDetail *string                                           `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Content            *SyncMiniGameOnlineStatusByAppResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	ErrorCode          *string                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg          *string                                           `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success            *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncMiniGameOnlineStatusByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncMiniGameOnlineStatusByAppResponseBody) GoString() string {
	return s.String()
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) GetContent() *SyncMiniGameOnlineStatusByAppResponseBodyContent {
	return s.Content
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) SetAccessDeniedDetail(v string) *SyncMiniGameOnlineStatusByAppResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) SetContent(v *SyncMiniGameOnlineStatusByAppResponseBodyContent) *SyncMiniGameOnlineStatusByAppResponseBody {
	s.Content = v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) SetErrorCode(v string) *SyncMiniGameOnlineStatusByAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) SetRequestId(v string) *SyncMiniGameOnlineStatusByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) SetResultMsg(v string) *SyncMiniGameOnlineStatusByAppResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) SetSuccess(v bool) *SyncMiniGameOnlineStatusByAppResponseBody {
	s.Success = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncMiniGameOnlineStatusByAppResponseBodyContent struct {
	CurrentOnlineStatus *int32  `json:"CurrentOnlineStatus,omitempty" xml:"CurrentOnlineStatus,omitempty"`
	MiniProgramCode     *string `json:"MiniProgramCode,omitempty" xml:"MiniProgramCode,omitempty"`
	StatusChanged       *bool   `json:"StatusChanged,omitempty" xml:"StatusChanged,omitempty"`
}

func (s SyncMiniGameOnlineStatusByAppResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s SyncMiniGameOnlineStatusByAppResponseBodyContent) GoString() string {
	return s.String()
}

func (s *SyncMiniGameOnlineStatusByAppResponseBodyContent) GetCurrentOnlineStatus() *int32 {
	return s.CurrentOnlineStatus
}

func (s *SyncMiniGameOnlineStatusByAppResponseBodyContent) GetMiniProgramCode() *string {
	return s.MiniProgramCode
}

func (s *SyncMiniGameOnlineStatusByAppResponseBodyContent) GetStatusChanged() *bool {
	return s.StatusChanged
}

func (s *SyncMiniGameOnlineStatusByAppResponseBodyContent) SetCurrentOnlineStatus(v int32) *SyncMiniGameOnlineStatusByAppResponseBodyContent {
	s.CurrentOnlineStatus = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBodyContent) SetMiniProgramCode(v string) *SyncMiniGameOnlineStatusByAppResponseBodyContent {
	s.MiniProgramCode = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBodyContent) SetStatusChanged(v bool) *SyncMiniGameOnlineStatusByAppResponseBodyContent {
	s.StatusChanged = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
