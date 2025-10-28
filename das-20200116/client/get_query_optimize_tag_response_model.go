// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeTagResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeTagResponseBody) *GetQueryOptimizeTagResponse
	GetBody() *GetQueryOptimizeTagResponseBody
}

type GetQueryOptimizeTagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeTagResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeTagResponse) GetBody() *GetQueryOptimizeTagResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeTagResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeTagResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeTagResponse) SetStatusCode(v int32) *GetQueryOptimizeTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeTagResponse) SetBody(v *GetQueryOptimizeTagResponseBody) *GetQueryOptimizeTagResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
