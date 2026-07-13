// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceEndpointResponseBody
	GetCode() *string
	SetData(v *CreateServiceEndpointResponseBodyData) *CreateServiceEndpointResponseBody
	GetData() *CreateServiceEndpointResponseBodyData
	SetHttpStatusCode(v int32) *CreateServiceEndpointResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateServiceEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceEndpointResponseBody
	GetSuccess() *bool
}

type CreateServiceEndpointResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateServiceEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2b7f1c2d-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceEndpointResponseBody) GetData() *CreateServiceEndpointResponseBodyData {
	return s.Data
}

func (s *CreateServiceEndpointResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateServiceEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceEndpointResponseBody) SetCode(v string) *CreateServiceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetData(v *CreateServiceEndpointResponseBodyData) *CreateServiceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetHttpStatusCode(v int32) *CreateServiceEndpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetMessage(v string) *CreateServiceEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetRequestId(v string) *CreateServiceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetSuccess(v bool) *CreateServiceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceEndpointResponseBodyData struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyData) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateServiceEndpointResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceEndpointResponseBodyData) SetEndpointId(v string) *CreateServiceEndpointResponseBodyData {
	s.EndpointId = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetStatus(v string) *CreateServiceEndpointResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
