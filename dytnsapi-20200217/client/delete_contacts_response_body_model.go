// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteContactsResponseBody
	GetCode() *string
	SetData(v bool) *DeleteContactsResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContactsResponseBody
	GetRequestId() *string
}

type DeleteContactsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContactsResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactsResponseBody) SetCode(v string) *DeleteContactsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactsResponseBody) SetData(v bool) *DeleteContactsResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteContactsResponseBody) SetMessage(v string) *DeleteContactsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactsResponseBody) SetRequestId(v string) *DeleteContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactsResponseBody) Validate() error {
	return dara.Validate(s)
}
