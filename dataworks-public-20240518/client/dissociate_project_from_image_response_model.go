// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateProjectFromImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateProjectFromImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateProjectFromImageResponse
	GetStatusCode() *int32
	SetBody(v *DissociateProjectFromImageResponseBody) *DissociateProjectFromImageResponse
	GetBody() *DissociateProjectFromImageResponseBody
}

type DissociateProjectFromImageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateProjectFromImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateProjectFromImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateProjectFromImageResponse) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateProjectFromImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateProjectFromImageResponse) GetBody() *DissociateProjectFromImageResponseBody {
	return s.Body
}

func (s *DissociateProjectFromImageResponse) SetHeaders(v map[string]*string) *DissociateProjectFromImageResponse {
	s.Headers = v
	return s
}

func (s *DissociateProjectFromImageResponse) SetStatusCode(v int32) *DissociateProjectFromImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateProjectFromImageResponse) SetBody(v *DissociateProjectFromImageResponseBody) *DissociateProjectFromImageResponse {
	s.Body = v
	return s
}

func (s *DissociateProjectFromImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
