// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterVpcEndpointConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateClusterVpcEndpointConnectionResponseBody
	GetCode() *string
	SetData(v *CreateClusterVpcEndpointConnectionResponseBodyData) *CreateClusterVpcEndpointConnectionResponseBody
	GetData() *CreateClusterVpcEndpointConnectionResponseBodyData
	SetMessage(v string) *CreateClusterVpcEndpointConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateClusterVpcEndpointConnectionResponseBody
	GetRequestId() *string
}

type CreateClusterVpcEndpointConnectionResponseBody struct {
	// example:
	//
	// Success
	Code *string                                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateClusterVpcEndpointConnectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateClusterVpcEndpointConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterVpcEndpointConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) GetData() *CreateClusterVpcEndpointConnectionResponseBodyData {
	return s.Data
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) SetCode(v string) *CreateClusterVpcEndpointConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) SetData(v *CreateClusterVpcEndpointConnectionResponseBodyData) *CreateClusterVpcEndpointConnectionResponseBody {
	s.Data = v
	return s
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) SetMessage(v string) *CreateClusterVpcEndpointConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) SetRequestId(v string) *CreateClusterVpcEndpointConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterVpcEndpointConnectionResponseBodyData struct {
	// example:
	//
	// ep-xxx
	EndpointConnectionId *string `json:"endpointConnectionId,omitempty" xml:"endpointConnectionId,omitempty"`
}

func (s CreateClusterVpcEndpointConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterVpcEndpointConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateClusterVpcEndpointConnectionResponseBodyData) GetEndpointConnectionId() *string {
	return s.EndpointConnectionId
}

func (s *CreateClusterVpcEndpointConnectionResponseBodyData) SetEndpointConnectionId(v string) *CreateClusterVpcEndpointConnectionResponseBodyData {
	s.EndpointConnectionId = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
