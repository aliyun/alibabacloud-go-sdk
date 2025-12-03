// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerlessInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServerlessInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServerlessInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServerlessInstanceResponseBody) *DeleteServerlessInstanceResponse
	GetBody() *DeleteServerlessInstanceResponseBody
}

type DeleteServerlessInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServerlessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServerlessInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerlessInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerlessInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServerlessInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServerlessInstanceResponse) GetBody() *DeleteServerlessInstanceResponseBody {
	return s.Body
}

func (s *DeleteServerlessInstanceResponse) SetHeaders(v map[string]*string) *DeleteServerlessInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerlessInstanceResponse) SetStatusCode(v int32) *DeleteServerlessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerlessInstanceResponse) SetBody(v *DeleteServerlessInstanceResponseBody) *DeleteServerlessInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteServerlessInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
