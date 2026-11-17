// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySourceUsageTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySourceUsageTrendResponse
	GetStatusCode() *int32
	SetBody(v *QuerySourceUsageTrendResponseBody) *QuerySourceUsageTrendResponse
	GetBody() *QuerySourceUsageTrendResponseBody
}

type QuerySourceUsageTrendResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySourceUsageTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySourceUsageTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageTrendResponse) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySourceUsageTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySourceUsageTrendResponse) GetBody() *QuerySourceUsageTrendResponseBody {
	return s.Body
}

func (s *QuerySourceUsageTrendResponse) SetHeaders(v map[string]*string) *QuerySourceUsageTrendResponse {
	s.Headers = v
	return s
}

func (s *QuerySourceUsageTrendResponse) SetStatusCode(v int32) *QuerySourceUsageTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySourceUsageTrendResponse) SetBody(v *QuerySourceUsageTrendResponseBody) *QuerySourceUsageTrendResponse {
	s.Body = v
	return s
}

func (s *QuerySourceUsageTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
