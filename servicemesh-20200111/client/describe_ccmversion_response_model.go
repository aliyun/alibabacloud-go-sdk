// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCCMVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCCMVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCCMVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCCMVersionResponseBody) *DescribeCCMVersionResponse
	GetBody() *DescribeCCMVersionResponseBody
}

type DescribeCCMVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCCMVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCCMVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCCMVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeCCMVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCCMVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCCMVersionResponse) GetBody() *DescribeCCMVersionResponseBody {
	return s.Body
}

func (s *DescribeCCMVersionResponse) SetHeaders(v map[string]*string) *DescribeCCMVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeCCMVersionResponse) SetStatusCode(v int32) *DescribeCCMVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCCMVersionResponse) SetBody(v *DescribeCCMVersionResponseBody) *DescribeCCMVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeCCMVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
