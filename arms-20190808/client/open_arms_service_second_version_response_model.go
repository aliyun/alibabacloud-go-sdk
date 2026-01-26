// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsServiceSecondVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenArmsServiceSecondVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenArmsServiceSecondVersionResponse
	GetStatusCode() *int32
	SetBody(v *OpenArmsServiceSecondVersionResponseBody) *OpenArmsServiceSecondVersionResponse
	GetBody() *OpenArmsServiceSecondVersionResponseBody
}

type OpenArmsServiceSecondVersionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenArmsServiceSecondVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenArmsServiceSecondVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsServiceSecondVersionResponse) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceSecondVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenArmsServiceSecondVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenArmsServiceSecondVersionResponse) GetBody() *OpenArmsServiceSecondVersionResponseBody {
	return s.Body
}

func (s *OpenArmsServiceSecondVersionResponse) SetHeaders(v map[string]*string) *OpenArmsServiceSecondVersionResponse {
	s.Headers = v
	return s
}

func (s *OpenArmsServiceSecondVersionResponse) SetStatusCode(v int32) *OpenArmsServiceSecondVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenArmsServiceSecondVersionResponse) SetBody(v *OpenArmsServiceSecondVersionResponseBody) *OpenArmsServiceSecondVersionResponse {
	s.Body = v
	return s
}

func (s *OpenArmsServiceSecondVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
