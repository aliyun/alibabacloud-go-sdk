// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveByPassOrShuntEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveByPassOrShuntEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveByPassOrShuntEventResponse
	GetStatusCode() *int32
	SetBody(v *SaveByPassOrShuntEventResponseBody) *SaveByPassOrShuntEventResponse
	GetBody() *SaveByPassOrShuntEventResponseBody
}

type SaveByPassOrShuntEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveByPassOrShuntEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveByPassOrShuntEventResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveByPassOrShuntEventResponse) GoString() string {
	return s.String()
}

func (s *SaveByPassOrShuntEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveByPassOrShuntEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveByPassOrShuntEventResponse) GetBody() *SaveByPassOrShuntEventResponseBody {
	return s.Body
}

func (s *SaveByPassOrShuntEventResponse) SetHeaders(v map[string]*string) *SaveByPassOrShuntEventResponse {
	s.Headers = v
	return s
}

func (s *SaveByPassOrShuntEventResponse) SetStatusCode(v int32) *SaveByPassOrShuntEventResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveByPassOrShuntEventResponse) SetBody(v *SaveByPassOrShuntEventResponseBody) *SaveByPassOrShuntEventResponse {
	s.Body = v
	return s
}

func (s *SaveByPassOrShuntEventResponse) Validate() error {
	return dara.Validate(s)
}
