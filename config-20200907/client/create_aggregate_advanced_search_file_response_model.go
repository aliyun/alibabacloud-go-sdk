// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateAdvancedSearchFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAggregateAdvancedSearchFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAggregateAdvancedSearchFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateAggregateAdvancedSearchFileResponseBody) *CreateAggregateAdvancedSearchFileResponse
	GetBody() *CreateAggregateAdvancedSearchFileResponseBody
}

type CreateAggregateAdvancedSearchFileResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAggregateAdvancedSearchFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAggregateAdvancedSearchFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateAdvancedSearchFileResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateAdvancedSearchFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAggregateAdvancedSearchFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAggregateAdvancedSearchFileResponse) GetBody() *CreateAggregateAdvancedSearchFileResponseBody {
	return s.Body
}

func (s *CreateAggregateAdvancedSearchFileResponse) SetHeaders(v map[string]*string) *CreateAggregateAdvancedSearchFileResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateAdvancedSearchFileResponse) SetStatusCode(v int32) *CreateAggregateAdvancedSearchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateAdvancedSearchFileResponse) SetBody(v *CreateAggregateAdvancedSearchFileResponseBody) *CreateAggregateAdvancedSearchFileResponse {
	s.Body = v
	return s
}

func (s *CreateAggregateAdvancedSearchFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
