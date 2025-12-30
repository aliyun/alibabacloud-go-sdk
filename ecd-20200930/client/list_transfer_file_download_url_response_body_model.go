// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransferFileDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTransferFileDownloadUrlResponseBody
	GetRequestId() *string
	SetUrls(v []*ListTransferFileDownloadUrlResponseBodyUrls) *ListTransferFileDownloadUrlResponseBody
	GetUrls() []*ListTransferFileDownloadUrlResponseBodyUrls
}

type ListTransferFileDownloadUrlResponseBody struct {
	// example:
	//
	// F1F01499-8F3D-5657-91AD-48177EB****
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Urls      []*ListTransferFileDownloadUrlResponseBodyUrls `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Repeated"`
}

func (s ListTransferFileDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransferFileDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransferFileDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransferFileDownloadUrlResponseBody) GetUrls() []*ListTransferFileDownloadUrlResponseBodyUrls {
	return s.Urls
}

func (s *ListTransferFileDownloadUrlResponseBody) SetRequestId(v string) *ListTransferFileDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransferFileDownloadUrlResponseBody) SetUrls(v []*ListTransferFileDownloadUrlResponseBodyUrls) *ListTransferFileDownloadUrlResponseBody {
	s.Urls = v
	return s
}

func (s *ListTransferFileDownloadUrlResponseBody) Validate() error {
	if s.Urls != nil {
		for _, item := range s.Urls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransferFileDownloadUrlResponseBodyUrls struct {
	// example:
	//
	// trf-i4pz8emx2k2fr****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// document.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// http://xxsy-transfer.oss-cn-beijing.aliyuncs.com/xxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListTransferFileDownloadUrlResponseBodyUrls) String() string {
	return dara.Prettify(s)
}

func (s ListTransferFileDownloadUrlResponseBodyUrls) GoString() string {
	return s.String()
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) GetFileId() *string {
	return s.FileId
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) GetFileName() *string {
	return s.FileName
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) GetStatus() *string {
	return s.Status
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) GetUrl() *string {
	return s.Url
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) SetFileId(v string) *ListTransferFileDownloadUrlResponseBodyUrls {
	s.FileId = &v
	return s
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) SetFileName(v string) *ListTransferFileDownloadUrlResponseBodyUrls {
	s.FileName = &v
	return s
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) SetStatus(v string) *ListTransferFileDownloadUrlResponseBodyUrls {
	s.Status = &v
	return s
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) SetUrl(v string) *ListTransferFileDownloadUrlResponseBodyUrls {
	s.Url = &v
	return s
}

func (s *ListTransferFileDownloadUrlResponseBodyUrls) Validate() error {
	return dara.Validate(s)
}
