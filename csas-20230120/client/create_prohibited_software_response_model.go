// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedSoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProhibitedSoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProhibitedSoftwareResponse
	GetStatusCode() *int32
	SetBody(v *CreateProhibitedSoftwareResponseBody) *CreateProhibitedSoftwareResponse
	GetBody() *CreateProhibitedSoftwareResponseBody
}

type CreateProhibitedSoftwareResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProhibitedSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProhibitedSoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedSoftwareResponse) GoString() string {
	return s.String()
}

func (s *CreateProhibitedSoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProhibitedSoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProhibitedSoftwareResponse) GetBody() *CreateProhibitedSoftwareResponseBody {
	return s.Body
}

func (s *CreateProhibitedSoftwareResponse) SetHeaders(v map[string]*string) *CreateProhibitedSoftwareResponse {
	s.Headers = v
	return s
}

func (s *CreateProhibitedSoftwareResponse) SetStatusCode(v int32) *CreateProhibitedSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProhibitedSoftwareResponse) SetBody(v *CreateProhibitedSoftwareResponseBody) *CreateProhibitedSoftwareResponse {
	s.Body = v
	return s
}

func (s *CreateProhibitedSoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
