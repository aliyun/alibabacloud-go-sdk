// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListFileUploadResponseBodyData) *ListFileUploadResponseBody
	GetData() []*ListFileUploadResponseBodyData
	SetErrorCode(v string) *ListFileUploadResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListFileUploadResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListFileUploadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFileUploadResponseBody
	GetSuccess() *bool
}

type ListFileUploadResponseBody struct {
	Data []*ListFileUploadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFileUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileUploadResponseBody) GetData() []*ListFileUploadResponseBodyData {
	return s.Data
}

func (s *ListFileUploadResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListFileUploadResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListFileUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileUploadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFileUploadResponseBody) SetData(v []*ListFileUploadResponseBodyData) *ListFileUploadResponseBody {
	s.Data = v
	return s
}

func (s *ListFileUploadResponseBody) SetErrorCode(v string) *ListFileUploadResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListFileUploadResponseBody) SetErrorMessage(v string) *ListFileUploadResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListFileUploadResponseBody) SetRequestId(v string) *ListFileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileUploadResponseBody) SetSuccess(v bool) *ListFileUploadResponseBody {
	s.Success = &v
	return s
}

func (s *ListFileUploadResponseBody) Validate() error {
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

type ListFileUploadResponseBodyData struct {
	// example:
	//
	// 16738266********
	AliyunParentUid *string `json:"AliyunParentUid,omitempty" xml:"AliyunParentUid,omitempty"`
	// example:
	//
	// 20372822********
	AliyunUid    *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// example:
	//
	// TextReport
	FileCategory *string `json:"FileCategory,omitempty" xml:"FileCategory,omitempty"`
	// example:
	//
	// Agent
	FileFrom *string `json:"FileFrom,omitempty" xml:"FileFrom,omitempty"`
	// example:
	//
	// f-8*******01m
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// samele_report.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 7453
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// md
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtCreated           *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// h8r********4fch
	SessionId      *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UploadLocation *string `json:"UploadLocation,omitempty" xml:"UploadLocation,omitempty"`
}

func (s ListFileUploadResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFileUploadResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFileUploadResponseBodyData) GetAliyunParentUid() *string {
	return s.AliyunParentUid
}

func (s *ListFileUploadResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListFileUploadResponseBodyData) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *ListFileUploadResponseBodyData) GetFileCategory() *string {
	return s.FileCategory
}

func (s *ListFileUploadResponseBodyData) GetFileFrom() *string {
	return s.FileFrom
}

func (s *ListFileUploadResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *ListFileUploadResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *ListFileUploadResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ListFileUploadResponseBodyData) GetFileType() *string {
	return s.FileType
}

func (s *ListFileUploadResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListFileUploadResponseBodyData) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *ListFileUploadResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *ListFileUploadResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *ListFileUploadResponseBodyData) GetUploadLocation() *string {
	return s.UploadLocation
}

func (s *ListFileUploadResponseBodyData) SetAliyunParentUid(v string) *ListFileUploadResponseBodyData {
	s.AliyunParentUid = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetAliyunUid(v string) *ListFileUploadResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetDownloadLink(v string) *ListFileUploadResponseBodyData {
	s.DownloadLink = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetFileCategory(v string) *ListFileUploadResponseBodyData {
	s.FileCategory = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetFileFrom(v string) *ListFileUploadResponseBodyData {
	s.FileFrom = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetFileId(v string) *ListFileUploadResponseBodyData {
	s.FileId = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetFileName(v string) *ListFileUploadResponseBodyData {
	s.FileName = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetFileSize(v int64) *ListFileUploadResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetFileType(v string) *ListFileUploadResponseBodyData {
	s.FileType = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetGmtCreated(v string) *ListFileUploadResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetIntranetDownloadLink(v string) *ListFileUploadResponseBodyData {
	s.IntranetDownloadLink = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetRegion(v string) *ListFileUploadResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetSessionId(v string) *ListFileUploadResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *ListFileUploadResponseBodyData) SetUploadLocation(v string) *ListFileUploadResponseBodyData {
	s.UploadLocation = &v
	return s
}

func (s *ListFileUploadResponseBodyData) Validate() error {
	return dara.Validate(s)
}
