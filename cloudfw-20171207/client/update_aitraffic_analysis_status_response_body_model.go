// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAITrafficAnalysisStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAITrafficAnalysisStatusResponseBody
	GetRequestId() *string
}

type UpdateAITrafficAnalysisStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 4E7F94C7-781F-5192-86CF-DB085013C810
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAITrafficAnalysisStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAITrafficAnalysisStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAITrafficAnalysisStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAITrafficAnalysisStatusResponseBody) SetRequestId(v string) *UpdateAITrafficAnalysisStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAITrafficAnalysisStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
