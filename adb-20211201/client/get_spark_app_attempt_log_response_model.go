// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppAttemptLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkAppAttemptLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkAppAttemptLogResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkAppAttemptLogResponseBody) *GetSparkAppAttemptLogResponse
	GetBody() *GetSparkAppAttemptLogResponseBody
}

type GetSparkAppAttemptLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppAttemptLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppAttemptLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppAttemptLogResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppAttemptLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkAppAttemptLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkAppAttemptLogResponse) GetBody() *GetSparkAppAttemptLogResponseBody {
	return s.Body
}

func (s *GetSparkAppAttemptLogResponse) SetHeaders(v map[string]*string) *GetSparkAppAttemptLogResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppAttemptLogResponse) SetStatusCode(v int32) *GetSparkAppAttemptLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppAttemptLogResponse) SetBody(v *GetSparkAppAttemptLogResponseBody) *GetSparkAppAttemptLogResponse {
	s.Body = v
	return s
}

func (s *GetSparkAppAttemptLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
