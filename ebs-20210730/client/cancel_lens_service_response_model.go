// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelLensServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelLensServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelLensServiceResponse
	GetStatusCode() *int32
	SetBody(v *CancelLensServiceResponseBody) *CancelLensServiceResponse
	GetBody() *CancelLensServiceResponseBody
}

type CancelLensServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelLensServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelLensServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelLensServiceResponse) GoString() string {
	return s.String()
}

func (s *CancelLensServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelLensServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelLensServiceResponse) GetBody() *CancelLensServiceResponseBody {
	return s.Body
}

func (s *CancelLensServiceResponse) SetHeaders(v map[string]*string) *CancelLensServiceResponse {
	s.Headers = v
	return s
}

func (s *CancelLensServiceResponse) SetStatusCode(v int32) *CancelLensServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelLensServiceResponse) SetBody(v *CancelLensServiceResponseBody) *CancelLensServiceResponse {
	s.Body = v
	return s
}

func (s *CancelLensServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
