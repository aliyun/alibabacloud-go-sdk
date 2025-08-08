// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeProjectAttachmentsResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeProjectAttachmentsResponseBodyResult) *DescribeProjectAttachmentsResponseBody
	GetResult() []*DescribeProjectAttachmentsResponseBodyResult
	SetSuccess(v bool) *DescribeProjectAttachmentsResponseBody
	GetSuccess() *bool
}

type DescribeProjectAttachmentsResponseBody struct {
	// example:
	//
	// e03a9f78-7b40-4fb3-a015-350913e37e3f
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeProjectAttachmentsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectAttachmentsResponseBody) GetResult() []*DescribeProjectAttachmentsResponseBodyResult {
	return s.Result
}

func (s *DescribeProjectAttachmentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProjectAttachmentsResponseBody) SetRequestId(v string) *DescribeProjectAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBody) SetResult(v []*DescribeProjectAttachmentsResponseBodyResult) *DescribeProjectAttachmentsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectAttachmentsResponseBody) SetSuccess(v bool) *DescribeProjectAttachmentsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProjectAttachmentsResponseBodyResult struct {
	// example:
	//
	// Mzc4NDAtODQ3MjY4MzI=
	AttachmentToken *string `json:"AttachmentToken,omitempty" xml:"AttachmentToken,omitempty"`
	// example:
	//
	// File
	AttachmentType *string `json:"AttachmentType,omitempty" xml:"AttachmentType,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// http://delivery-center.oss-cn-shanghai.aliyuncs.com/6a8****0e2/e0a***f3.jpg?Expires=1589334682&OSSAccessKeyId=wI2r*********&Signature=JWB39pUxs7RCqrcw58qXPEGu6M0%3D
	FileLink *string `json:"FileLink,omitempty" xml:"FileLink,omitempty"`
	// example:
	//
	// 1589334682404
	FileLinkGmtExpired *int64 `json:"FileLinkGmtExpired,omitempty" xml:"FileLinkGmtExpired,omitempty"`
	// example:
	//
	// f8-test-perview.jpeg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 109124
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// jpg
	FileSuffix *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
	// example:
	//
	// 1587968858000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 8472
	NodeId   *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// 45261111****
	Operator     *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// example:
	//
	// Provider
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	// example:
	//
	// 3
	StepNo *int32 `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
}

func (s DescribeProjectAttachmentsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectAttachmentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetAttachmentToken() *string {
	return s.AttachmentToken
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetAttachmentType() *string {
	return s.AttachmentType
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetFileLink() *string {
	return s.FileLink
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetFileLinkGmtExpired() *int64 {
	return s.FileLinkGmtExpired
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetFileName() *string {
	return s.FileName
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetFileSuffix() *string {
	return s.FileSuffix
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetNodeId() *int64 {
	return s.NodeId
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetOperator() *int64 {
	return s.Operator
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetOperatorRole() *string {
	return s.OperatorRole
}

func (s *DescribeProjectAttachmentsResponseBodyResult) GetStepNo() *int32 {
	return s.StepNo
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetAttachmentToken(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.AttachmentToken = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetAttachmentType(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.AttachmentType = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetContent(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileLink(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileLink = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileLinkGmtExpired(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileLinkGmtExpired = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileName(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileName = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileSize(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileSuffix(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileSuffix = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetGmtCreate(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetNodeId(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.NodeId = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetNodeName(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.NodeName = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetOperator(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.Operator = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetOperatorName(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetOperatorRole(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.OperatorRole = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetStepNo(v int32) *DescribeProjectAttachmentsResponseBodyResult {
	s.StepNo = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
