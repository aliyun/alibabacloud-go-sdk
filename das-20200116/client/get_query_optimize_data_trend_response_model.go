// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeDataTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeDataTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeDataTrendResponseBody) *GetQueryOptimizeDataTrendResponse
	GetBody() *GetQueryOptimizeDataTrendResponseBody
}

type GetQueryOptimizeDataTrendResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeDataTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeDataTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTrendResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeDataTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeDataTrendResponse) GetBody() *GetQueryOptimizeDataTrendResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeDataTrendResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeDataTrendResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeDataTrendResponse) SetStatusCode(v int32) *GetQueryOptimizeDataTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeDataTrendResponse) SetBody(v *GetQueryOptimizeDataTrendResponseBody) *GetQueryOptimizeDataTrendResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeDataTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
