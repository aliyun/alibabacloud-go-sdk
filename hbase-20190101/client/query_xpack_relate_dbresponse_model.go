// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryXpackRelateDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryXpackRelateDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryXpackRelateDBResponse
	GetStatusCode() *int32
	SetBody(v *QueryXpackRelateDBResponseBody) *QueryXpackRelateDBResponse
	GetBody() *QueryXpackRelateDBResponseBody
}

type QueryXpackRelateDBResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryXpackRelateDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryXpackRelateDBResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelateDBResponse) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryXpackRelateDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryXpackRelateDBResponse) GetBody() *QueryXpackRelateDBResponseBody {
	return s.Body
}

func (s *QueryXpackRelateDBResponse) SetHeaders(v map[string]*string) *QueryXpackRelateDBResponse {
	s.Headers = v
	return s
}

func (s *QueryXpackRelateDBResponse) SetStatusCode(v int32) *QueryXpackRelateDBResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryXpackRelateDBResponse) SetBody(v *QueryXpackRelateDBResponseBody) *QueryXpackRelateDBResponse {
	s.Body = v
	return s
}

func (s *QueryXpackRelateDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
