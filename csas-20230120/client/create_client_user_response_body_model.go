// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateClientUserResponseBody
	GetData() *string
	SetRequestId(v string) *CreateClientUserResponseBody
	GetRequestId() *string
}

type CreateClientUserResponseBody struct {
	// example:
	//
	// 726
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClientUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientUserResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateClientUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClientUserResponseBody) SetData(v string) *CreateClientUserResponseBody {
	s.Data = &v
	return s
}

func (s *CreateClientUserResponseBody) SetRequestId(v string) *CreateClientUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientUserResponseBody) Validate() error {
	return dara.Validate(s)
}
