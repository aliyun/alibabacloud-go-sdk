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
	// The files.
	Files []*ListTransferFilesResponseBodyFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The number of entries to return on each page.
	//
	// Maximum value: 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of `NextToken` is a pagination token, which can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
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
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransferFilesResponseBodyFiles struct {
	// The URL of the file icon.
	//
	// >
	//
	// 	- For image file types (.png, .jpg, .jpeg, .gif, .webp, and .svg), the URL of the specific image is returned.
	//
	// 	- For other file types, the URL of the default image is returned.
	//
	// example:
	//
	// https://app-center-icon-pre-hangzhou.oss-cn-hangzhou.aliyuncs.com/tenant****
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The file ID.
	//
	// example:
	//
	// trf-a213msf****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The file name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the object stored in OSS.
	//
	// >
	//
	// 	- A value is returned for this parameter only when the object is stored in a custom OSS bucket.
	OssFileName *string `json:"OssFileName,omitempty" xml:"OssFileName,omitempty"`
	// The path of the object in the OSS bucket.
	//
	// >
	//
	// 	- A value is returned for this parameter only when the object is stored in a custom OSS bucket.
	//
	// example:
	//
	// transfer/1244234/****
	OssFilePath *string `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
	// The file size.
	//
	// example:
	//
	// 10853079
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The file status.
	//
	// Valid values:
	//
	// 	- DELETING
	//
	// 	- DELETED
	//
	// 	- UPLOADED
	//
	// example:
	//
	// DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The file type.
	//
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
