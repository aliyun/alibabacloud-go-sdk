// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemAttachmentCreatemetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetWorkitemAttachmentCreatemetaResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetWorkitemAttachmentCreatemetaResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetWorkitemAttachmentCreatemetaResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetWorkitemAttachmentCreatemetaResponseBody
	GetSuccess() *string
	SetUploadInfo(v *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) *GetWorkitemAttachmentCreatemetaResponseBody
	GetUploadInfo() *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo
}

type GetWorkitemAttachmentCreatemetaResponseBody struct {
	// example:
	//
	// Invalid.IdNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Forbidden.UserNotInCurrentOrganization
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success    *string                                                `json:"success,omitempty" xml:"success,omitempty"`
	UploadInfo *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo `json:"uploadInfo,omitempty" xml:"uploadInfo,omitempty" type:"Struct"`
}

func (s GetWorkitemAttachmentCreatemetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemAttachmentCreatemetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) GetUploadInfo() *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo {
	return s.UploadInfo
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) SetErrorCode(v string) *GetWorkitemAttachmentCreatemetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) SetErrorMessage(v string) *GetWorkitemAttachmentCreatemetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) SetRequestId(v string) *GetWorkitemAttachmentCreatemetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) SetSuccess(v string) *GetWorkitemAttachmentCreatemetaResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) SetUploadInfo(v *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) *GetWorkitemAttachmentCreatemetaResponseBody {
	s.UploadInfo = v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBody) Validate() error {
	if s.UploadInfo != nil {
		if err := s.UploadInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo struct {
	// example:
	//
	// xxxxxxx
	Accessid *string `json:"accessid,omitempty" xml:"accessid,omitempty"`
	// example:
	//
	// ddd/dddd
	Dir *string `json:"dir,omitempty" xml:"dir,omitempty"`
	// example:
	//
	// xxxxx
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// xxxxxxx
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// example:
	//
	// xdWcrl/yTmIUA0kE7a3B0Ox4Vu8=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) GoString() string {
	return s.String()
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) GetAccessid() *string {
	return s.Accessid
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) GetDir() *string {
	return s.Dir
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) GetHost() *string {
	return s.Host
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) GetPolicy() *string {
	return s.Policy
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) GetSignature() *string {
	return s.Signature
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) SetAccessid(v string) *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo {
	s.Accessid = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) SetDir(v string) *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo {
	s.Dir = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) SetHost(v string) *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo {
	s.Host = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) SetPolicy(v string) *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo {
	s.Policy = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) SetSignature(v string) *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo {
	s.Signature = &v
	return s
}

func (s *GetWorkitemAttachmentCreatemetaResponseBodyUploadInfo) Validate() error {
	return dara.Validate(s)
}
