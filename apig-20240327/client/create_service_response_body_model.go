// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceResponseBody
	GetCode() *string
	SetData(v *CreateServiceResponseBodyData) *CreateServiceResponseBody
	GetData() *CreateServiceResponseBodyData
	SetMessage(v string) *CreateServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceResponseBody
	GetRequestId() *string
}

type CreateServiceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *CreateServiceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C67DED2B-F19B-5BEC-88C1-D6EB854CD0D4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceResponseBody) GetData() *CreateServiceResponseBodyData {
	return s.Data
}

func (s *CreateServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceResponseBody) SetCode(v string) *CreateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceResponseBody) SetData(v *CreateServiceResponseBodyData) *CreateServiceResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceResponseBody) SetMessage(v string) *CreateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceResponseBodyData struct {
	// The list of service IDs.
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
}

func (s CreateServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyData) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *CreateServiceResponseBodyData) SetServiceIds(v []*string) *CreateServiceResponseBodyData {
	s.ServiceIds = v
	return s
}

func (s *CreateServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
