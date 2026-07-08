// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateFlowResponseBody
	GetData() *string
	SetRequestId(v string) *CreateFlowResponseBody
	GetRequestId() *string
}

type CreateFlowResponseBody struct {
	// The information of the newly created product returned when the call succeeds.
	//
	// example:
	//
	// flow-9691a04fc7f94525aac1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32F02021-11D2-5196-BC85-72890E9AA090
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowResponseBody) SetData(v string) *CreateFlowResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
