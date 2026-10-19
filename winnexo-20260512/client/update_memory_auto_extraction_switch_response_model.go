// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryAutoExtractionSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMemoryAutoExtractionSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMemoryAutoExtractionSwitchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMemoryAutoExtractionSwitchResponseBody) *UpdateMemoryAutoExtractionSwitchResponse
	GetBody() *UpdateMemoryAutoExtractionSwitchResponseBody
}

type UpdateMemoryAutoExtractionSwitchResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemoryAutoExtractionSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryAutoExtractionSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryAutoExtractionSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemoryAutoExtractionSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMemoryAutoExtractionSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMemoryAutoExtractionSwitchResponse) GetBody() *UpdateMemoryAutoExtractionSwitchResponseBody {
	return s.Body
}

func (s *UpdateMemoryAutoExtractionSwitchResponse) SetHeaders(v map[string]*string) *UpdateMemoryAutoExtractionSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponse) SetStatusCode(v int32) *UpdateMemoryAutoExtractionSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponse) SetBody(v *UpdateMemoryAutoExtractionSwitchResponseBody) *UpdateMemoryAutoExtractionSwitchResponse {
	s.Body = v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
