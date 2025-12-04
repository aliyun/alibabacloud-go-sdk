// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSyndbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSyndbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSyndbResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSyndbResponseBody) *DeleteSyndbResponse
	GetBody() *DeleteSyndbResponseBody
}

type DeleteSyndbResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSyndbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSyndbResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSyndbResponse) GoString() string {
	return s.String()
}

func (s *DeleteSyndbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSyndbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSyndbResponse) GetBody() *DeleteSyndbResponseBody {
	return s.Body
}

func (s *DeleteSyndbResponse) SetHeaders(v map[string]*string) *DeleteSyndbResponse {
	s.Headers = v
	return s
}

func (s *DeleteSyndbResponse) SetStatusCode(v int32) *DeleteSyndbResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSyndbResponse) SetBody(v *DeleteSyndbResponseBody) *DeleteSyndbResponse {
	s.Body = v
	return s
}

func (s *DeleteSyndbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
