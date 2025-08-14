// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrlUploadInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExists(v []*string) *GetUrlUploadInfosResponseBody
	GetNonExists() []*string
	SetRequestId(v string) *GetUrlUploadInfosResponseBody
	GetRequestId() *string
	SetURLUploadInfoList(v []*GetUrlUploadInfosResponseBodyURLUploadInfoList) *GetUrlUploadInfosResponseBody
	GetURLUploadInfoList() []*GetUrlUploadInfosResponseBodyURLUploadInfoList
}

type GetUrlUploadInfosResponseBody struct {
	// The job IDs or upload URLs that do not exist.
	NonExists []*string `json:"NonExists,omitempty" xml:"NonExists,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about URL-based upload jobs.
	URLUploadInfoList []*GetUrlUploadInfosResponseBodyURLUploadInfoList `json:"URLUploadInfoList,omitempty" xml:"URLUploadInfoList,omitempty" type:"Repeated"`
}

func (s GetUrlUploadInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUrlUploadInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetUrlUploadInfosResponseBody) GetNonExists() []*string {
	return s.NonExists
}

func (s *GetUrlUploadInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUrlUploadInfosResponseBody) GetURLUploadInfoList() []*GetUrlUploadInfosResponseBodyURLUploadInfoList {
	return s.URLUploadInfoList
}

func (s *GetUrlUploadInfosResponseBody) SetNonExists(v []*string) *GetUrlUploadInfosResponseBody {
	s.NonExists = v
	return s
}

func (s *GetUrlUploadInfosResponseBody) SetRequestId(v string) *GetUrlUploadInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUrlUploadInfosResponseBody) SetURLUploadInfoList(v []*GetUrlUploadInfosResponseBodyURLUploadInfoList) *GetUrlUploadInfosResponseBody {
	s.URLUploadInfoList = v
	return s
}

func (s *GetUrlUploadInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUrlUploadInfosResponseBodyURLUploadInfoList struct {
	// The time when the upload job was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-26 21:47:37
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the upload job was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-07T10:03:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error code returned if the upload job failed.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the upload job failed.
	//
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 64610
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The ID of the upload job.
	//
	// example:
	//
	// 3829500c0fef429fa4ec1680b122d***
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the uploaded media file.
	//
	// example:
	//
	// 5014ca70f08171ecbf940764a0fd6***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The status of the upload job. For more information about the valid values of the parameter, see the "Status: the status of a URL-based upload job" section of the [Basic data types](https://help.aliyun.com/document_detail/52839.html) topic.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The upload URL of the source file.
	//
	// >  A maximum of 100 URLs can be returned.
	//
	// example:
	//
	// http://****.mp4
	UploadURL *string `json:"UploadURL,omitempty" xml:"UploadURL,omitempty"`
	// The user data. The value is a JSON string.
	//
	// example:
	//
	// {"MessageCallback":"{"CallbackURL":"http://example.aliyundoc.com"}", "Extend":"{"localId":"***", "test":"www"}"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetUrlUploadInfosResponseBodyURLUploadInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetUrlUploadInfosResponseBodyURLUploadInfoList) GoString() string {
	return s.String()
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetFileSize() *string {
	return s.FileSize
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetJobId() *string {
	return s.JobId
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetMediaId() *string {
	return s.MediaId
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetUploadURL() *string {
	return s.UploadURL
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) GetUserData() *string {
	return s.UserData
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetCompleteTime(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.CompleteTime = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetCreationTime(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetErrorCode(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.ErrorCode = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetErrorMessage(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetFileSize(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.FileSize = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetJobId(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.JobId = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetMediaId(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.MediaId = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetStatus(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.Status = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetUploadURL(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.UploadURL = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) SetUserData(v string) *GetUrlUploadInfosResponseBodyURLUploadInfoList {
	s.UserData = &v
	return s
}

func (s *GetUrlUploadInfosResponseBodyURLUploadInfoList) Validate() error {
	return dara.Validate(s)
}
