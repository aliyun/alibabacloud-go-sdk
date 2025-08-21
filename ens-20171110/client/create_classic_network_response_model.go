// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClassicNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClassicNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClassicNetworkResponse
	GetStatusCode() *int32
	SetBody(v *CreateClassicNetworkResponseBody) *CreateClassicNetworkResponse
	GetBody() *CreateClassicNetworkResponseBody
}

type CreateClassicNetworkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClassicNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClassicNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClassicNetworkResponse) GoString() string {
	return s.String()
}

func (s *CreateClassicNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClassicNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClassicNetworkResponse) GetBody() *CreateClassicNetworkResponseBody {
	return s.Body
}

func (s *CreateClassicNetworkResponse) SetHeaders(v map[string]*string) *CreateClassicNetworkResponse {
	s.Headers = v
	return s
}

func (s *CreateClassicNetworkResponse) SetStatusCode(v int32) *CreateClassicNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClassicNetworkResponse) SetBody(v *CreateClassicNetworkResponseBody) *CreateClassicNetworkResponse {
	s.Body = v
	return s
}

func (s *CreateClassicNetworkResponse) Validate() error {
	return dara.Validate(s)
}
