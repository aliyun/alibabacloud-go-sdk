// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadBackupSetStorageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDownloadBackupSetStorageInfoResponseBody
	GetCode() *string
	SetData(v *DescribeDownloadBackupSetStorageInfoResponseBodyData) *DescribeDownloadBackupSetStorageInfoResponseBody
	GetData() *DescribeDownloadBackupSetStorageInfoResponseBodyData
	SetErrCode(v string) *DescribeDownloadBackupSetStorageInfoResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDownloadBackupSetStorageInfoResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeDownloadBackupSetStorageInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDownloadBackupSetStorageInfoResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDownloadBackupSetStorageInfoResponseBody
	GetSuccess() *string
}

type DescribeDownloadBackupSetStorageInfoResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeDownloadBackupSetStorageInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Argument: regionCode Must not be empty
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Argument: regionCode Must not be empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 44B8C2F5-919D-5D29-BCD5-DEB03467****
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDownloadBackupSetStorageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadBackupSetStorageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) GetData() *DescribeDownloadBackupSetStorageInfoResponseBodyData {
	return s.Data
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetCode(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetData(v *DescribeDownloadBackupSetStorageInfoResponseBodyData) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetErrCode(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetErrMessage(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetMessage(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetRequestId(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetSuccess(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDownloadBackupSetStorageInfoResponseBodyData struct {
	// The validity period of the URL.
	//
	// > This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1661329050
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The private download URL of the backup set.
	PrivateUrl *string `json:"PrivateUrl,omitempty" xml:"PrivateUrl,omitempty"`
	// The public download URL of the backup set.
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
}

func (s DescribeDownloadBackupSetStorageInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadBackupSetStorageInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) GetPrivateUrl() *string {
	return s.PrivateUrl
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) SetExpirationTime(v int64) *DescribeDownloadBackupSetStorageInfoResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) SetPrivateUrl(v string) *DescribeDownloadBackupSetStorageInfoResponseBodyData {
	s.PrivateUrl = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) SetPublicUrl(v string) *DescribeDownloadBackupSetStorageInfoResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
