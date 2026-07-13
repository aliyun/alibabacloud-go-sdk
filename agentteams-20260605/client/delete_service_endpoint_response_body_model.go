// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteServiceEndpointResponseBody
	GetCode() *string
	SetData(v bool) *DeleteServiceEndpointResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteServiceEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteServiceEndpointResponseBody
	GetSuccess() *bool
}

type DeleteServiceEndpointResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteServiceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteServiceEndpointResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteServiceEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteServiceEndpointResponseBody) SetCode(v string) *DeleteServiceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceEndpointResponseBody) SetData(v bool) *DeleteServiceEndpointResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteServiceEndpointResponseBody) SetMessage(v string) *DeleteServiceEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceEndpointResponseBody) SetRequestId(v string) *DeleteServiceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceEndpointResponseBody) SetSuccess(v bool) *DeleteServiceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteServiceEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
