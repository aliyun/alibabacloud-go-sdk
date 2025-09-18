// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicNetworkStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePublicNetworkStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePublicNetworkStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePublicNetworkStatusResponseBody) *UpdatePublicNetworkStatusResponse
	GetBody() *UpdatePublicNetworkStatusResponseBody
}

type UpdatePublicNetworkStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublicNetworkStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublicNetworkStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePublicNetworkStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePublicNetworkStatusResponse) GetBody() *UpdatePublicNetworkStatusResponseBody {
	return s.Body
}

func (s *UpdatePublicNetworkStatusResponse) SetHeaders(v map[string]*string) *UpdatePublicNetworkStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicNetworkStatusResponse) SetStatusCode(v int32) *UpdatePublicNetworkStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponse) SetBody(v *UpdatePublicNetworkStatusResponseBody) *UpdatePublicNetworkStatusResponse {
	s.Body = v
	return s
}

func (s *UpdatePublicNetworkStatusResponse) Validate() error {
	return dara.Validate(s)
}
