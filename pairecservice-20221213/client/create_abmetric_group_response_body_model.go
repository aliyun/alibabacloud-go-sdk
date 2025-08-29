// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABMetricGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetABMetricGroupId(v string) *CreateABMetricGroupResponseBody
	GetABMetricGroupId() *string
	SetRequestId(v string) *CreateABMetricGroupResponseBody
	GetRequestId() *string
}

type CreateABMetricGroupResponseBody struct {
	// example:
	//
	// 1
	ABMetricGroupId *string `json:"ABMetricGroupId,omitempty" xml:"ABMetricGroupId,omitempty"`
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateABMetricGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABMetricGroupResponseBody) GetABMetricGroupId() *string {
	return s.ABMetricGroupId
}

func (s *CreateABMetricGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateABMetricGroupResponseBody) SetABMetricGroupId(v string) *CreateABMetricGroupResponseBody {
	s.ABMetricGroupId = &v
	return s
}

func (s *CreateABMetricGroupResponseBody) SetRequestId(v string) *CreateABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABMetricGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
