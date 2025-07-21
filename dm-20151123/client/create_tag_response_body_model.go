// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTagResponseBody
	GetRequestId() *string
	SetTagId(v string) *CreateTagResponseBody
	GetTagId() *string
}

type CreateTagResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tag ID
	//
	// example:
	//
	// 91141
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagResponseBody) GetTagId() *string {
	return s.TagId
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) SetTagId(v string) *CreateTagResponseBody {
	s.TagId = &v
	return s
}

func (s *CreateTagResponseBody) Validate() error {
	return dara.Validate(s)
}
