// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReverseWriterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *StartReverseWriterResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartReverseWriterResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *StartReverseWriterResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StartReverseWriterResponseBody
	GetSuccess() *string
}

type StartReverseWriterResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 291BA4F1-2035-3FAA-6D5A-5D2015CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartReverseWriterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartReverseWriterResponseBody) GoString() string {
	return s.String()
}

func (s *StartReverseWriterResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartReverseWriterResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartReverseWriterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartReverseWriterResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StartReverseWriterResponseBody) SetErrCode(v string) *StartReverseWriterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartReverseWriterResponseBody) SetErrMessage(v string) *StartReverseWriterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartReverseWriterResponseBody) SetRequestId(v string) *StartReverseWriterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartReverseWriterResponseBody) SetSuccess(v string) *StartReverseWriterResponseBody {
	s.Success = &v
	return s
}

func (s *StartReverseWriterResponseBody) Validate() error {
	return dara.Validate(s)
}
