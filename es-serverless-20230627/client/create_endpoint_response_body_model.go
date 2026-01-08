// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEndpointResponseBody
	GetRequestId() *string
	SetResult(v *CreateEndpointResponseBodyResult) *CreateEndpointResponseBody
	GetResult() *CreateEndpointResponseBodyResult
}

type CreateEndpointResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateEndpointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEndpointResponseBody) GetResult() *CreateEndpointResponseBodyResult {
	return s.Result
}

func (s *CreateEndpointResponseBody) SetRequestId(v string) *CreateEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEndpointResponseBody) SetResult(v *CreateEndpointResponseBodyResult) *CreateEndpointResponseBody {
	s.Result = v
	return s
}

func (s *CreateEndpointResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEndpointResponseBodyResult struct {
	// example:
	//
	// essep-abd***dks
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
}

func (s CreateEndpointResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponseBodyResult) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateEndpointResponseBodyResult) SetEndpointId(v string) *CreateEndpointResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *CreateEndpointResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
