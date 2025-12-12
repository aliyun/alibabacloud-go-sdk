// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCheckTypeToSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCheckTypeToSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCheckTypeToSchemeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCheckTypeToSchemeResponseBody) *DeleteCheckTypeToSchemeResponse
	GetBody() *DeleteCheckTypeToSchemeResponseBody
}

type DeleteCheckTypeToSchemeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCheckTypeToSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCheckTypeToSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckTypeToSchemeResponse) GoString() string {
	return s.String()
}

func (s *DeleteCheckTypeToSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCheckTypeToSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCheckTypeToSchemeResponse) GetBody() *DeleteCheckTypeToSchemeResponseBody {
	return s.Body
}

func (s *DeleteCheckTypeToSchemeResponse) SetHeaders(v map[string]*string) *DeleteCheckTypeToSchemeResponse {
	s.Headers = v
	return s
}

func (s *DeleteCheckTypeToSchemeResponse) SetStatusCode(v int32) *DeleteCheckTypeToSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCheckTypeToSchemeResponse) SetBody(v *DeleteCheckTypeToSchemeResponseBody) *DeleteCheckTypeToSchemeResponse {
	s.Body = v
	return s
}

func (s *DeleteCheckTypeToSchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
