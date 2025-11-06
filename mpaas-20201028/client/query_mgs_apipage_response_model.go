// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsApipageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMgsApipageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMgsApipageResponse
	GetStatusCode() *int32
	SetBody(v *QueryMgsApipageResponseBody) *QueryMgsApipageResponse
	GetBody() *QueryMgsApipageResponseBody
}

type QueryMgsApipageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMgsApipageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMgsApipageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsApipageResponse) GoString() string {
	return s.String()
}

func (s *QueryMgsApipageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMgsApipageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMgsApipageResponse) GetBody() *QueryMgsApipageResponseBody {
	return s.Body
}

func (s *QueryMgsApipageResponse) SetHeaders(v map[string]*string) *QueryMgsApipageResponse {
	s.Headers = v
	return s
}

func (s *QueryMgsApipageResponse) SetStatusCode(v int32) *QueryMgsApipageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMgsApipageResponse) SetBody(v *QueryMgsApipageResponseBody) *QueryMgsApipageResponse {
	s.Body = v
	return s
}

func (s *QueryMgsApipageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
