// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSavepointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSavepointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSavepointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSavepointResponseBody) *DeleteSavepointResponse
	GetBody() *DeleteSavepointResponseBody
}

type DeleteSavepointResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSavepointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSavepointResponse) GoString() string {
	return s.String()
}

func (s *DeleteSavepointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSavepointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSavepointResponse) GetBody() *DeleteSavepointResponseBody {
	return s.Body
}

func (s *DeleteSavepointResponse) SetHeaders(v map[string]*string) *DeleteSavepointResponse {
	s.Headers = v
	return s
}

func (s *DeleteSavepointResponse) SetStatusCode(v int32) *DeleteSavepointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSavepointResponse) SetBody(v *DeleteSavepointResponseBody) *DeleteSavepointResponse {
	s.Body = v
	return s
}

func (s *DeleteSavepointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
