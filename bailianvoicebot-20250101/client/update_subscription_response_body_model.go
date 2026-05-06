// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSubscriptionResponseBody
	GetCode() *string
	SetData(v string) *UpdateSubscriptionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSubscriptionResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateSubscriptionResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateSubscriptionResponseBody
	GetRequestId() *string
}

type UpdateSubscriptionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance llm-rj6aqmctjcit4acy does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSubscriptionResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSubscriptionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubscriptionResponseBody) SetCode(v string) *UpdateSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetData(v string) *UpdateSubscriptionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetHttpStatusCode(v int32) *UpdateSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetMessage(v string) *UpdateSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetParams(v []*string) *UpdateSubscriptionResponseBody {
	s.Params = v
	return s
}

func (s *UpdateSubscriptionResponseBody) SetRequestId(v string) *UpdateSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
