// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseOrgsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLangfuseOrgsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLangfuseOrgsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLangfuseOrgsResponseBody) *DescribeLangfuseOrgsResponse
	GetBody() *DescribeLangfuseOrgsResponseBody
}

type DescribeLangfuseOrgsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLangfuseOrgsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLangfuseOrgsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLangfuseOrgsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLangfuseOrgsResponse) GetBody() *DescribeLangfuseOrgsResponseBody {
	return s.Body
}

func (s *DescribeLangfuseOrgsResponse) SetHeaders(v map[string]*string) *DescribeLangfuseOrgsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLangfuseOrgsResponse) SetStatusCode(v int32) *DescribeLangfuseOrgsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLangfuseOrgsResponse) SetBody(v *DescribeLangfuseOrgsResponseBody) *DescribeLangfuseOrgsResponse {
	s.Body = v
	return s
}

func (s *DescribeLangfuseOrgsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
