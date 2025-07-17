// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkResponseBody) *CreateNetworkResponse
	GetBody() *CreateNetworkResponseBody
}

type CreateNetworkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkResponse) GetBody() *CreateNetworkResponseBody {
	return s.Body
}

func (s *CreateNetworkResponse) SetHeaders(v map[string]*string) *CreateNetworkResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkResponse) SetStatusCode(v int32) *CreateNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkResponse) SetBody(v *CreateNetworkResponseBody) *CreateNetworkResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkResponse) Validate() error {
	return dara.Validate(s)
}
