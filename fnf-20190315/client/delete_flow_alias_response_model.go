// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowAliasResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlowAliasResponseBody) *DeleteFlowAliasResponse
	GetBody() *DeleteFlowAliasResponseBody
}

type DeleteFlowAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowAliasResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowAliasResponse) GetBody() *DeleteFlowAliasResponseBody {
	return s.Body
}

func (s *DeleteFlowAliasResponse) SetHeaders(v map[string]*string) *DeleteFlowAliasResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowAliasResponse) SetStatusCode(v int32) *DeleteFlowAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowAliasResponse) SetBody(v *DeleteFlowAliasResponseBody) *DeleteFlowAliasResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
