// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachments(v []*ListWorkitemAttachmentsResponseBodyAttachments) *ListWorkitemAttachmentsResponseBody
	GetAttachments() []*ListWorkitemAttachmentsResponseBodyAttachments
	SetErrorCode(v string) *ListWorkitemAttachmentsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListWorkitemAttachmentsResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ListWorkitemAttachmentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkitemAttachmentsResponseBody
	GetSuccess() *bool
}

type ListWorkitemAttachmentsResponseBody struct {
	Attachments []*ListWorkitemAttachmentsResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListWorkitemAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkitemAttachmentsResponseBody) GetAttachments() []*ListWorkitemAttachmentsResponseBodyAttachments {
	return s.Attachments
}

func (s *ListWorkitemAttachmentsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkitemAttachmentsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListWorkitemAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkitemAttachmentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkitemAttachmentsResponseBody) SetAttachments(v []*ListWorkitemAttachmentsResponseBodyAttachments) *ListWorkitemAttachmentsResponseBody {
	s.Attachments = v
	return s
}

func (s *ListWorkitemAttachmentsResponseBody) SetErrorCode(v string) *ListWorkitemAttachmentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBody) SetErrorMsg(v string) *ListWorkitemAttachmentsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBody) SetRequestId(v string) *ListWorkitemAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBody) SetSuccess(v bool) *ListWorkitemAttachmentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBody) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkitemAttachmentsResponseBodyAttachments struct {
	// example:
	//
	// 237109
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// dflkjlsdddsdl234lkjfg
	FileIdentifier *string `json:"fileIdentifier,omitempty" xml:"fileIdentifier,omitempty"`
	// example:
	//
	// Artifacts_1565193_1.tgz
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// tgz
	FileSuffix *string `json:"fileSuffix,omitempty" xml:"fileSuffix,omitempty"`
	// example:
	//
	// 1545726028000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 50
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// http://grace-share.oss-cn-hangzhou.aliyuncs.com/qf%2Fheap.bin?Expires=1675750082&OSSAccessKeyId=LTAI5t8irN2Wu3BGrBpffZue&Signature=RqRUEuHiwW8wuahYz6CenHaWWs4%3D
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListWorkitemAttachmentsResponseBodyAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemAttachmentsResponseBodyAttachments) GoString() string {
	return s.String()
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) GetFileIdentifier() *string {
	return s.FileIdentifier
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) GetFileName() *string {
	return s.FileName
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) GetFileSuffix() *string {
	return s.FileSuffix
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) GetSize() *string {
	return s.Size
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) GetUrl() *string {
	return s.Url
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) SetCreator(v string) *ListWorkitemAttachmentsResponseBodyAttachments {
	s.Creator = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) SetFileIdentifier(v string) *ListWorkitemAttachmentsResponseBodyAttachments {
	s.FileIdentifier = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) SetFileName(v string) *ListWorkitemAttachmentsResponseBodyAttachments {
	s.FileName = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) SetFileSuffix(v string) *ListWorkitemAttachmentsResponseBodyAttachments {
	s.FileSuffix = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) SetGmtCreate(v int64) *ListWorkitemAttachmentsResponseBodyAttachments {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) SetSize(v string) *ListWorkitemAttachmentsResponseBodyAttachments {
	s.Size = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) SetUrl(v string) *ListWorkitemAttachmentsResponseBodyAttachments {
	s.Url = &v
	return s
}

func (s *ListWorkitemAttachmentsResponseBodyAttachments) Validate() error {
	return dara.Validate(s)
}
