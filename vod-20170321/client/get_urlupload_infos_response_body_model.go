// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetURLUploadInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExists(v []*string) *GetURLUploadInfosResponseBody
	GetNonExists() []*string
	SetRequestId(v string) *GetURLUploadInfosResponseBody
	GetRequestId() *string
	SetURLUploadInfoList(v []*GetURLUploadInfosResponseBodyURLUploadInfoList) *GetURLUploadInfosResponseBody
	GetURLUploadInfoList() []*GetURLUploadInfosResponseBodyURLUploadInfoList
}

type GetURLUploadInfosResponseBody struct {
	// The job IDs or upload URLs that do not exist.
	NonExists []*string `json:"NonExists,omitempty" xml:"NonExists,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about URL-based upload jobs. For more information, see the "URLUploadInfo: the information about a URL-based upload job" section of the [Basic structures](https://help.aliyun.com/document_detail/52839.html) topic.
	URLUploadInfoList []*GetURLUploadInfosResponseBodyURLUploadInfoList `json:"URLUploadInfoList,omitempty" xml:"URLUploadInfoList,omitempty" type:"Repeated"`
}

func (s GetURLUploadInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetURLUploadInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetURLUploadInfosResponseBody) GetNonExists() []*string {
	return s.NonExists
}

func (s *GetURLUploadInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetURLUploadInfosResponseBody) GetURLUploadInfoList() []*GetURLUploadInfosResponseBodyURLUploadInfoList {
	return s.URLUploadInfoList
}

func (s *GetURLUploadInfosResponseBody) SetNonExists(v []*string) *GetURLUploadInfosResponseBody {
	s.NonExists = v
	return s
}

func (s *GetURLUploadInfosResponseBody) SetRequestId(v string) *GetURLUploadInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetURLUploadInfosResponseBody) SetURLUploadInfoList(v []*GetURLUploadInfosResponseBodyURLUploadInfoList) *GetURLUploadInfosResponseBody {
	s.URLUploadInfoList = v
	return s
}

func (s *GetURLUploadInfosResponseBody) Validate() error {
	if s.URLUploadInfoList != nil {
		for _, item := range s.URLUploadInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetURLUploadInfosResponseBodyURLUploadInfoList struct {
	// The time when the upload job was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-01T01:11:01Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the upload job was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-01T01:01:01Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// error_message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The size of the uploaded media file. Unit: byte.
	//
	// example:
	//
	// 24
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The ID of the upload job.
	//
	// example:
	//
	// 86c1925fba0****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the uploaded media file.
	//
	// example:
	//
	// 93ab850b4f6f54b6e91d24d81d4****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The status of the URL-based upload job. For more information about the valid values and value description of the parameter, see the "Status: the status of a video" section of the [Basic structures](https://help.aliyun.com/document_detail/52839.html) topic.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The upload URL of the source file.
	//
	// > A maximum of 100 URLs can be returned.
	//
	// example:
	//
	// http://****.mp4
	UploadURL *string `json:"UploadURL,omitempty" xml:"UploadURL,omitempty"`
	// The custom configurations. The value is a JSON string. For more information, see the "UserData: specifies the custom configurations for media upload" section of the [Request parameters](https://help.aliyun.com/document_detail/86952.html) topic.
	//
	// example:
	//
	// {"MessageCallback":"{"CallbackURL":"http://example.aliyundoc.com"}", "Extend":"{"localId":"***", "test":"www"}"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetURLUploadInfosResponseBodyURLUploadInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetURLUploadInfosResponseBodyURLUploadInfoList) GoString() string {
	return s.String()
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetFileSize() *string {
	return s.FileSize
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetJobId() *string {
	return s.JobId
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetMediaId() *string {
	return s.MediaId
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetStatus() *string {
	return s.Status
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetUploadURL() *string {
	return s.UploadURL
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetUserData() *string {
	return s.UserData
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetCompleteTime(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.CompleteTime = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetCreationTime(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetErrorCode(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.ErrorCode = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetErrorMessage(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetFileSize(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.FileSize = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetJobId(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.JobId = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetMediaId(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.MediaId = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetStatus(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.Status = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetUploadURL(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.UploadURL = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetUserData(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.UserData = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) Validate() error {
	return dara.Validate(s)
}
