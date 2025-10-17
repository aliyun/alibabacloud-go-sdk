// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsCopyWorkloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApsCopyWorkloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApsCopyWorkloadResponse
	GetStatusCode() *int32
	SetBody(v *CreateApsCopyWorkloadResponseBody) *CreateApsCopyWorkloadResponse
	GetBody() *CreateApsCopyWorkloadResponseBody
}

type CreateApsCopyWorkloadResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApsCopyWorkloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApsCopyWorkloadResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApsCopyWorkloadResponse) GoString() string {
	return s.String()
}

func (s *CreateApsCopyWorkloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApsCopyWorkloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApsCopyWorkloadResponse) GetBody() *CreateApsCopyWorkloadResponseBody {
	return s.Body
}

func (s *CreateApsCopyWorkloadResponse) SetHeaders(v map[string]*string) *CreateApsCopyWorkloadResponse {
	s.Headers = v
	return s
}

func (s *CreateApsCopyWorkloadResponse) SetStatusCode(v int32) *CreateApsCopyWorkloadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApsCopyWorkloadResponse) SetBody(v *CreateApsCopyWorkloadResponseBody) *CreateApsCopyWorkloadResponse {
	s.Body = v
	return s
}

func (s *CreateApsCopyWorkloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
