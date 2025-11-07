// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkPathResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkPathResponseBody) *CreateNetworkPathResponse
	GetBody() *CreateNetworkPathResponseBody
}

type CreateNetworkPathResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkPathResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPathResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkPathResponse) GetBody() *CreateNetworkPathResponseBody {
	return s.Body
}

func (s *CreateNetworkPathResponse) SetHeaders(v map[string]*string) *CreateNetworkPathResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkPathResponse) SetStatusCode(v int32) *CreateNetworkPathResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkPathResponse) SetBody(v *CreateNetworkPathResponseBody) *CreateNetworkPathResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
