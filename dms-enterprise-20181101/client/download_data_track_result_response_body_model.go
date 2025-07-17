// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDataTrackResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadKeyId(v string) *DownloadDataTrackResultResponseBody
	GetDownloadKeyId() *string
	SetErrorCode(v string) *DownloadDataTrackResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DownloadDataTrackResultResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DownloadDataTrackResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadDataTrackResultResponseBody
	GetSuccess() *bool
}

type DownloadDataTrackResultResponseBody struct {
	// The ID of the download key, which is used to download the parsing result of the data tracking task.
	//
	// example:
	//
	// e23dd7ec-a19f-4a69-8eb3-8ffd26e6****
	DownloadKeyId *string `json:"DownloadKeyId,omitempty" xml:"DownloadKeyId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// B43AD641-49C2-5299-9E06-1B37EC1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadDataTrackResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadDataTrackResultResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDataTrackResultResponseBody) GetDownloadKeyId() *string {
	return s.DownloadKeyId
}

func (s *DownloadDataTrackResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DownloadDataTrackResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DownloadDataTrackResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadDataTrackResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadDataTrackResultResponseBody) SetDownloadKeyId(v string) *DownloadDataTrackResultResponseBody {
	s.DownloadKeyId = &v
	return s
}

func (s *DownloadDataTrackResultResponseBody) SetErrorCode(v string) *DownloadDataTrackResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DownloadDataTrackResultResponseBody) SetErrorMessage(v string) *DownloadDataTrackResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DownloadDataTrackResultResponseBody) SetRequestId(v string) *DownloadDataTrackResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadDataTrackResultResponseBody) SetSuccess(v bool) *DownloadDataTrackResultResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadDataTrackResultResponseBody) Validate() error {
	return dara.Validate(s)
}
