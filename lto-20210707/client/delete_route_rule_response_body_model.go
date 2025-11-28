// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRouteRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteRouteRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteRouteRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRouteRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRouteRuleResponseBody
	GetSuccess() *bool
}

type DeleteRouteRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRouteRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRouteRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRouteRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRouteRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRouteRuleResponseBody) SetCode(v string) *DeleteRouteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRouteRuleResponseBody) SetHttpStatusCode(v int32) *DeleteRouteRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRouteRuleResponseBody) SetMessage(v string) *DeleteRouteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRouteRuleResponseBody) SetRequestId(v string) *DeleteRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteRuleResponseBody) SetSuccess(v bool) *DeleteRouteRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRouteRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
