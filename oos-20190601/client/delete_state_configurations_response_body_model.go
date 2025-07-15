// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStateConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStateConfigurationsResponseBody
	GetRequestId() *string
}

type DeleteStateConfigurationsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 12345B731-0FCE-48BA-8D42-605abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStateConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStateConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStateConfigurationsResponseBody) SetRequestId(v string) *DeleteStateConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStateConfigurationsResponseBody) Validate() error {
	return dara.Validate(s)
}
