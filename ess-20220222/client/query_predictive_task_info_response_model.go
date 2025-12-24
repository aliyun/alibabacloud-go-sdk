// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPredictiveTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPredictiveTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryPredictiveTaskInfoResponseBody) *QueryPredictiveTaskInfoResponse
	GetBody() *QueryPredictiveTaskInfoResponseBody
}

type QueryPredictiveTaskInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPredictiveTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPredictiveTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPredictiveTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPredictiveTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPredictiveTaskInfoResponse) GetBody() *QueryPredictiveTaskInfoResponseBody {
	return s.Body
}

func (s *QueryPredictiveTaskInfoResponse) SetHeaders(v map[string]*string) *QueryPredictiveTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPredictiveTaskInfoResponse) SetStatusCode(v int32) *QueryPredictiveTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPredictiveTaskInfoResponse) SetBody(v *QueryPredictiveTaskInfoResponseBody) *QueryPredictiveTaskInfoResponse {
	s.Body = v
	return s
}

func (s *QueryPredictiveTaskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
