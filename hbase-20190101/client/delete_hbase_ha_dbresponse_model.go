// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHBaseHaDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHBaseHaDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHBaseHaDBResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHBaseHaDBResponseBody) *DeleteHBaseHaDBResponse
	GetBody() *DeleteHBaseHaDBResponseBody
}

type DeleteHBaseHaDBResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHBaseHaDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHBaseHaDBResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHBaseHaDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHBaseHaDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHBaseHaDBResponse) GetBody() *DeleteHBaseHaDBResponseBody {
	return s.Body
}

func (s *DeleteHBaseHaDBResponse) SetHeaders(v map[string]*string) *DeleteHBaseHaDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteHBaseHaDBResponse) SetStatusCode(v int32) *DeleteHBaseHaDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHBaseHaDBResponse) SetBody(v *DeleteHBaseHaDBResponseBody) *DeleteHBaseHaDBResponse {
	s.Body = v
	return s
}

func (s *DeleteHBaseHaDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
