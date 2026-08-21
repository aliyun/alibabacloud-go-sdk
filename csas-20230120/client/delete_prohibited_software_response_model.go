// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProhibitedSoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProhibitedSoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProhibitedSoftwareResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProhibitedSoftwareResponseBody) *DeleteProhibitedSoftwareResponse
	GetBody() *DeleteProhibitedSoftwareResponseBody
}

type DeleteProhibitedSoftwareResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProhibitedSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProhibitedSoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProhibitedSoftwareResponse) GoString() string {
	return s.String()
}

func (s *DeleteProhibitedSoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProhibitedSoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProhibitedSoftwareResponse) GetBody() *DeleteProhibitedSoftwareResponseBody {
	return s.Body
}

func (s *DeleteProhibitedSoftwareResponse) SetHeaders(v map[string]*string) *DeleteProhibitedSoftwareResponse {
	s.Headers = v
	return s
}

func (s *DeleteProhibitedSoftwareResponse) SetStatusCode(v int32) *DeleteProhibitedSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProhibitedSoftwareResponse) SetBody(v *DeleteProhibitedSoftwareResponseBody) *DeleteProhibitedSoftwareResponse {
	s.Body = v
	return s
}

func (s *DeleteProhibitedSoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
