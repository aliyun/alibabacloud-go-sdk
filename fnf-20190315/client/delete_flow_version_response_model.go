// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlowVersionResponseBody) *DeleteFlowVersionResponse
	GetBody() *DeleteFlowVersionResponseBody
}

type DeleteFlowVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowVersionResponse) GetBody() *DeleteFlowVersionResponseBody {
	return s.Body
}

func (s *DeleteFlowVersionResponse) SetHeaders(v map[string]*string) *DeleteFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowVersionResponse) SetStatusCode(v int32) *DeleteFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowVersionResponse) SetBody(v *DeleteFlowVersionResponseBody) *DeleteFlowVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
