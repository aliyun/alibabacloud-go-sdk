// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddSecurityGroupRuleResponseBody
	GetCode() *int32
	SetData(v string) *AddSecurityGroupRuleResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *AddSecurityGroupRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddSecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSecurityGroupRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddSecurityGroupRuleResponseBody
	GetSuccess() *bool
}

type AddSecurityGroupRuleResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 12
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D0DB055C-51F2-5BB2-82A6-CD8A3C6EE6BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddSecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddSecurityGroupRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *AddSecurityGroupRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddSecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSecurityGroupRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddSecurityGroupRuleResponseBody) SetCode(v int32) *AddSecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddSecurityGroupRuleResponseBody) SetData(v string) *AddSecurityGroupRuleResponseBody {
	s.Data = &v
	return s
}

func (s *AddSecurityGroupRuleResponseBody) SetHttpStatusCode(v int32) *AddSecurityGroupRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddSecurityGroupRuleResponseBody) SetMessage(v string) *AddSecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddSecurityGroupRuleResponseBody) SetRequestId(v string) *AddSecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSecurityGroupRuleResponseBody) SetSuccess(v bool) *AddSecurityGroupRuleResponseBody {
	s.Success = &v
	return s
}

func (s *AddSecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
