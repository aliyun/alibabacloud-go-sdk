// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePerformanceDataCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePerformanceDataCollectionResponseBody
	GetRequestId() *string
}

type UpdatePerformanceDataCollectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePerformanceDataCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePerformanceDataCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePerformanceDataCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePerformanceDataCollectionResponseBody) SetRequestId(v string) *UpdatePerformanceDataCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePerformanceDataCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
