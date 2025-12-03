// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryXpackRelatedDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryXpackRelatedDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryXpackRelatedDBResponse
	GetStatusCode() *int32
	SetBody(v *QueryXpackRelatedDBResponseBody) *QueryXpackRelatedDBResponse
	GetBody() *QueryXpackRelatedDBResponseBody
}

type QueryXpackRelatedDBResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryXpackRelatedDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryXpackRelatedDBResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelatedDBResponse) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryXpackRelatedDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryXpackRelatedDBResponse) GetBody() *QueryXpackRelatedDBResponseBody {
	return s.Body
}

func (s *QueryXpackRelatedDBResponse) SetHeaders(v map[string]*string) *QueryXpackRelatedDBResponse {
	s.Headers = v
	return s
}

func (s *QueryXpackRelatedDBResponse) SetStatusCode(v int32) *QueryXpackRelatedDBResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryXpackRelatedDBResponse) SetBody(v *QueryXpackRelatedDBResponseBody) *QueryXpackRelatedDBResponse {
	s.Body = v
	return s
}

func (s *QueryXpackRelatedDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
