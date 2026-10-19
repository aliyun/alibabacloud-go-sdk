// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryAutoExtractionSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryAutoExtractionSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryAutoExtractionSwitchResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryAutoExtractionSwitchResponseBody) *GetMemoryAutoExtractionSwitchResponse
	GetBody() *GetMemoryAutoExtractionSwitchResponseBody
}

type GetMemoryAutoExtractionSwitchResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryAutoExtractionSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryAutoExtractionSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryAutoExtractionSwitchResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryAutoExtractionSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryAutoExtractionSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryAutoExtractionSwitchResponse) GetBody() *GetMemoryAutoExtractionSwitchResponseBody {
	return s.Body
}

func (s *GetMemoryAutoExtractionSwitchResponse) SetHeaders(v map[string]*string) *GetMemoryAutoExtractionSwitchResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponse) SetStatusCode(v int32) *GetMemoryAutoExtractionSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponse) SetBody(v *GetMemoryAutoExtractionSwitchResponseBody) *GetMemoryAutoExtractionSwitchResponse {
	s.Body = v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
