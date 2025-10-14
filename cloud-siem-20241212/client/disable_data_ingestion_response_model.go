// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDataIngestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDataIngestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDataIngestionResponse
	GetStatusCode() *int32
	SetBody(v *DisableDataIngestionResponseBody) *DisableDataIngestionResponse
	GetBody() *DisableDataIngestionResponseBody
}

type DisableDataIngestionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDataIngestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDataIngestionResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDataIngestionResponse) GoString() string {
	return s.String()
}

func (s *DisableDataIngestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDataIngestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDataIngestionResponse) GetBody() *DisableDataIngestionResponseBody {
	return s.Body
}

func (s *DisableDataIngestionResponse) SetHeaders(v map[string]*string) *DisableDataIngestionResponse {
	s.Headers = v
	return s
}

func (s *DisableDataIngestionResponse) SetStatusCode(v int32) *DisableDataIngestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDataIngestionResponse) SetBody(v *DisableDataIngestionResponseBody) *DisableDataIngestionResponse {
	s.Body = v
	return s
}

func (s *DisableDataIngestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
