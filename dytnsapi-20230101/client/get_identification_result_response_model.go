// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentificationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentificationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentificationResultResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentificationResultResponseBody) *GetIdentificationResultResponse
	GetBody() *GetIdentificationResultResponseBody
}

type GetIdentificationResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentificationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentificationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentificationResultResponse) GoString() string {
	return s.String()
}

func (s *GetIdentificationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentificationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentificationResultResponse) GetBody() *GetIdentificationResultResponseBody {
	return s.Body
}

func (s *GetIdentificationResultResponse) SetHeaders(v map[string]*string) *GetIdentificationResultResponse {
	s.Headers = v
	return s
}

func (s *GetIdentificationResultResponse) SetStatusCode(v int32) *GetIdentificationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentificationResultResponse) SetBody(v *GetIdentificationResultResponseBody) *GetIdentificationResultResponse {
	s.Body = v
	return s
}

func (s *GetIdentificationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
