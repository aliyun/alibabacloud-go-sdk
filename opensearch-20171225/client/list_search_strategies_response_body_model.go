// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSearchStrategiesResponseBody
	GetRequestId() *string
}

type ListSearchStrategiesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9C6351F5-2E2E-5249-888B-88A74E1B8A65
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSearchStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchStrategiesResponseBody) SetRequestId(v string) *ListSearchStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchStrategiesResponseBody) Validate() error {
	return dara.Validate(s)
}
