// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardWordRootResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStandardWordRootResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStandardWordRootResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStandardWordRootResponseBody) *DeleteStandardWordRootResponse
	GetBody() *DeleteStandardWordRootResponseBody
}

type DeleteStandardWordRootResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStandardWordRootResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStandardWordRootResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardWordRootResponse) GoString() string {
	return s.String()
}

func (s *DeleteStandardWordRootResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStandardWordRootResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStandardWordRootResponse) GetBody() *DeleteStandardWordRootResponseBody {
	return s.Body
}

func (s *DeleteStandardWordRootResponse) SetHeaders(v map[string]*string) *DeleteStandardWordRootResponse {
	s.Headers = v
	return s
}

func (s *DeleteStandardWordRootResponse) SetStatusCode(v int32) *DeleteStandardWordRootResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStandardWordRootResponse) SetBody(v *DeleteStandardWordRootResponseBody) *DeleteStandardWordRootResponse {
	s.Body = v
	return s
}

func (s *DeleteStandardWordRootResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
