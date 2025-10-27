// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQAConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQAConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQAConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQAConfigResponseBody) *DescribeSQAConfigResponse
	GetBody() *DescribeSQAConfigResponseBody
}

type DescribeSQAConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQAConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQAConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQAConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQAConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQAConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQAConfigResponse) GetBody() *DescribeSQAConfigResponseBody {
	return s.Body
}

func (s *DescribeSQAConfigResponse) SetHeaders(v map[string]*string) *DescribeSQAConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQAConfigResponse) SetStatusCode(v int32) *DescribeSQAConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQAConfigResponse) SetBody(v *DescribeSQAConfigResponseBody) *DescribeSQAConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSQAConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
