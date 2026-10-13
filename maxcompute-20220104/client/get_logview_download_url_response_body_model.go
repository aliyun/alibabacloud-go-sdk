// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogviewDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetLogviewDownloadUrlResponseBody
	GetData() *string
	SetErrorCode(v string) *GetLogviewDownloadUrlResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetLogviewDownloadUrlResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetLogviewDownloadUrlResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetLogviewDownloadUrlResponseBody
	GetRequestId() *string
}

type GetLogviewDownloadUrlResponseBody struct {
	// The presigned OSS download URL for the offline Logview file. The URL is valid for 10 minutes from the time it is generated. The file name extension is .txt and the content is JSON data in UTF-8 encoding. The download URL is determined by useVpcEndpoint. If the link expires, invoke this operation again. The sample URL is desensitized and cannot be used for actual downloads.
	//
	// example:
	//
	// https://example-bucket.oss-cn-hangzhou.aliyuncs.com/logview/offline/logview_detail_odps_project_20260409110419512gx6ofeaj8fb.txt?OSSAccessKeyId=LTAI5tExample****&Expires=1791777900&Signature=example****
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. This field is not assigned a value when this operation is successfully processed. The sample value NO_PERMISSION is used to illustrate a common permission error and must not be combined with the success response example.
	//
	// example:
	//
	// NO_PERMISSION
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This field is not assigned a value when this operation is successfully processed. The sample value Forbidden is used to illustrate a common permission error and must not be combined with the success response example.
	//
	// example:
	//
	// Forbidden
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code field. This field is not assigned a value when this operation is successfully processed. Do not rely solely on whether this field equals 200 to determine success. Process the result based on the actual HTTP response status and returned content. The sample value 200 is used only to illustrate the field format.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID. This ID is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLogviewDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogviewDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogviewDownloadUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetLogviewDownloadUrlResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetLogviewDownloadUrlResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetLogviewDownloadUrlResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetLogviewDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogviewDownloadUrlResponseBody) SetData(v string) *GetLogviewDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetLogviewDownloadUrlResponseBody) SetErrorCode(v string) *GetLogviewDownloadUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLogviewDownloadUrlResponseBody) SetErrorMsg(v string) *GetLogviewDownloadUrlResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetLogviewDownloadUrlResponseBody) SetHttpCode(v int32) *GetLogviewDownloadUrlResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetLogviewDownloadUrlResponseBody) SetRequestId(v string) *GetLogviewDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogviewDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
