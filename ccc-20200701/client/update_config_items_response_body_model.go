// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConfigItemsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateConfigItemsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateConfigItemsResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateConfigItemsResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateConfigItemsResponseBody
	GetRequestId() *string
}

type UpdateConfigItemsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 8E7Y5B39-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigItemsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConfigItemsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateConfigItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConfigItemsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateConfigItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigItemsResponseBody) SetCode(v string) *UpdateConfigItemsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConfigItemsResponseBody) SetHttpStatusCode(v int32) *UpdateConfigItemsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConfigItemsResponseBody) SetMessage(v string) *UpdateConfigItemsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConfigItemsResponseBody) SetParams(v []*string) *UpdateConfigItemsResponseBody {
	s.Params = v
	return s
}

func (s *UpdateConfigItemsResponseBody) SetRequestId(v string) *UpdateConfigItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
