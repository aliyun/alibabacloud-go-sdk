// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedSoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProhibitedSoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProhibitedSoftwareResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProhibitedSoftwareResponseBody) *UpdateProhibitedSoftwareResponse
	GetBody() *UpdateProhibitedSoftwareResponseBody
}

type UpdateProhibitedSoftwareResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProhibitedSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProhibitedSoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedSoftwareResponse) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedSoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProhibitedSoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProhibitedSoftwareResponse) GetBody() *UpdateProhibitedSoftwareResponseBody {
	return s.Body
}

func (s *UpdateProhibitedSoftwareResponse) SetHeaders(v map[string]*string) *UpdateProhibitedSoftwareResponse {
	s.Headers = v
	return s
}

func (s *UpdateProhibitedSoftwareResponse) SetStatusCode(v int32) *UpdateProhibitedSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProhibitedSoftwareResponse) SetBody(v *UpdateProhibitedSoftwareResponseBody) *UpdateProhibitedSoftwareResponse {
	s.Body = v
	return s
}

func (s *UpdateProhibitedSoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
