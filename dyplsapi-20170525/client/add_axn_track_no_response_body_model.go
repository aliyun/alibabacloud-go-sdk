// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAxnTrackNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddAxnTrackNoResponseBody
	GetCode() *string
	SetMessage(v string) *AddAxnTrackNoResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAxnTrackNoResponseBody
	GetRequestId() *string
}

type AddAxnTrackNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
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
	// 1A4CADEF-8516-5890-9177-A1A4D97F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAxnTrackNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAxnTrackNoResponseBody) GoString() string {
	return s.String()
}

func (s *AddAxnTrackNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddAxnTrackNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAxnTrackNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAxnTrackNoResponseBody) SetCode(v string) *AddAxnTrackNoResponseBody {
	s.Code = &v
	return s
}

func (s *AddAxnTrackNoResponseBody) SetMessage(v string) *AddAxnTrackNoResponseBody {
	s.Message = &v
	return s
}

func (s *AddAxnTrackNoResponseBody) SetRequestId(v string) *AddAxnTrackNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAxnTrackNoResponseBody) Validate() error {
	return dara.Validate(s)
}
