// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTransferCalleePoolConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetTransferCalleePoolConfigResponseBody
	GetCode() *string
	SetData(v bool) *SetTransferCalleePoolConfigResponseBody
	GetData() *bool
	SetMessage(v string) *SetTransferCalleePoolConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetTransferCalleePoolConfigResponseBody
	GetRequestId() *string
}

type SetTransferCalleePoolConfigResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the phone numbers for transferring the call were configured.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// E7518CE1-B0FF-4C6F-9252-BF80271B2F99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetTransferCalleePoolConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetTransferCalleePoolConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetTransferCalleePoolConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetTransferCalleePoolConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetTransferCalleePoolConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetTransferCalleePoolConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetTransferCalleePoolConfigResponseBody) SetCode(v string) *SetTransferCalleePoolConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponseBody) SetData(v bool) *SetTransferCalleePoolConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponseBody) SetMessage(v string) *SetTransferCalleePoolConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponseBody) SetRequestId(v string) *SetTransferCalleePoolConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
