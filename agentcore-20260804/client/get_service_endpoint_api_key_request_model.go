// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEndpointApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetServiceEndpointApiKeyRequest struct {
}

func (s GetServiceEndpointApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointApiKeyRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
