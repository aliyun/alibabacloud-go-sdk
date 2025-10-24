// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFreeUserEventTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFreeUserEventTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFreeUserEventTypesResponseBody) *DescribeFreeUserEventTypesResponse
	GetBody() *DescribeFreeUserEventTypesResponseBody
}

type DescribeFreeUserEventTypesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFreeUserEventTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFreeUserEventTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFreeUserEventTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFreeUserEventTypesResponse) GetBody() *DescribeFreeUserEventTypesResponseBody {
	return s.Body
}

func (s *DescribeFreeUserEventTypesResponse) SetHeaders(v map[string]*string) *DescribeFreeUserEventTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFreeUserEventTypesResponse) SetStatusCode(v int32) *DescribeFreeUserEventTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFreeUserEventTypesResponse) SetBody(v *DescribeFreeUserEventTypesResponseBody) *DescribeFreeUserEventTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeFreeUserEventTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
