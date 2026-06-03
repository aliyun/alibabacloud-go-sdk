// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcknowledgeTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcknowledgeTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcknowledgeTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *AcknowledgeTaskResultResponseBody) *AcknowledgeTaskResultResponse
	GetBody() *AcknowledgeTaskResultResponseBody
}

type AcknowledgeTaskResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcknowledgeTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcknowledgeTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s AcknowledgeTaskResultResponse) GoString() string {
	return s.String()
}

func (s *AcknowledgeTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcknowledgeTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcknowledgeTaskResultResponse) GetBody() *AcknowledgeTaskResultResponseBody {
	return s.Body
}

func (s *AcknowledgeTaskResultResponse) SetHeaders(v map[string]*string) *AcknowledgeTaskResultResponse {
	s.Headers = v
	return s
}

func (s *AcknowledgeTaskResultResponse) SetStatusCode(v int32) *AcknowledgeTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *AcknowledgeTaskResultResponse) SetBody(v *AcknowledgeTaskResultResponseBody) *AcknowledgeTaskResultResponse {
	s.Body = v
	return s
}

func (s *AcknowledgeTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
