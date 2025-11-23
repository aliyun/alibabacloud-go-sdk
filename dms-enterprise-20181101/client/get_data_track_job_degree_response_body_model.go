// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackJobDegreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataTrackJobDegreeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataTrackJobDegreeResponseBody
	GetErrorMessage() *string
	SetJobDegree(v *GetDataTrackJobDegreeResponseBodyJobDegree) *GetDataTrackJobDegreeResponseBody
	GetJobDegree() *GetDataTrackJobDegreeResponseBodyJobDegree
	SetRequestId(v string) *GetDataTrackJobDegreeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataTrackJobDegreeResponseBody
	GetSuccess() *bool
}

type GetDataTrackJobDegreeResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress details of the data tracking task.
	JobDegree *GetDataTrackJobDegreeResponseBodyJobDegree `json:"JobDegree,omitempty" xml:"JobDegree,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C51420E3-144A-4A94-B473-8662FCF4AD10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataTrackJobDegreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobDegreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobDegreeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataTrackJobDegreeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataTrackJobDegreeResponseBody) GetJobDegree() *GetDataTrackJobDegreeResponseBodyJobDegree {
	return s.JobDegree
}

func (s *GetDataTrackJobDegreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataTrackJobDegreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataTrackJobDegreeResponseBody) SetErrorCode(v string) *GetDataTrackJobDegreeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBody) SetErrorMessage(v string) *GetDataTrackJobDegreeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBody) SetJobDegree(v *GetDataTrackJobDegreeResponseBodyJobDegree) *GetDataTrackJobDegreeResponseBody {
	s.JobDegree = v
	return s
}

func (s *GetDataTrackJobDegreeResponseBody) SetRequestId(v string) *GetDataTrackJobDegreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBody) SetSuccess(v bool) *GetDataTrackJobDegreeResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBody) Validate() error {
	if s.JobDegree != nil {
		if err := s.JobDegree.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataTrackJobDegreeResponseBodyJobDegree struct {
	// The progress of binary log download. Valid values: 0 to 1. A value of 1 indicates that binary log download is complete.
	//
	// example:
	//
	// 1
	DownloadCompletionDegree *float64 `json:"DownloadCompletionDegree,omitempty" xml:"DownloadCompletionDegree,omitempty"`
	// The progress of binary log parsing. Valid values: 0 to 1. A value of 1 indicates that binary log parsing is complete.
	//
	// example:
	//
	// 1
	FilterCompletionDegree *float64 `json:"FilterCompletionDegree,omitempty" xml:"FilterCompletionDegree,omitempty"`
	// The status of the data tracking task. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **LISTING**: The binary logs are being obtained.
	//
	// 	- **LIST_SUCCESS**: The binary logs are successfully obtained.
	//
	// 	- **DOWNLOADING**: The binary logs are being downloaded.
	//
	// 	- **DOWNLOAD_FAIL**: The binary logs failed to be downloaded.
	//
	// 	- **DOWNLOAD_SUCCESS**: The binary logs are successfully downloaded.
	//
	// 	- **FILTERING**: The binary logs are being parsed.
	//
	// 	- **FILTER_FAIL**: The binary logs failed to be parsed.
	//
	// 	- **FILTER_SUCCESS**: The binary logs are successfully parsed.
	//
	// example:
	//
	// FILTER_SUCCESS
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The progress of binary log obtaining. Valid values: 0 to 1. A value of 1 indicates that binary log obtaining is complete.
	//
	// example:
	//
	// 1
	ListCompletionDegree *float64 `json:"ListCompletionDegree,omitempty" xml:"ListCompletionDegree,omitempty"`
	// The description of the task status.
	//
	// example:
	//
	// searching success
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
}

func (s GetDataTrackJobDegreeResponseBodyJobDegree) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobDegreeResponseBodyJobDegree) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) GetDownloadCompletionDegree() *float64 {
	return s.DownloadCompletionDegree
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) GetFilterCompletionDegree() *float64 {
	return s.FilterCompletionDegree
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) GetListCompletionDegree() *float64 {
	return s.ListCompletionDegree
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) SetDownloadCompletionDegree(v float64) *GetDataTrackJobDegreeResponseBodyJobDegree {
	s.DownloadCompletionDegree = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) SetFilterCompletionDegree(v float64) *GetDataTrackJobDegreeResponseBodyJobDegree {
	s.FilterCompletionDegree = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) SetJobStatus(v string) *GetDataTrackJobDegreeResponseBodyJobDegree {
	s.JobStatus = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) SetListCompletionDegree(v float64) *GetDataTrackJobDegreeResponseBodyJobDegree {
	s.ListCompletionDegree = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) SetStatusDesc(v string) *GetDataTrackJobDegreeResponseBodyJobDegree {
	s.StatusDesc = &v
	return s
}

func (s *GetDataTrackJobDegreeResponseBodyJobDegree) Validate() error {
	return dara.Validate(s)
}
