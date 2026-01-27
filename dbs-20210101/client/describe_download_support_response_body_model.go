// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadSupportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDownloadSupportResponseBody
	GetCode() *string
	SetData(v string) *DescribeDownloadSupportResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeDownloadSupportResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDownloadSupportResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeDownloadSupportResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDownloadSupportResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDownloadSupportResponseBody
	GetSuccess() *string
}

type DescribeDownloadSupportResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the advanced download feature is supported. Valid values:
	//
	// 	- **true**: The advanced download feature is supported.
	//
	// 	- **false**: The advanced download feature is not supported.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// F1A186F7-7B34-5C11-A903-EE23876B****
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

func (s DescribeDownloadSupportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadSupportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSupportResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDownloadSupportResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeDownloadSupportResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDownloadSupportResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDownloadSupportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDownloadSupportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadSupportResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDownloadSupportResponseBody) SetCode(v string) *DescribeDownloadSupportResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetData(v string) *DescribeDownloadSupportResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetErrCode(v string) *DescribeDownloadSupportResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetErrMessage(v string) *DescribeDownloadSupportResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetMessage(v string) *DescribeDownloadSupportResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetRequestId(v string) *DescribeDownloadSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetSuccess(v string) *DescribeDownloadSupportResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) Validate() error {
	return dara.Validate(s)
}
