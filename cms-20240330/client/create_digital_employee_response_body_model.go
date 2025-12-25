// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalEmployeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateDigitalEmployeeResponseBody
	GetName() *string
	SetRequestId(v string) *CreateDigitalEmployeeResponseBody
	GetRequestId() *string
}

type CreateDigitalEmployeeResponseBody struct {
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 3B311FD9-A60B-55E0-A896-A0C73*********
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDigitalEmployeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateDigitalEmployeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDigitalEmployeeResponseBody) SetName(v string) *CreateDigitalEmployeeResponseBody {
	s.Name = &v
	return s
}

func (s *CreateDigitalEmployeeResponseBody) SetRequestId(v string) *CreateDigitalEmployeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDigitalEmployeeResponseBody) Validate() error {
	return dara.Validate(s)
}
