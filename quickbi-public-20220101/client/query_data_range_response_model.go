// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataRangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDataRangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDataRangeResponse
	GetStatusCode() *int32
	SetBody(v *QueryDataRangeResponseBody) *QueryDataRangeResponse
	GetBody() *QueryDataRangeResponseBody
}

type QueryDataRangeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataRangeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDataRangeResponse) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDataRangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDataRangeResponse) GetBody() *QueryDataRangeResponseBody {
	return s.Body
}

func (s *QueryDataRangeResponse) SetHeaders(v map[string]*string) *QueryDataRangeResponse {
	s.Headers = v
	return s
}

func (s *QueryDataRangeResponse) SetStatusCode(v int32) *QueryDataRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataRangeResponse) SetBody(v *QueryDataRangeResponseBody) *QueryDataRangeResponse {
	s.Body = v
	return s
}

func (s *QueryDataRangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
