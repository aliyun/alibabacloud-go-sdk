// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTableMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTableMetaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTableMetaResponseBody) *DeleteTableMetaResponse
	GetBody() *DeleteTableMetaResponseBody
}

type DeleteTableMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTableMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTableMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTableMetaResponse) GetBody() *DeleteTableMetaResponseBody {
	return s.Body
}

func (s *DeleteTableMetaResponse) SetHeaders(v map[string]*string) *DeleteTableMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableMetaResponse) SetStatusCode(v int32) *DeleteTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableMetaResponse) SetBody(v *DeleteTableMetaResponseBody) *DeleteTableMetaResponse {
	s.Body = v
	return s
}

func (s *DeleteTableMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
