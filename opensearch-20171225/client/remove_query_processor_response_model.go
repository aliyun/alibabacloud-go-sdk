// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveQueryProcessorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveQueryProcessorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveQueryProcessorResponse
	GetStatusCode() *int32
	SetBody(v *RemoveQueryProcessorResponseBody) *RemoveQueryProcessorResponse
	GetBody() *RemoveQueryProcessorResponseBody
}

type RemoveQueryProcessorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveQueryProcessorResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *RemoveQueryProcessorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveQueryProcessorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveQueryProcessorResponse) GetBody() *RemoveQueryProcessorResponseBody {
	return s.Body
}

func (s *RemoveQueryProcessorResponse) SetHeaders(v map[string]*string) *RemoveQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *RemoveQueryProcessorResponse) SetStatusCode(v int32) *RemoveQueryProcessorResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveQueryProcessorResponse) SetBody(v *RemoveQueryProcessorResponseBody) *RemoveQueryProcessorResponse {
	s.Body = v
	return s
}

func (s *RemoveQueryProcessorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
