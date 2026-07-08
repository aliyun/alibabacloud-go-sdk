// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DisableFlowResponseBody
	GetData() *string
	SetRequestId(v string) *DisableFlowResponseBody
	GetRequestId() *string
}

type DisableFlowResponseBody struct {
	// The data returned by the operation.
	//
	// example:
	//
	// 无
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C0CEBCD9-CE5D-5BDE-B8E5-F3ADD1BB943F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DisableFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *DisableFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableFlowResponseBody) SetData(v string) *DisableFlowResponseBody {
	s.Data = &v
	return s
}

func (s *DisableFlowResponseBody) SetRequestId(v string) *DisableFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
