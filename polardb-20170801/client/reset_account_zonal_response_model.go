// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAccountZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAccountZonalResponse
	GetStatusCode() *int32
	SetBody(v *ResetAccountZonalResponseBody) *ResetAccountZonalResponse
	GetBody() *ResetAccountZonalResponseBody
}

type ResetAccountZonalResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAccountZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAccountZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountZonalResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAccountZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAccountZonalResponse) GetBody() *ResetAccountZonalResponseBody {
	return s.Body
}

func (s *ResetAccountZonalResponse) SetHeaders(v map[string]*string) *ResetAccountZonalResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountZonalResponse) SetStatusCode(v int32) *ResetAccountZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountZonalResponse) SetBody(v *ResetAccountZonalResponseBody) *ResetAccountZonalResponse {
	s.Body = v
	return s
}

func (s *ResetAccountZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
