// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentAccuracyTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataAgentAccuracyTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataAgentAccuracyTestResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataAgentAccuracyTestResponseBody) *CreateDataAgentAccuracyTestResponse
	GetBody() *CreateDataAgentAccuracyTestResponseBody
}

type CreateDataAgentAccuracyTestResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataAgentAccuracyTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataAgentAccuracyTestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentAccuracyTestResponse) GoString() string {
	return s.String()
}

func (s *CreateDataAgentAccuracyTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataAgentAccuracyTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataAgentAccuracyTestResponse) GetBody() *CreateDataAgentAccuracyTestResponseBody {
	return s.Body
}

func (s *CreateDataAgentAccuracyTestResponse) SetHeaders(v map[string]*string) *CreateDataAgentAccuracyTestResponse {
	s.Headers = v
	return s
}

func (s *CreateDataAgentAccuracyTestResponse) SetStatusCode(v int32) *CreateDataAgentAccuracyTestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataAgentAccuracyTestResponse) SetBody(v *CreateDataAgentAccuracyTestResponseBody) *CreateDataAgentAccuracyTestResponse {
	s.Body = v
	return s
}

func (s *CreateDataAgentAccuracyTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
