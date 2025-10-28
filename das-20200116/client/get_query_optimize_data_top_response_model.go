// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataTopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeDataTopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeDataTopResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeDataTopResponseBody) *GetQueryOptimizeDataTopResponse
	GetBody() *GetQueryOptimizeDataTopResponseBody
}

type GetQueryOptimizeDataTopResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeDataTopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeDataTopResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataTopResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataTopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeDataTopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeDataTopResponse) GetBody() *GetQueryOptimizeDataTopResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeDataTopResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeDataTopResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeDataTopResponse) SetStatusCode(v int32) *GetQueryOptimizeDataTopResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeDataTopResponse) SetBody(v *GetQueryOptimizeDataTopResponseBody) *GetQueryOptimizeDataTopResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeDataTopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
