// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoTerrorismRecogJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIVideoTerrorismRecogJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIVideoTerrorismRecogJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIVideoTerrorismRecogJobResponseBody) *ListAIVideoTerrorismRecogJobResponse
	GetBody() *ListAIVideoTerrorismRecogJobResponseBody
}

type ListAIVideoTerrorismRecogJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIVideoTerrorismRecogJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIVideoTerrorismRecogJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTerrorismRecogJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIVideoTerrorismRecogJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIVideoTerrorismRecogJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIVideoTerrorismRecogJobResponse) GetBody() *ListAIVideoTerrorismRecogJobResponseBody {
	return s.Body
}

func (s *ListAIVideoTerrorismRecogJobResponse) SetHeaders(v map[string]*string) *ListAIVideoTerrorismRecogJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponse) SetStatusCode(v int32) *ListAIVideoTerrorismRecogJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponse) SetBody(v *ListAIVideoTerrorismRecogJobResponseBody) *ListAIVideoTerrorismRecogJobResponse {
	s.Body = v
	return s
}

func (s *ListAIVideoTerrorismRecogJobResponse) Validate() error {
	return dara.Validate(s)
}
