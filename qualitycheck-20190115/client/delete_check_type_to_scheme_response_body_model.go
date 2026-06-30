// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCheckTypeToSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCheckTypeToSchemeResponseBody
	GetCode() *string
	SetData(v string) *DeleteCheckTypeToSchemeResponseBody
	GetData() *string
	SetMessage(v string) *DeleteCheckTypeToSchemeResponseBody
	GetMessage() *string
	SetMessages(v []*string) *DeleteCheckTypeToSchemeResponseBody
	GetMessages() []*string
	SetRequestId(v string) *DeleteCheckTypeToSchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCheckTypeToSchemeResponseBody
	GetSuccess() *bool
}

type DeleteCheckTypeToSchemeResponseBody struct {
	// The result code. A value of **200*	- indicates that the request was successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 48864
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned. If the request is successful, \\`successful\\` is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A list of returned messages.
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates success. A value of false or null indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCheckTypeToSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckTypeToSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCheckTypeToSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCheckTypeToSchemeResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCheckTypeToSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCheckTypeToSchemeResponseBody) GetMessages() []*string {
	return s.Messages
}

func (s *DeleteCheckTypeToSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCheckTypeToSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCheckTypeToSchemeResponseBody) SetCode(v string) *DeleteCheckTypeToSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCheckTypeToSchemeResponseBody) SetData(v string) *DeleteCheckTypeToSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCheckTypeToSchemeResponseBody) SetMessage(v string) *DeleteCheckTypeToSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCheckTypeToSchemeResponseBody) SetMessages(v []*string) *DeleteCheckTypeToSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteCheckTypeToSchemeResponseBody) SetRequestId(v string) *DeleteCheckTypeToSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCheckTypeToSchemeResponseBody) SetSuccess(v bool) *DeleteCheckTypeToSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCheckTypeToSchemeResponseBody) Validate() error {
	return dara.Validate(s)
}
