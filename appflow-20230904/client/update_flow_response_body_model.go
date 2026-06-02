// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateFlowResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateFlowResponseBody
	GetRequestId() *string
}

type UpdateFlowResponseBody struct {
	// example:
	//
	// None
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A053FC9D-AB9D-5258-9355-8FA57EE888C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFlowResponseBody) SetData(v string) *UpdateFlowResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateFlowResponseBody) SetRequestId(v string) *UpdateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
