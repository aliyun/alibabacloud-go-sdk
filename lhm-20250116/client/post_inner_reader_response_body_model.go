// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerReaderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *PostInnerReaderResponseBody
	GetData() *string
	SetErrCode(v string) *PostInnerReaderResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *PostInnerReaderResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *PostInnerReaderResponseBody
	GetRequestId() *string
	SetSuccess(v string) *PostInnerReaderResponseBody
	GetSuccess() *string
}

type PostInnerReaderResponseBody struct {
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PostInnerReaderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostInnerReaderResponseBody) GoString() string {
	return s.String()
}

func (s *PostInnerReaderResponseBody) GetData() *string {
	return s.Data
}

func (s *PostInnerReaderResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *PostInnerReaderResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *PostInnerReaderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostInnerReaderResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PostInnerReaderResponseBody) SetData(v string) *PostInnerReaderResponseBody {
	s.Data = &v
	return s
}

func (s *PostInnerReaderResponseBody) SetErrCode(v string) *PostInnerReaderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PostInnerReaderResponseBody) SetErrMessage(v string) *PostInnerReaderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PostInnerReaderResponseBody) SetRequestId(v string) *PostInnerReaderResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostInnerReaderResponseBody) SetSuccess(v string) *PostInnerReaderResponseBody {
	s.Success = &v
	return s
}

func (s *PostInnerReaderResponseBody) Validate() error {
	return dara.Validate(s)
}
