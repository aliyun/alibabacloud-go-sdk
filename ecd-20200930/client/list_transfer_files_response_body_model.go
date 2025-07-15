// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransferFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*ListTransferFilesResponseBodyFiles) *ListTransferFilesResponseBody
	GetFiles() []*ListTransferFilesResponseBodyFiles
	SetMaxResults(v int32) *ListTransferFilesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransferFilesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransferFilesResponseBody
	GetRequestId() *string
}

type ListTransferFilesResponseBody struct {
	Files []*ListTransferFilesResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTransferFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransferFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransferFilesResponseBody) GetFiles() []*ListTransferFilesResponseBodyFiles {
	return s.Files
}

func (s *ListTransferFilesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransferFilesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransferFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransferFilesResponseBody) SetFiles(v []*ListTransferFilesResponseBodyFiles) *ListTransferFilesResponseBody {
	s.Files = v
	return s
}

func (s *ListTransferFilesResponseBody) SetMaxResults(v int32) *ListTransferFilesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransferFilesResponseBody) SetNextToken(v string) *ListTransferFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransferFilesResponseBody) SetRequestId(v string) *ListTransferFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransferFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTransferFilesResponseBodyFiles struct {
	// example:
	//
	// https://app-center-icon-pre-hangzhou.oss-cn-hangzhou.aliyuncs.com/tenant****
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// example:
	//
	// trf-a213msf****
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssFileName *string `json:"OssFileName,omitempty" xml:"OssFileName,omitempty"`
	// example:
	//
	// transfer/1244234/****
	OssFilePath *string `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
	// example:
	//
	// 10853079
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// txt
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTransferFilesResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s ListTransferFilesResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *ListTransferFilesResponseBodyFiles) GetIconUrl() *string {
	return s.IconUrl
}

func (s *ListTransferFilesResponseBodyFiles) GetId() *string {
	return s.Id
}

func (s *ListTransferFilesResponseBodyFiles) GetName() *string {
	return s.Name
}

func (s *ListTransferFilesResponseBodyFiles) GetOssFileName() *string {
	return s.OssFileName
}

func (s *ListTransferFilesResponseBodyFiles) GetOssFilePath() *string {
	return s.OssFilePath
}

func (s *ListTransferFilesResponseBodyFiles) GetSize() *string {
	return s.Size
}

func (s *ListTransferFilesResponseBodyFiles) GetStatus() *string {
	return s.Status
}

func (s *ListTransferFilesResponseBodyFiles) GetType() *string {
	return s.Type
}

func (s *ListTransferFilesResponseBodyFiles) SetIconUrl(v string) *ListTransferFilesResponseBodyFiles {
	s.IconUrl = &v
	return s
}

func (s *ListTransferFilesResponseBodyFiles) SetId(v string) *ListTransferFilesResponseBodyFiles {
	s.Id = &v
	return s
}

func (s *ListTransferFilesResponseBodyFiles) SetName(v string) *ListTransferFilesResponseBodyFiles {
	s.Name = &v
	return s
}

func (s *ListTransferFilesResponseBodyFiles) SetOssFileName(v string) *ListTransferFilesResponseBodyFiles {
	s.OssFileName = &v
	return s
}

func (s *ListTransferFilesResponseBodyFiles) SetOssFilePath(v string) *ListTransferFilesResponseBodyFiles {
	s.OssFilePath = &v
	return s
}

func (s *ListTransferFilesResponseBodyFiles) SetSize(v string) *ListTransferFilesResponseBodyFiles {
	s.Size = &v
	return s
}

func (s *ListTransferFilesResponseBodyFiles) SetStatus(v string) *ListTransferFilesResponseBodyFiles {
	s.Status = &v
	return s
}

func (s *ListTransferFilesResponseBodyFiles) SetType(v string) *ListTransferFilesResponseBodyFiles {
	s.Type = &v
	return s
}

func (s *ListTransferFilesResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}
