// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEaisEiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachEaisEiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachEaisEiResponse
	GetStatusCode() *int32
	SetBody(v *AttachEaisEiResponseBody) *AttachEaisEiResponse
	GetBody() *AttachEaisEiResponseBody
}

type AttachEaisEiResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachEaisEiResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachEaisEiResponse) GoString() string {
	return s.String()
}

func (s *AttachEaisEiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachEaisEiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachEaisEiResponse) GetBody() *AttachEaisEiResponseBody {
	return s.Body
}

func (s *AttachEaisEiResponse) SetHeaders(v map[string]*string) *AttachEaisEiResponse {
	s.Headers = v
	return s
}

func (s *AttachEaisEiResponse) SetStatusCode(v int32) *AttachEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEaisEiResponse) SetBody(v *AttachEaisEiResponseBody) *AttachEaisEiResponse {
	s.Body = v
	return s
}

func (s *AttachEaisEiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
