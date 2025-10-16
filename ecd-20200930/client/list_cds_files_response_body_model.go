// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCdsFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCdsFilesResponseBody
	GetCode() *string
	SetCount(v string) *ListCdsFilesResponseBody
	GetCount() *string
	SetFileModels(v []*ListCdsFilesResponseBodyFileModels) *ListCdsFilesResponseBody
	GetFileModels() []*ListCdsFilesResponseBodyFileModels
	SetMessage(v string) *ListCdsFilesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListCdsFilesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCdsFilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCdsFilesResponseBody
	GetSuccess() *bool
}

type ListCdsFilesResponseBody struct {
	// The result of the operation. A value of success indicates that the operation is successful. If the operation failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The files.
	FileModels []*ListCdsFilesResponseBodyFileModels `json:"FileModels,omitempty" xml:"FileModels,omitempty" type:"Repeated"`
	// The error message returned if the request failed. This parameter is empty if the value of Code is success.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token used for the next query. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// WyI2MzhmMjA2ZTFmZGZlZGVjZDk3OTRlNzVhZmIwMDFiZmM5NWQ3YTgwIiwibiIsIm4iLDEsLTEsMTY3MDMyNDMzNTAzMSwiNjM4ZjIwNmZjNDFkMzIwOTZmZWU0NGYxODkwY2I5ZjI0Mjg0NzM2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID. An ID is the unique identifier of the request.
	//
	// example:
	//
	// 40D86754-20FD-53DC-A9B8-25F7FECC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- <!-- -->
	//
	//     false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCdsFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCdsFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCdsFilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCdsFilesResponseBody) GetCount() *string {
	return s.Count
}

func (s *ListCdsFilesResponseBody) GetFileModels() []*ListCdsFilesResponseBodyFileModels {
	return s.FileModels
}

func (s *ListCdsFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCdsFilesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCdsFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCdsFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCdsFilesResponseBody) SetCode(v string) *ListCdsFilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCdsFilesResponseBody) SetCount(v string) *ListCdsFilesResponseBody {
	s.Count = &v
	return s
}

func (s *ListCdsFilesResponseBody) SetFileModels(v []*ListCdsFilesResponseBodyFileModels) *ListCdsFilesResponseBody {
	s.FileModels = v
	return s
}

func (s *ListCdsFilesResponseBody) SetMessage(v string) *ListCdsFilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCdsFilesResponseBody) SetNextToken(v string) *ListCdsFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCdsFilesResponseBody) SetRequestId(v string) *ListCdsFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCdsFilesResponseBody) SetSuccess(v bool) *ListCdsFilesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCdsFilesResponseBody) Validate() error {
	if s.FileModels != nil {
		for _, item := range s.FileModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCdsFilesResponseBodyFileModels struct {
	// The file category. PDS categorizes files based on their suffixes and MIME types. The following major categories are included: doc, image, audio, and video.
	//
	// example:
	//
	// image
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content type of the file.
	//
	// example:
	//
	// application/json
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The time when the file was created.
	//
	// example:
	//
	// 2022-09-06T07:27:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The file creator.
	//
	// example:
	//
	// user01
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The file description.
	//
	// example:
	//
	// test1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL that is used to download the file. The download URL is valid for only 15 minutes. If the URL is expired, you can call the GetFile operation to obtain the file.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899ef****?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The filename extension.
	//
	// example:
	//
	// pdf
	FileExtension *string `json:"FileExtension,omitempty" xml:"FileExtension,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 637725ff2f63db8470984e6c92c692b87d52****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file path.
	//
	// example:
	//
	// isv/1019236948660053/temp/
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The file type.
	//
	// example:
	//
	// file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The MD5 value of the file.
	//
	// example:
	//
	// 63c83ececb4e6926c51448fc5ecb****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The time when the file was last modified.
	//
	// example:
	//
	// 2022-09-06T07:27:08Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The user who modified the file.
	//
	// example:
	//
	// user02
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The file name.
	//
	// example:
	//
	// fvt-appcenterp4qwa
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the file was last opened.
	//
	// example:
	//
	// 2022-09-06T07:27:08Z
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// The timestamp that indicates the time when the file was last opened.
	//
	// example:
	//
	// 168951245231
	OpenTimeStamp *int64 `json:"OpenTimeStamp,omitempty" xml:"OpenTimeStamp,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// 3343213ff2f63db8470984e6c92c3213dfdw****
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The region ID You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SHA 1 file.
	//
	// example:
	//
	// EA4942AA8761213890A5C386F88E6464D2C3****
	Sha1 *string `json:"Sha1,omitempty" xml:"Sha1,omitempty"`
	// The file size. Unit: bits.
	//
	// example:
	//
	// 102400
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The URL of the thumbnail.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899ef****?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
}

func (s ListCdsFilesResponseBodyFileModels) String() string {
	return dara.Prettify(s)
}

func (s ListCdsFilesResponseBodyFileModels) GoString() string {
	return s.String()
}

func (s *ListCdsFilesResponseBodyFileModels) GetCategory() *string {
	return s.Category
}

func (s *ListCdsFilesResponseBodyFileModels) GetContentType() *string {
	return s.ContentType
}

func (s *ListCdsFilesResponseBodyFileModels) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCdsFilesResponseBodyFileModels) GetCreator() *string {
	return s.Creator
}

func (s *ListCdsFilesResponseBodyFileModels) GetDescription() *string {
	return s.Description
}

func (s *ListCdsFilesResponseBodyFileModels) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListCdsFilesResponseBodyFileModels) GetFileExtension() *string {
	return s.FileExtension
}

func (s *ListCdsFilesResponseBodyFileModels) GetFileId() *string {
	return s.FileId
}

func (s *ListCdsFilesResponseBodyFileModels) GetFilePath() *string {
	return s.FilePath
}

func (s *ListCdsFilesResponseBodyFileModels) GetFileType() *string {
	return s.FileType
}

func (s *ListCdsFilesResponseBodyFileModels) GetMd5() *string {
	return s.Md5
}

func (s *ListCdsFilesResponseBodyFileModels) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListCdsFilesResponseBodyFileModels) GetModifier() *string {
	return s.Modifier
}

func (s *ListCdsFilesResponseBodyFileModels) GetName() *string {
	return s.Name
}

func (s *ListCdsFilesResponseBodyFileModels) GetOpenTime() *string {
	return s.OpenTime
}

func (s *ListCdsFilesResponseBodyFileModels) GetOpenTimeStamp() *int64 {
	return s.OpenTimeStamp
}

func (s *ListCdsFilesResponseBodyFileModels) GetParentId() *string {
	return s.ParentId
}

func (s *ListCdsFilesResponseBodyFileModels) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCdsFilesResponseBodyFileModels) GetSha1() *string {
	return s.Sha1
}

func (s *ListCdsFilesResponseBodyFileModels) GetSize() *int64 {
	return s.Size
}

func (s *ListCdsFilesResponseBodyFileModels) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *ListCdsFilesResponseBodyFileModels) SetCategory(v string) *ListCdsFilesResponseBodyFileModels {
	s.Category = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetContentType(v string) *ListCdsFilesResponseBodyFileModels {
	s.ContentType = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetCreateTime(v string) *ListCdsFilesResponseBodyFileModels {
	s.CreateTime = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetCreator(v string) *ListCdsFilesResponseBodyFileModels {
	s.Creator = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetDescription(v string) *ListCdsFilesResponseBodyFileModels {
	s.Description = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetDownloadUrl(v string) *ListCdsFilesResponseBodyFileModels {
	s.DownloadUrl = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetFileExtension(v string) *ListCdsFilesResponseBodyFileModels {
	s.FileExtension = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetFileId(v string) *ListCdsFilesResponseBodyFileModels {
	s.FileId = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetFilePath(v string) *ListCdsFilesResponseBodyFileModels {
	s.FilePath = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetFileType(v string) *ListCdsFilesResponseBodyFileModels {
	s.FileType = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetMd5(v string) *ListCdsFilesResponseBodyFileModels {
	s.Md5 = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetModifiedTime(v string) *ListCdsFilesResponseBodyFileModels {
	s.ModifiedTime = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetModifier(v string) *ListCdsFilesResponseBodyFileModels {
	s.Modifier = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetName(v string) *ListCdsFilesResponseBodyFileModels {
	s.Name = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetOpenTime(v string) *ListCdsFilesResponseBodyFileModels {
	s.OpenTime = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetOpenTimeStamp(v int64) *ListCdsFilesResponseBodyFileModels {
	s.OpenTimeStamp = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetParentId(v string) *ListCdsFilesResponseBodyFileModels {
	s.ParentId = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetRegionId(v string) *ListCdsFilesResponseBodyFileModels {
	s.RegionId = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetSha1(v string) *ListCdsFilesResponseBodyFileModels {
	s.Sha1 = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetSize(v int64) *ListCdsFilesResponseBodyFileModels {
	s.Size = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) SetThumbnail(v string) *ListCdsFilesResponseBodyFileModels {
	s.Thumbnail = &v
	return s
}

func (s *ListCdsFilesResponseBodyFileModels) Validate() error {
	return dara.Validate(s)
}
