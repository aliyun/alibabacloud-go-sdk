// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySubsIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySubsIdResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySubsIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySubsIdResponseBody
	GetRequestId() *string
	SetSubsId(v string) *QuerySubsIdResponseBody
	GetSubsId() *string
}

type QuerySubsIdResponseBody struct {
	// The response code. The value OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7F99446-8191-43C0-99B5-F58A6AEAD779
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The binding ID.
	//
	// example:
	//
	// 11111111****
	SubsId *string `json:"SubsId,omitempty" xml:"SubsId,omitempty"`
}

func (s QuerySubsIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySubsIdResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubsIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySubsIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySubsIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySubsIdResponseBody) GetSubsId() *string {
	return s.SubsId
}

func (s *QuerySubsIdResponseBody) SetCode(v string) *QuerySubsIdResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetMessage(v string) *QuerySubsIdResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetRequestId(v string) *QuerySubsIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySubsIdResponseBody) SetSubsId(v string) *QuerySubsIdResponseBody {
	s.SubsId = &v
	return s
}

func (s *QuerySubsIdResponseBody) Validate() error {
	return dara.Validate(s)
}
