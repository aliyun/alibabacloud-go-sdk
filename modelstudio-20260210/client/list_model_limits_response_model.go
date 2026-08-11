// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelLimitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelLimitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelLimitsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelLimitsResponseBody) *ListModelLimitsResponse
	GetBody() *ListModelLimitsResponseBody
}

type ListModelLimitsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelLimitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelLimitsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelLimitsResponse) GoString() string {
	return s.String()
}

func (s *ListModelLimitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelLimitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelLimitsResponse) GetBody() *ListModelLimitsResponseBody {
	return s.Body
}

func (s *ListModelLimitsResponse) SetHeaders(v map[string]*string) *ListModelLimitsResponse {
	s.Headers = v
	return s
}

func (s *ListModelLimitsResponse) SetStatusCode(v int32) *ListModelLimitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelLimitsResponse) SetBody(v *ListModelLimitsResponseBody) *ListModelLimitsResponse {
	s.Body = v
	return s
}

func (s *ListModelLimitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
