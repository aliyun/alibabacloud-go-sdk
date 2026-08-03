// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInsightSelectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInsightSelectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInsightSelectorsResponse
	GetStatusCode() *int32
	SetBody(v *GetInsightSelectorsResponseBody) *GetInsightSelectorsResponse
	GetBody() *GetInsightSelectorsResponseBody
}

type GetInsightSelectorsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInsightSelectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInsightSelectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInsightSelectorsResponse) GoString() string {
	return s.String()
}

func (s *GetInsightSelectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInsightSelectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInsightSelectorsResponse) GetBody() *GetInsightSelectorsResponseBody {
	return s.Body
}

func (s *GetInsightSelectorsResponse) SetHeaders(v map[string]*string) *GetInsightSelectorsResponse {
	s.Headers = v
	return s
}

func (s *GetInsightSelectorsResponse) SetStatusCode(v int32) *GetInsightSelectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInsightSelectorsResponse) SetBody(v *GetInsightSelectorsResponseBody) *GetInsightSelectorsResponse {
	s.Body = v
	return s
}

func (s *GetInsightSelectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
