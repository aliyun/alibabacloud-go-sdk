// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetRegionResponseBody) *GetRegionResponse
	GetBody() *GetRegionResponseBody
}

type GetRegionResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegionResponse) GoString() string {
	return s.String()
}

func (s *GetRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegionResponse) GetBody() *GetRegionResponseBody {
	return s.Body
}

func (s *GetRegionResponse) SetHeaders(v map[string]*string) *GetRegionResponse {
	s.Headers = v
	return s
}

func (s *GetRegionResponse) SetStatusCode(v int32) *GetRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionResponse) SetBody(v *GetRegionResponseBody) *GetRegionResponse {
	s.Body = v
	return s
}

func (s *GetRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
