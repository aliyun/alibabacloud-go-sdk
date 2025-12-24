// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPredictiveValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPredictiveValueResponse
	GetStatusCode() *int32
	SetBody(v *QueryPredictiveValueResponseBody) *QueryPredictiveValueResponse
	GetBody() *QueryPredictiveValueResponseBody
}

type QueryPredictiveValueResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPredictiveValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPredictiveValueResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveValueResponse) GoString() string {
	return s.String()
}

func (s *QueryPredictiveValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPredictiveValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPredictiveValueResponse) GetBody() *QueryPredictiveValueResponseBody {
	return s.Body
}

func (s *QueryPredictiveValueResponse) SetHeaders(v map[string]*string) *QueryPredictiveValueResponse {
	s.Headers = v
	return s
}

func (s *QueryPredictiveValueResponse) SetStatusCode(v int32) *QueryPredictiveValueResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPredictiveValueResponse) SetBody(v *QueryPredictiveValueResponseBody) *QueryPredictiveValueResponse {
	s.Body = v
	return s
}

func (s *QueryPredictiveValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
