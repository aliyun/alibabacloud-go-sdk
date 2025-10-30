// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafStartConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSafStartConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSafStartConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSafStartConfigResponseBody) *DescribeSafStartConfigResponse
	GetBody() *DescribeSafStartConfigResponseBody
}

type DescribeSafStartConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSafStartConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSafStartConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafStartConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSafStartConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSafStartConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSafStartConfigResponse) GetBody() *DescribeSafStartConfigResponseBody {
	return s.Body
}

func (s *DescribeSafStartConfigResponse) SetHeaders(v map[string]*string) *DescribeSafStartConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSafStartConfigResponse) SetStatusCode(v int32) *DescribeSafStartConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSafStartConfigResponse) SetBody(v *DescribeSafStartConfigResponseBody) *DescribeSafStartConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSafStartConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
