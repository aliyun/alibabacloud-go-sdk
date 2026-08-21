// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProhibitedSoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProhibitedSoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProhibitedSoftwareResponse
	GetStatusCode() *int32
	SetBody(v *GetProhibitedSoftwareResponseBody) *GetProhibitedSoftwareResponse
	GetBody() *GetProhibitedSoftwareResponseBody
}

type GetProhibitedSoftwareResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProhibitedSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProhibitedSoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedSoftwareResponse) GoString() string {
	return s.String()
}

func (s *GetProhibitedSoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProhibitedSoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProhibitedSoftwareResponse) GetBody() *GetProhibitedSoftwareResponseBody {
	return s.Body
}

func (s *GetProhibitedSoftwareResponse) SetHeaders(v map[string]*string) *GetProhibitedSoftwareResponse {
	s.Headers = v
	return s
}

func (s *GetProhibitedSoftwareResponse) SetStatusCode(v int32) *GetProhibitedSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProhibitedSoftwareResponse) SetBody(v *GetProhibitedSoftwareResponseBody) *GetProhibitedSoftwareResponse {
	s.Body = v
	return s
}

func (s *GetProhibitedSoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
