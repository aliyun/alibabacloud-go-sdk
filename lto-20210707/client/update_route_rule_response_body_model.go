// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRouteRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateRouteRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateRouteRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRouteRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRouteRuleResponseBody
	GetSuccess() *bool
}

type UpdateRouteRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRouteRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRouteRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRouteRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateRouteRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRouteRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRouteRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRouteRuleResponseBody) SetCode(v string) *UpdateRouteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRouteRuleResponseBody) SetHttpStatusCode(v int32) *UpdateRouteRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRouteRuleResponseBody) SetMessage(v string) *UpdateRouteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRouteRuleResponseBody) SetRequestId(v string) *UpdateRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRouteRuleResponseBody) SetSuccess(v bool) *UpdateRouteRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRouteRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
