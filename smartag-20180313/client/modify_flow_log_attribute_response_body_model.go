// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowLogAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFlowLogAttributeResponseBody
	GetRequestId() *string
}

type ModifyFlowLogAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AAC70A63-1A2E-4857-9CA3-5DE5B4041D1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowLogAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowLogAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFlowLogAttributeResponseBody) SetRequestId(v string) *ModifyFlowLogAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
