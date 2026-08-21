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
	// The list of upload task IDs or URLs that do not exist.
	NonExists []*string `json:"NonExists,omitempty" xml:"NonExists,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of URL upload information. For more information about the fields and descriptions, see [URLUploadInfo](https://help.aliyun.com/document_detail/52839.html).
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
	// The completion time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-01-01T01:11:01Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The creation time. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2019-01-01T01:01:01Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error_message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 24
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The upload task ID.
	//
	// example:
	//
	// 86c1925fba0****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The media ID.
	//
	// example:
	//
	// 93ab850b4f6f54b6e91d24d81d4****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 93ab850b4f654b6e91d24d81d44****
	RegisteredMediaId *string `json:"RegisteredMediaId,omitempty" xml:"RegisteredMediaId,omitempty"`
	// The status of the URL-based upload task. For more information about the status values and descriptions, see [Status](https://help.aliyun.com/document_detail/52839.html).
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The URL of the source video file.
	//
	// > A maximum of 100 records can be returned.
	//
	// example:
	//
	// http://****.mp4
	UploadURL *string `json:"UploadURL,omitempty" xml:"UploadURL,omitempty"`
	// The custom settings. The value is a JSON string. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
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

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) GetRegisteredMediaId() *string {
	return s.RegisteredMediaId
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

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetRegisteredMediaId(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.RegisteredMediaId = &v
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
