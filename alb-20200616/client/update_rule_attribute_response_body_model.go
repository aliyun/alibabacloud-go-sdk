// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateRuleAttributeResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateRuleAttributeResponseBody
	GetRequestId() *string
}

type UpdateRuleAttributeResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F5378-41F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRuleAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateRuleAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRuleAttributeResponseBody) SetJobId(v string) *UpdateRuleAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateRuleAttributeResponseBody) SetRequestId(v string) *UpdateRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
