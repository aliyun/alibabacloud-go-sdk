// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRouteRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddRouteRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddRouteRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddRouteRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddRouteRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRouteRuleResponseBody
	GetSuccess() *bool
}

type AddRouteRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRouteRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddRouteRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddRouteRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddRouteRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddRouteRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRouteRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRouteRuleResponseBody) SetCode(v string) *AddRouteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddRouteRuleResponseBody) SetHttpStatusCode(v int32) *AddRouteRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddRouteRuleResponseBody) SetMessage(v string) *AddRouteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddRouteRuleResponseBody) SetRequestId(v string) *AddRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRouteRuleResponseBody) SetSuccess(v bool) *AddRouteRuleResponseBody {
	s.Success = &v
	return s
}

func (s *AddRouteRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
