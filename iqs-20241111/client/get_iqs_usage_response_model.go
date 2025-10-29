// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIqsUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIqsUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIqsUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetIqsUsageResult) *GetIqsUsageResponse
	GetBody() *GetIqsUsageResult
}

type GetIqsUsageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIqsUsageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIqsUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIqsUsageResponse) GoString() string {
	return s.String()
}

func (s *GetIqsUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIqsUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIqsUsageResponse) GetBody() *GetIqsUsageResult {
	return s.Body
}

func (s *GetIqsUsageResponse) SetHeaders(v map[string]*string) *GetIqsUsageResponse {
	s.Headers = v
	return s
}

func (s *GetIqsUsageResponse) SetStatusCode(v int32) *GetIqsUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIqsUsageResponse) SetBody(v *GetIqsUsageResult) *GetIqsUsageResponse {
	s.Body = v
	return s
}

func (s *GetIqsUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
