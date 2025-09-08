// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostAutomateResponseConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PostAutomateResponseConfigResponseBody
	GetCode() *int32
	SetData(v string) *PostAutomateResponseConfigResponseBody
	GetData() *string
	SetMessage(v string) *PostAutomateResponseConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *PostAutomateResponseConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostAutomateResponseConfigResponseBody
	GetSuccess() *bool
}

type PostAutomateResponseConfigResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostAutomateResponseConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostAutomateResponseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PostAutomateResponseConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *PostAutomateResponseConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PostAutomateResponseConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostAutomateResponseConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostAutomateResponseConfigResponseBody) SetCode(v int32) *PostAutomateResponseConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetData(v string) *PostAutomateResponseConfigResponseBody {
	s.Data = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetMessage(v string) *PostAutomateResponseConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetRequestId(v string) *PostAutomateResponseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetSuccess(v bool) *PostAutomateResponseConfigResponseBody {
	s.Success = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
