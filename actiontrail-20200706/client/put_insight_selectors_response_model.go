// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutInsightSelectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutInsightSelectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutInsightSelectorsResponse
	GetStatusCode() *int32
	SetBody(v *PutInsightSelectorsResponseBody) *PutInsightSelectorsResponse
	GetBody() *PutInsightSelectorsResponseBody
}

type PutInsightSelectorsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutInsightSelectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutInsightSelectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutInsightSelectorsResponse) GoString() string {
	return s.String()
}

func (s *PutInsightSelectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutInsightSelectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutInsightSelectorsResponse) GetBody() *PutInsightSelectorsResponseBody {
	return s.Body
}

func (s *PutInsightSelectorsResponse) SetHeaders(v map[string]*string) *PutInsightSelectorsResponse {
	s.Headers = v
	return s
}

func (s *PutInsightSelectorsResponse) SetStatusCode(v int32) *PutInsightSelectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutInsightSelectorsResponse) SetBody(v *PutInsightSelectorsResponseBody) *PutInsightSelectorsResponse {
	s.Body = v
	return s
}

func (s *PutInsightSelectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
