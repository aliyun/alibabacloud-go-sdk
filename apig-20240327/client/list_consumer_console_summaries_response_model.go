// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerConsoleSummariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerConsoleSummariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerConsoleSummariesResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerConsoleSummariesResponseBody) *ListConsumerConsoleSummariesResponse
	GetBody() *ListConsumerConsoleSummariesResponseBody
}

type ListConsumerConsoleSummariesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerConsoleSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerConsoleSummariesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerConsoleSummariesResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerConsoleSummariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerConsoleSummariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerConsoleSummariesResponse) GetBody() *ListConsumerConsoleSummariesResponseBody {
	return s.Body
}

func (s *ListConsumerConsoleSummariesResponse) SetHeaders(v map[string]*string) *ListConsumerConsoleSummariesResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerConsoleSummariesResponse) SetStatusCode(v int32) *ListConsumerConsoleSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerConsoleSummariesResponse) SetBody(v *ListConsumerConsoleSummariesResponseBody) *ListConsumerConsoleSummariesResponse {
	s.Body = v
	return s
}

func (s *ListConsumerConsoleSummariesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
