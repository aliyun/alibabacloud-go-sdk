// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageBasicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ReadPageItem) *ReadPageBasicResponseBody
	GetData() *ReadPageItem
	SetErrorCode(v string) *ReadPageBasicResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ReadPageBasicResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ReadPageBasicResponseBody
	GetRequestId() *string
}

type ReadPageBasicResponseBody struct {
	Data         *ReadPageItem `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ReadPageBasicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadPageBasicResponseBody) GoString() string {
	return s.String()
}

func (s *ReadPageBasicResponseBody) GetData() *ReadPageItem {
	return s.Data
}

func (s *ReadPageBasicResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReadPageBasicResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReadPageBasicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadPageBasicResponseBody) SetData(v *ReadPageItem) *ReadPageBasicResponseBody {
	s.Data = v
	return s
}

func (s *ReadPageBasicResponseBody) SetErrorCode(v string) *ReadPageBasicResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReadPageBasicResponseBody) SetErrorMessage(v string) *ReadPageBasicResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReadPageBasicResponseBody) SetRequestId(v string) *ReadPageBasicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadPageBasicResponseBody) Validate() error {
	return dara.Validate(s)
}
