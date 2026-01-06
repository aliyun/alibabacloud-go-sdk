// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceObservabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceObservabilityResponseBody
	GetRequestId() *string
}

type CreateServiceObservabilityResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceObservabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceObservabilityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceObservabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceObservabilityResponseBody) SetRequestId(v string) *CreateServiceObservabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceObservabilityResponseBody) Validate() error {
	return dara.Validate(s)
}
