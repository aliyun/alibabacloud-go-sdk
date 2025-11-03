// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiDestinationResponse
	GetStatusCode() *int32
	SetBody(v *GetApiDestinationResponseBody) *GetApiDestinationResponse
	GetBody() *GetApiDestinationResponseBody
}

type GetApiDestinationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiDestinationResponse) GoString() string {
	return s.String()
}

func (s *GetApiDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiDestinationResponse) GetBody() *GetApiDestinationResponseBody {
	return s.Body
}

func (s *GetApiDestinationResponse) SetHeaders(v map[string]*string) *GetApiDestinationResponse {
	s.Headers = v
	return s
}

func (s *GetApiDestinationResponse) SetStatusCode(v int32) *GetApiDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiDestinationResponse) SetBody(v *GetApiDestinationResponseBody) *GetApiDestinationResponse {
	s.Body = v
	return s
}

func (s *GetApiDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
