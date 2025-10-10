// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRulesAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateRulesAttributeResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateRulesAttributeResponseBody
	GetRequestId() *string
}

type UpdateRulesAttributeResponseBody struct {
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
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRulesAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateRulesAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRulesAttributeResponseBody) SetJobId(v string) *UpdateRulesAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateRulesAttributeResponseBody) SetRequestId(v string) *UpdateRulesAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRulesAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
