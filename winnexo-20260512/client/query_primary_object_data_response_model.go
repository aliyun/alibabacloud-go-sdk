// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPrimaryObjectDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPrimaryObjectDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPrimaryObjectDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryPrimaryObjectDataResponseBody) *QueryPrimaryObjectDataResponse
	GetBody() *QueryPrimaryObjectDataResponseBody
}

type QueryPrimaryObjectDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPrimaryObjectDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPrimaryObjectDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPrimaryObjectDataResponse) GoString() string {
	return s.String()
}

func (s *QueryPrimaryObjectDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPrimaryObjectDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPrimaryObjectDataResponse) GetBody() *QueryPrimaryObjectDataResponseBody {
	return s.Body
}

func (s *QueryPrimaryObjectDataResponse) SetHeaders(v map[string]*string) *QueryPrimaryObjectDataResponse {
	s.Headers = v
	return s
}

func (s *QueryPrimaryObjectDataResponse) SetStatusCode(v int32) *QueryPrimaryObjectDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPrimaryObjectDataResponse) SetBody(v *QueryPrimaryObjectDataResponseBody) *QueryPrimaryObjectDataResponse {
	s.Body = v
	return s
}

func (s *QueryPrimaryObjectDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
