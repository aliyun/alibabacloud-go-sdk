// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertGraphDraftResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevertGraphDraftResourceResponseBody
	GetCode() *string
	SetMessage(v string) *RevertGraphDraftResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevertGraphDraftResourceResponseBody
	GetRequestId() *string
	SetReverted(v bool) *RevertGraphDraftResourceResponseBody
	GetReverted() *bool
}

type RevertGraphDraftResourceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the draft is actually revoked (true / false).
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Reverted *bool `json:"reverted,omitempty" xml:"reverted,omitempty"`
}

func (s RevertGraphDraftResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevertGraphDraftResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RevertGraphDraftResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevertGraphDraftResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevertGraphDraftResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevertGraphDraftResourceResponseBody) GetReverted() *bool {
	return s.Reverted
}

func (s *RevertGraphDraftResourceResponseBody) SetCode(v string) *RevertGraphDraftResourceResponseBody {
	s.Code = &v
	return s
}

func (s *RevertGraphDraftResourceResponseBody) SetMessage(v string) *RevertGraphDraftResourceResponseBody {
	s.Message = &v
	return s
}

func (s *RevertGraphDraftResourceResponseBody) SetRequestId(v string) *RevertGraphDraftResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertGraphDraftResourceResponseBody) SetReverted(v bool) *RevertGraphDraftResourceResponseBody {
	s.Reverted = &v
	return s
}

func (s *RevertGraphDraftResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
