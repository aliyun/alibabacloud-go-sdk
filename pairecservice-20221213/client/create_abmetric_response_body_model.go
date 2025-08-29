// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetABMetricId(v string) *CreateABMetricResponseBody
	GetABMetricId() *string
	SetRequestId(v string) *CreateABMetricResponseBody
	GetRequestId() *string
}

type CreateABMetricResponseBody struct {
	// example:
	//
	// 1
	ABMetricId *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateABMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateABMetricResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABMetricResponseBody) GetABMetricId() *string {
	return s.ABMetricId
}

func (s *CreateABMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateABMetricResponseBody) SetABMetricId(v string) *CreateABMetricResponseBody {
	s.ABMetricId = &v
	return s
}

func (s *CreateABMetricResponseBody) SetRequestId(v string) *CreateABMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
