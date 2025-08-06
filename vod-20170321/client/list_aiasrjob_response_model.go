// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIASRJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIASRJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIASRJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIASRJobResponseBody) *ListAIASRJobResponse
	GetBody() *ListAIASRJobResponseBody
}

type ListAIASRJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIASRJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIASRJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIASRJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIASRJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIASRJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIASRJobResponse) GetBody() *ListAIASRJobResponseBody {
	return s.Body
}

func (s *ListAIASRJobResponse) SetHeaders(v map[string]*string) *ListAIASRJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIASRJobResponse) SetStatusCode(v int32) *ListAIASRJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIASRJobResponse) SetBody(v *ListAIASRJobResponseBody) *ListAIASRJobResponse {
	s.Body = v
	return s
}

func (s *ListAIASRJobResponse) Validate() error {
	return dara.Validate(s)
}
