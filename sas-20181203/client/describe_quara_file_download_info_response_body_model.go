// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuaraFileDownloadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *DescribeQuaraFileDownloadInfoResponseBody
	GetDownloadUrl() *string
	SetMd5(v string) *DescribeQuaraFileDownloadInfoResponseBody
	GetMd5() *string
	SetPath(v string) *DescribeQuaraFileDownloadInfoResponseBody
	GetPath() *string
	SetQuaraFileId(v int32) *DescribeQuaraFileDownloadInfoResponseBody
	GetQuaraFileId() *int32
	SetRequestId(v string) *DescribeQuaraFileDownloadInfoResponseBody
	GetRequestId() *string
	SetTag(v string) *DescribeQuaraFileDownloadInfoResponseBody
	GetTag() *string
	SetUuid(v string) *DescribeQuaraFileDownloadInfoResponseBody
	GetUuid() *string
}

type DescribeQuaraFileDownloadInfoResponseBody struct {
	// The URL that is used to download the file. The URL is valid for five minutes.
	//
	// example:
	//
	// https://xxxxxxxx.oss-cn-hangzhou-1.aliyuncs.com/xxxxx/xxxxxxxxxxxxxx?Expires=1671448125&OSSAccessKeyId=xxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The MD5 hash value of the quarantined file.
	//
	// example:
	//
	// bb62ef1311bc564377a0378d3axxxxxx
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The file path.
	//
	// example:
	//
	// /etc/test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The ID of the quarantined file.
	//
	// example:
	//
	// 123
	QuaraFileId *int32 `json:"QuaraFileId,omitempty" xml:"QuaraFileId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 29874225-EAAC-5415-8501-32DD20FXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag that is added to the related alert.
	//
	// example:
	//
	// 6d4ff40a22b15c86adecf2aa48xxxxx
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeQuaraFileDownloadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuaraFileDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) GetMd5() *string {
	return s.Md5
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) GetPath() *string {
	return s.Path
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) GetQuaraFileId() *int32 {
	return s.QuaraFileId
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) GetTag() *string {
	return s.Tag
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) SetDownloadUrl(v string) *DescribeQuaraFileDownloadInfoResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) SetMd5(v string) *DescribeQuaraFileDownloadInfoResponseBody {
	s.Md5 = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) SetPath(v string) *DescribeQuaraFileDownloadInfoResponseBody {
	s.Path = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) SetQuaraFileId(v int32) *DescribeQuaraFileDownloadInfoResponseBody {
	s.QuaraFileId = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) SetRequestId(v string) *DescribeQuaraFileDownloadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) SetTag(v string) *DescribeQuaraFileDownloadInfoResponseBody {
	s.Tag = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) SetUuid(v string) *DescribeQuaraFileDownloadInfoResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeQuaraFileDownloadInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
