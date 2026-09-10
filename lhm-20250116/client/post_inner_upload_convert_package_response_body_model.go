// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerUploadConvertPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *PostInnerUploadConvertPackageResponseBody
	GetData() *string
	SetErrCode(v string) *PostInnerUploadConvertPackageResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *PostInnerUploadConvertPackageResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *PostInnerUploadConvertPackageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostInnerUploadConvertPackageResponseBody
	GetSuccess() *bool
}

type PostInnerUploadConvertPackageResponseBody struct {
	// The business data returned by the operation as a string. The specific content varies by operation.
	//
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. This value is an empty string if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. This value is an empty string if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID that uniquely identifies this call. Provide this value when troubleshooting issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Troubleshoot by using errCode and errMessage.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PostInnerUploadConvertPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostInnerUploadConvertPackageResponseBody) GoString() string {
	return s.String()
}

func (s *PostInnerUploadConvertPackageResponseBody) GetData() *string {
	return s.Data
}

func (s *PostInnerUploadConvertPackageResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *PostInnerUploadConvertPackageResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *PostInnerUploadConvertPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostInnerUploadConvertPackageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostInnerUploadConvertPackageResponseBody) SetData(v string) *PostInnerUploadConvertPackageResponseBody {
	s.Data = &v
	return s
}

func (s *PostInnerUploadConvertPackageResponseBody) SetErrCode(v string) *PostInnerUploadConvertPackageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PostInnerUploadConvertPackageResponseBody) SetErrMessage(v string) *PostInnerUploadConvertPackageResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PostInnerUploadConvertPackageResponseBody) SetRequestId(v string) *PostInnerUploadConvertPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostInnerUploadConvertPackageResponseBody) SetSuccess(v bool) *PostInnerUploadConvertPackageResponseBody {
	s.Success = &v
	return s
}

func (s *PostInnerUploadConvertPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
