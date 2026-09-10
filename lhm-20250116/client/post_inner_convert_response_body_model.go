// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerConvertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *PostInnerConvertResponseBody
	GetData() *string
	SetErrCode(v string) *PostInnerConvertResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *PostInnerConvertResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *PostInnerConvertResponseBody
	GetRequestId() *string
	SetSuccess(v string) *PostInnerConvertResponseBody
	GetSuccess() *string
}

type PostInnerConvertResponseBody struct {
	// Business data, the identifier of the transformation task returned upon successful submission of this interface.
	//
	// example:
	//
	// c8f3a1b2e9d74c5f
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Error code, which is an empty string when the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// Error message, which is an empty string when the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Request ID, used for locating and troubleshooting issues with this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful: `true` indicates success, `false` indicates failure. In case of failure, troubleshoot using `errCode` and `errMessage`.
	//
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PostInnerConvertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostInnerConvertResponseBody) GoString() string {
	return s.String()
}

func (s *PostInnerConvertResponseBody) GetData() *string {
	return s.Data
}

func (s *PostInnerConvertResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *PostInnerConvertResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *PostInnerConvertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostInnerConvertResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PostInnerConvertResponseBody) SetData(v string) *PostInnerConvertResponseBody {
	s.Data = &v
	return s
}

func (s *PostInnerConvertResponseBody) SetErrCode(v string) *PostInnerConvertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PostInnerConvertResponseBody) SetErrMessage(v string) *PostInnerConvertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PostInnerConvertResponseBody) SetRequestId(v string) *PostInnerConvertResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostInnerConvertResponseBody) SetSuccess(v string) *PostInnerConvertResponseBody {
	s.Success = &v
	return s
}

func (s *PostInnerConvertResponseBody) Validate() error {
	return dara.Validate(s)
}
