// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMMLActiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMMLActiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMMLActiveResponse
	GetStatusCode() *int32
	SetBody(v *QueryMMLActiveResponseBody) *QueryMMLActiveResponse
	GetBody() *QueryMMLActiveResponseBody
}

type QueryMMLActiveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMMLActiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMMLActiveResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMMLActiveResponse) GoString() string {
	return s.String()
}

func (s *QueryMMLActiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMMLActiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMMLActiveResponse) GetBody() *QueryMMLActiveResponseBody {
	return s.Body
}

func (s *QueryMMLActiveResponse) SetHeaders(v map[string]*string) *QueryMMLActiveResponse {
	s.Headers = v
	return s
}

func (s *QueryMMLActiveResponse) SetStatusCode(v int32) *QueryMMLActiveResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMMLActiveResponse) SetBody(v *QueryMMLActiveResponseBody) *QueryMMLActiveResponse {
	s.Body = v
	return s
}

func (s *QueryMMLActiveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
