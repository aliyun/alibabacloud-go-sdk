// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFpShotFileList(v *ListFpShotFilesResponseBodyFpShotFileList) *ListFpShotFilesResponseBody
	GetFpShotFileList() *ListFpShotFilesResponseBodyFpShotFileList
	SetNextPageToken(v string) *ListFpShotFilesResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListFpShotFilesResponseBody
	GetRequestId() *string
}

type ListFpShotFilesResponseBody struct {
	// The media fingerprint files. For more information, see the "FpShotFile" section of the [Data types](https://help.aliyun.com/document_detail/29251.html) topic.
	FpShotFileList *ListFpShotFilesResponseBodyFpShotFileList `json:"FpShotFileList,omitempty" xml:"FpShotFileList,omitempty" type:"Struct"`
	// The returned value of NextPageToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ae0fd49c0840e14daf0d66a75b83****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFpShotFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFpShotFilesResponseBody) GetFpShotFileList() *ListFpShotFilesResponseBodyFpShotFileList {
	return s.FpShotFileList
}

func (s *ListFpShotFilesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListFpShotFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFpShotFilesResponseBody) SetFpShotFileList(v *ListFpShotFilesResponseBodyFpShotFileList) *ListFpShotFilesResponseBody {
	s.FpShotFileList = v
	return s
}

func (s *ListFpShotFilesResponseBody) SetNextPageToken(v string) *ListFpShotFilesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListFpShotFilesResponseBody) SetRequestId(v string) *ListFpShotFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFpShotFilesResponseBody) Validate() error {
	if s.FpShotFileList != nil {
		if err := s.FpShotFileList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFpShotFilesResponseBodyFpShotFileList struct {
	FpShotFile []*ListFpShotFilesResponseBodyFpShotFileListFpShotFile `json:"FpShotFile,omitempty" xml:"FpShotFile,omitempty" type:"Repeated"`
}

func (s ListFpShotFilesResponseBodyFpShotFileList) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotFilesResponseBodyFpShotFileList) GoString() string {
	return s.String()
}

func (s *ListFpShotFilesResponseBodyFpShotFileList) GetFpShotFile() []*ListFpShotFilesResponseBodyFpShotFileListFpShotFile {
	return s.FpShotFile
}

func (s *ListFpShotFilesResponseBodyFpShotFileList) SetFpShotFile(v []*ListFpShotFilesResponseBodyFpShotFileListFpShotFile) *ListFpShotFilesResponseBodyFpShotFileList {
	s.FpShotFile = v
	return s
}

func (s *ListFpShotFilesResponseBodyFpShotFileList) Validate() error {
	if s.FpShotFile != nil {
		for _, item := range s.FpShotFile {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFpShotFilesResponseBodyFpShotFileListFpShotFile struct {
	// The ID of the video file.
	//
	// example:
	//
	// 41e6536e4f2250e2e9bf26cdea19****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The information about the input file.
	InputFile *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The unique primary key of the input video.
	//
	// example:
	//
	// fb712a6890464059b1b2ea7c8647****
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The time when the media fingerprint file was stored. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	//
	// > This parameter is available only in the China (Beijing), China (Hangzhou), and China (Shanghai) regions.
	//
	// example:
	//
	// 2022-09-08T23:32:56Z
	StoreTime *string `json:"StoreTime,omitempty" xml:"StoreTime,omitempty"`
}

func (s ListFpShotFilesResponseBodyFpShotFileListFpShotFile) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotFilesResponseBodyFpShotFileListFpShotFile) GoString() string {
	return s.String()
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) GetFileId() *string {
	return s.FileId
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) GetInputFile() *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile {
	return s.InputFile
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) GetStoreTime() *string {
	return s.StoreTime
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) SetFileId(v string) *ListFpShotFilesResponseBodyFpShotFileListFpShotFile {
	s.FileId = &v
	return s
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) SetInputFile(v *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) *ListFpShotFilesResponseBodyFpShotFileListFpShotFile {
	s.InputFile = v
	return s
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) SetPrimaryKey(v string) *ListFpShotFilesResponseBodyFpShotFileListFpShotFile {
	s.PrimaryKey = &v
	return s
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) SetStoreTime(v string) *ListFpShotFilesResponseBodyFpShotFileListFpShotFile {
	s.StoreTime = &v
	return s
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFile) Validate() error {
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile struct {
	// The name of the OSS bucket in which the input file is stored.
	//
	// example:
	//
	// example-bucket-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the OSS region in which the input file resides.
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

func (s ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) GoString() string {
	return s.String()
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) GetLocation() *string {
	return s.Location
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) GetObject() *string {
	return s.Object
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) SetBucket(v string) *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile {
	s.Bucket = &v
	return s
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) SetLocation(v string) *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile {
	s.Location = &v
	return s
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) SetObject(v string) *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile {
	s.Object = &v
	return s
}

func (s *ListFpShotFilesResponseBodyFpShotFileListFpShotFileInputFile) Validate() error {
	return dara.Validate(s)
}
