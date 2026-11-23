// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseUsageResponseBody) *GetContextDatabaseUsageResponse
	GetBody() *GetContextDatabaseUsageResponseBody
}

type GetContextDatabaseUsageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseUsageResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseUsageResponse) GetBody() *GetContextDatabaseUsageResponseBody {
	return s.Body
}

func (s *GetContextDatabaseUsageResponse) SetHeaders(v map[string]*string) *GetContextDatabaseUsageResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseUsageResponse) SetStatusCode(v int32) *GetContextDatabaseUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseUsageResponse) SetBody(v *GetContextDatabaseUsageResponseBody) *GetContextDatabaseUsageResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
