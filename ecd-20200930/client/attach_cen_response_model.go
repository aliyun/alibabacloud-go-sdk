// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachCenResponse
	GetStatusCode() *int32
	SetBody(v *AttachCenResponseBody) *AttachCenResponse
	GetBody() *AttachCenResponseBody
}

type AttachCenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachCenResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachCenResponse) GoString() string {
	return s.String()
}

func (s *AttachCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachCenResponse) GetBody() *AttachCenResponseBody {
	return s.Body
}

func (s *AttachCenResponse) SetHeaders(v map[string]*string) *AttachCenResponse {
	s.Headers = v
	return s
}

func (s *AttachCenResponse) SetStatusCode(v int32) *AttachCenResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachCenResponse) SetBody(v *AttachCenResponseBody) *AttachCenResponse {
	s.Body = v
	return s
}

func (s *AttachCenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
