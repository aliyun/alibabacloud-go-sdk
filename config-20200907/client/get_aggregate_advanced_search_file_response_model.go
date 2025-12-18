// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateAdvancedSearchFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateAdvancedSearchFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateAdvancedSearchFileResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateAdvancedSearchFileResponseBody) *GetAggregateAdvancedSearchFileResponse
	GetBody() *GetAggregateAdvancedSearchFileResponseBody
}

type GetAggregateAdvancedSearchFileResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateAdvancedSearchFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateAdvancedSearchFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAdvancedSearchFileResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateAdvancedSearchFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateAdvancedSearchFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateAdvancedSearchFileResponse) GetBody() *GetAggregateAdvancedSearchFileResponseBody {
	return s.Body
}

func (s *GetAggregateAdvancedSearchFileResponse) SetHeaders(v map[string]*string) *GetAggregateAdvancedSearchFileResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateAdvancedSearchFileResponse) SetStatusCode(v int32) *GetAggregateAdvancedSearchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateAdvancedSearchFileResponse) SetBody(v *GetAggregateAdvancedSearchFileResponseBody) *GetAggregateAdvancedSearchFileResponse {
	s.Body = v
	return s
}

func (s *GetAggregateAdvancedSearchFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
