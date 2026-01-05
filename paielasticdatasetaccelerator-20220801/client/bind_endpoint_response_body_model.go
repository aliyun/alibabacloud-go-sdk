// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindEndpointResponseBody
	GetRequestId() *string
}

type BindEndpointResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *BindEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindEndpointResponseBody) SetRequestId(v string) *BindEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
