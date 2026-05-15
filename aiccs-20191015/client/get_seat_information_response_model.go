// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSeatInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSeatInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSeatInformationResponse
	GetStatusCode() *int32
	SetBody(v *GetSeatInformationResponseBody) *GetSeatInformationResponse
	GetBody() *GetSeatInformationResponseBody
}

type GetSeatInformationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSeatInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSeatInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSeatInformationResponse) GoString() string {
	return s.String()
}

func (s *GetSeatInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSeatInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSeatInformationResponse) GetBody() *GetSeatInformationResponseBody {
	return s.Body
}

func (s *GetSeatInformationResponse) SetHeaders(v map[string]*string) *GetSeatInformationResponse {
	s.Headers = v
	return s
}

func (s *GetSeatInformationResponse) SetStatusCode(v int32) *GetSeatInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSeatInformationResponse) SetBody(v *GetSeatInformationResponseBody) *GetSeatInformationResponse {
	s.Body = v
	return s
}

func (s *GetSeatInformationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
