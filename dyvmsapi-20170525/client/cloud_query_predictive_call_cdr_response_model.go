// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryPredictiveCallCdrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryPredictiveCallCdrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryPredictiveCallCdrResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryPredictiveCallCdrResponseBody) *CloudQueryPredictiveCallCdrResponse
	GetBody() *CloudQueryPredictiveCallCdrResponseBody
}

type CloudQueryPredictiveCallCdrResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryPredictiveCallCdrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryPredictiveCallCdrResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryPredictiveCallCdrResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryPredictiveCallCdrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryPredictiveCallCdrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryPredictiveCallCdrResponse) GetBody() *CloudQueryPredictiveCallCdrResponseBody {
	return s.Body
}

func (s *CloudQueryPredictiveCallCdrResponse) SetHeaders(v map[string]*string) *CloudQueryPredictiveCallCdrResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponse) SetStatusCode(v int32) *CloudQueryPredictiveCallCdrResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponse) SetBody(v *CloudQueryPredictiveCallCdrResponseBody) *CloudQueryPredictiveCallCdrResponse {
	s.Body = v
	return s
}

func (s *CloudQueryPredictiveCallCdrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
