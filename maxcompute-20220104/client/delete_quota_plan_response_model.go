// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQuotaPlanResponseBody) *DeleteQuotaPlanResponse
	GetBody() *DeleteQuotaPlanResponseBody
}

type DeleteQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQuotaPlanResponse) GetBody() *DeleteQuotaPlanResponseBody {
	return s.Body
}

func (s *DeleteQuotaPlanResponse) SetHeaders(v map[string]*string) *DeleteQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuotaPlanResponse) SetStatusCode(v int32) *DeleteQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQuotaPlanResponse) SetBody(v *DeleteQuotaPlanResponseBody) *DeleteQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteQuotaPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
