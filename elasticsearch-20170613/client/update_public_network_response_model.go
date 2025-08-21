// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePublicNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePublicNetworkResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePublicNetworkResponseBody) *UpdatePublicNetworkResponse
	GetBody() *UpdatePublicNetworkResponseBody
}

type UpdatePublicNetworkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublicNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublicNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePublicNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePublicNetworkResponse) GetBody() *UpdatePublicNetworkResponseBody {
	return s.Body
}

func (s *UpdatePublicNetworkResponse) SetHeaders(v map[string]*string) *UpdatePublicNetworkResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicNetworkResponse) SetStatusCode(v int32) *UpdatePublicNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublicNetworkResponse) SetBody(v *UpdatePublicNetworkResponseBody) *UpdatePublicNetworkResponse {
	s.Body = v
	return s
}

func (s *UpdatePublicNetworkResponse) Validate() error {
	return dara.Validate(s)
}
