// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetIntegrationStateResponseBody
	GetCode() *int32
	SetMessage(v string) *GetIntegrationStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIntegrationStateResponseBody
	GetRequestId() *string
	SetState(v bool) *GetIntegrationStateResponseBody
	GetState() *bool
}

type GetIntegrationStateResponseBody struct {
	// Status code. 200 means success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The prompt information of the returned result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The integration state of Prometheus dashboards and collection rules. Valid values:
	//
	// 	- `true`: The Prometheus dashboards and collection rules that monitor the software are integrated.
	//
	// 	- `false`: The Prometheus dashboards and collection rules that monitor the software are not integrated.
	//
	// example:
	//
	// true
	State *bool `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetIntegrationStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegrationStateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetIntegrationStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIntegrationStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntegrationStateResponseBody) GetState() *bool {
	return s.State
}

func (s *GetIntegrationStateResponseBody) SetCode(v int32) *GetIntegrationStateResponseBody {
	s.Code = &v
	return s
}

func (s *GetIntegrationStateResponseBody) SetMessage(v string) *GetIntegrationStateResponseBody {
	s.Message = &v
	return s
}

func (s *GetIntegrationStateResponseBody) SetRequestId(v string) *GetIntegrationStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntegrationStateResponseBody) SetState(v bool) *GetIntegrationStateResponseBody {
	s.State = &v
	return s
}

func (s *GetIntegrationStateResponseBody) Validate() error {
	return dara.Validate(s)
}
