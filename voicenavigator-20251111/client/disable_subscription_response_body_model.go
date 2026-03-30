// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableSubscriptionResponseBody
	GetCode() *string
	SetData(v string) *DisableSubscriptionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DisableSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableSubscriptionResponseBody
	GetMessage() *string
	SetParams(v []*string) *DisableSubscriptionResponseBody
	GetParams() []*string
	SetRequestId(v string) *DisableSubscriptionResponseBody
	GetRequestId() *string
}

type DisableSubscriptionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableSubscriptionResponseBody) GetData() *string {
	return s.Data
}

func (s *DisableSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableSubscriptionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DisableSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSubscriptionResponseBody) SetCode(v string) *DisableSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetData(v string) *DisableSubscriptionResponseBody {
	s.Data = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetHttpStatusCode(v int32) *DisableSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetMessage(v string) *DisableSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *DisableSubscriptionResponseBody) SetParams(v []*string) *DisableSubscriptionResponseBody {
	s.Params = v
	return s
}

func (s *DisableSubscriptionResponseBody) SetRequestId(v string) *DisableSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
