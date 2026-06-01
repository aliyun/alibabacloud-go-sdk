// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRayHistoryServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRayHistoryServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRayHistoryServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateRayHistoryServerResponseBody) *CreateRayHistoryServerResponse
	GetBody() *CreateRayHistoryServerResponseBody
}

type CreateRayHistoryServerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRayHistoryServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRayHistoryServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRayHistoryServerResponse) GoString() string {
	return s.String()
}

func (s *CreateRayHistoryServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRayHistoryServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRayHistoryServerResponse) GetBody() *CreateRayHistoryServerResponseBody {
	return s.Body
}

func (s *CreateRayHistoryServerResponse) SetHeaders(v map[string]*string) *CreateRayHistoryServerResponse {
	s.Headers = v
	return s
}

func (s *CreateRayHistoryServerResponse) SetStatusCode(v int32) *CreateRayHistoryServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRayHistoryServerResponse) SetBody(v *CreateRayHistoryServerResponseBody) *CreateRayHistoryServerResponse {
	s.Body = v
	return s
}

func (s *CreateRayHistoryServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
