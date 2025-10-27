// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarSecurityEventsQueryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSimilarSecurityEventsQueryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSimilarSecurityEventsQueryTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSimilarSecurityEventsQueryTaskResponseBody) *CreateSimilarSecurityEventsQueryTaskResponse
	GetBody() *CreateSimilarSecurityEventsQueryTaskResponseBody
}

type CreateSimilarSecurityEventsQueryTaskResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimilarSecurityEventsQueryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimilarSecurityEventsQueryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) GetBody() *CreateSimilarSecurityEventsQueryTaskResponseBody {
	return s.Body
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) SetHeaders(v map[string]*string) *CreateSimilarSecurityEventsQueryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) SetStatusCode(v int32) *CreateSimilarSecurityEventsQueryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) SetBody(v *CreateSimilarSecurityEventsQueryTaskResponseBody) *CreateSimilarSecurityEventsQueryTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
