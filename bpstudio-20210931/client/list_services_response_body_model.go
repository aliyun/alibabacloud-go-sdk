// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListServicesResponseBody
	GetCode() *string
	SetData(v []*string) *ListServicesResponseBody
	GetData() []*string
	SetMessage(v string) *ListServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListServicesResponseBody
	GetRequestId() *string
}

type ListServicesResponseBody struct {
	// example:
	//
	// 200
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListServicesResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServicesResponseBody) SetCode(v string) *ListServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListServicesResponseBody) SetData(v []*string) *ListServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListServicesResponseBody) SetMessage(v string) *ListServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) Validate() error {
	return dara.Validate(s)
}
