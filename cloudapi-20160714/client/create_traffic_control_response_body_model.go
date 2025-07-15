// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrafficControlResponseBody
	GetRequestId() *string
	SetTrafficControlId(v string) *CreateTrafficControlResponseBody
	GetTrafficControlId() *string
}

type CreateTrafficControlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the throttling policy.
	//
	// example:
	//
	// tf123456
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s CreateTrafficControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrafficControlResponseBody) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *CreateTrafficControlResponseBody) SetRequestId(v string) *CreateTrafficControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficControlResponseBody) SetTrafficControlId(v string) *CreateTrafficControlResponseBody {
	s.TrafficControlId = &v
	return s
}

func (s *CreateTrafficControlResponseBody) Validate() error {
	return dara.Validate(s)
}
