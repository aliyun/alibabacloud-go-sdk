// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2RevocationEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OAuth2RevocationEndpointResponseBody
	GetRequestId() *string
}

type OAuth2RevocationEndpointResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 4070039E-5822-1F32-9295-1D2883E48BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OAuth2RevocationEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OAuth2RevocationEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *OAuth2RevocationEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OAuth2RevocationEndpointResponseBody) SetRequestId(v string) *OAuth2RevocationEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *OAuth2RevocationEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
