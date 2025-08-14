// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDNAFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileList(v []*ListDNAFilesResponseBodyFileList) *ListDNAFilesResponseBody
	GetFileList() []*ListDNAFilesResponseBodyFileList
	SetNextPageToken(v string) *ListDNAFilesResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListDNAFilesResponseBody
	GetRequestId() *string
}

type ListDNAFilesResponseBody struct {
	// The queried files.
	FileList []*ListDNAFilesResponseBodyFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ae0fd49c0840e14daf0d66a75b83****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2AE89FA5-E620-56C7-9B80-75D09757385A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDNAFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDNAFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDNAFilesResponseBody) GetFileList() []*ListDNAFilesResponseBodyFileList {
	return s.FileList
}

func (s *ListDNAFilesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDNAFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDNAFilesResponseBody) SetFileList(v []*ListDNAFilesResponseBodyFileList) *ListDNAFilesResponseBody {
	s.FileList = v
	return s
}

func (s *ListDNAFilesResponseBody) SetNextPageToken(v string) *ListDNAFilesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListDNAFilesResponseBody) SetRequestId(v string) *ListDNAFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDNAFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDNAFilesResponseBodyFileList struct {
	// The Object Storage Service (OSS) information about the input file.
	InputFile *ListDNAFilesResponseBodyFileListInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The primary key of the file.
	//
	// example:
	//
	// ae0fd49c0840e14daf0d66a75b83****
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
}

func (s ListDNAFilesResponseBodyFileList) String() string {
	return dara.Prettify(s)
}

func (s ListDNAFilesResponseBodyFileList) GoString() string {
	return s.String()
}

func (s *ListDNAFilesResponseBodyFileList) GetInputFile() *ListDNAFilesResponseBodyFileListInputFile {
	return s.InputFile
}

func (s *ListDNAFilesResponseBodyFileList) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ListDNAFilesResponseBodyFileList) SetInputFile(v *ListDNAFilesResponseBodyFileListInputFile) *ListDNAFilesResponseBodyFileList {
	s.InputFile = v
	return s
}

func (s *ListDNAFilesResponseBodyFileList) SetPrimaryKey(v string) *ListDNAFilesResponseBodyFileList {
	s.PrimaryKey = &v
	return s
}

func (s *ListDNAFilesResponseBodyFileList) Validate() error {
	return dara.Validate(s)
}

type ListDNAFilesResponseBodyFileListInputFile struct {
	// The name of the OSS bucket in which the input file is stored.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the input file resides.
	//
	// example:
	//
	// oss-cn-beijing
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the input file.
	//
	// example:
	//
	// example-****.mp4
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s ListDNAFilesResponseBodyFileListInputFile) String() string {
	return dara.Prettify(s)
}

func (s ListDNAFilesResponseBodyFileListInputFile) GoString() string {
	return s.String()
}

func (s *ListDNAFilesResponseBodyFileListInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *ListDNAFilesResponseBodyFileListInputFile) GetLocation() *string {
	return s.Location
}

func (s *ListDNAFilesResponseBodyFileListInputFile) GetObject() *string {
	return s.Object
}

func (s *ListDNAFilesResponseBodyFileListInputFile) SetBucket(v string) *ListDNAFilesResponseBodyFileListInputFile {
	s.Bucket = &v
	return s
}

func (s *ListDNAFilesResponseBodyFileListInputFile) SetLocation(v string) *ListDNAFilesResponseBodyFileListInputFile {
	s.Location = &v
	return s
}

func (s *ListDNAFilesResponseBodyFileListInputFile) SetObject(v string) *ListDNAFilesResponseBodyFileListInputFile {
	s.Object = &v
	return s
}

func (s *ListDNAFilesResponseBodyFileListInputFile) Validate() error {
	return dara.Validate(s)
}
