// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConcernNecessityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConcernNecessityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConcernNecessityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConcernNecessityResponseBody) *DescribeConcernNecessityResponse
	GetBody() *DescribeConcernNecessityResponseBody
}

type DescribeConcernNecessityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConcernNecessityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConcernNecessityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConcernNecessityResponse) GoString() string {
	return s.String()
}

func (s *DescribeConcernNecessityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConcernNecessityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConcernNecessityResponse) GetBody() *DescribeConcernNecessityResponseBody {
	return s.Body
}

func (s *DescribeConcernNecessityResponse) SetHeaders(v map[string]*string) *DescribeConcernNecessityResponse {
	s.Headers = v
	return s
}

func (s *DescribeConcernNecessityResponse) SetStatusCode(v int32) *DescribeConcernNecessityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConcernNecessityResponse) SetBody(v *DescribeConcernNecessityResponseBody) *DescribeConcernNecessityResponse {
	s.Body = v
	return s
}

func (s *DescribeConcernNecessityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
