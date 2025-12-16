// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPassThroughAuthInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPassThroughAuthInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPassThroughAuthInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetPassThroughAuthInfoResponseBody) *GetPassThroughAuthInfoResponse
	GetBody() *GetPassThroughAuthInfoResponseBody
}

type GetPassThroughAuthInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPassThroughAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPassThroughAuthInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPassThroughAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPassThroughAuthInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPassThroughAuthInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPassThroughAuthInfoResponse) GetBody() *GetPassThroughAuthInfoResponseBody {
	return s.Body
}

func (s *GetPassThroughAuthInfoResponse) SetHeaders(v map[string]*string) *GetPassThroughAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPassThroughAuthInfoResponse) SetStatusCode(v int32) *GetPassThroughAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPassThroughAuthInfoResponse) SetBody(v *GetPassThroughAuthInfoResponseBody) *GetPassThroughAuthInfoResponse {
	s.Body = v
	return s
}

func (s *GetPassThroughAuthInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
