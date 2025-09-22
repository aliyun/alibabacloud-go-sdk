// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRunningPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetRunningPlanResponseBody
	GetData() *bool
	SetRequestId(v string) *SetRunningPlanResponseBody
	GetRequestId() *string
}

type SetRunningPlanResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SetRunningPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRunningPlanResponseBody) GoString() string {
	return s.String()
}

func (s *SetRunningPlanResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetRunningPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRunningPlanResponseBody) SetData(v bool) *SetRunningPlanResponseBody {
	s.Data = &v
	return s
}

func (s *SetRunningPlanResponseBody) SetRequestId(v string) *SetRunningPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRunningPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
