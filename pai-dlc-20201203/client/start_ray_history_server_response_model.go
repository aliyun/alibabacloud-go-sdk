// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRayHistoryServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRayHistoryServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRayHistoryServerResponse
	GetStatusCode() *int32
	SetBody(v *StartRayHistoryServerResponseBody) *StartRayHistoryServerResponse
	GetBody() *StartRayHistoryServerResponseBody
}

type StartRayHistoryServerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRayHistoryServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRayHistoryServerResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRayHistoryServerResponse) GoString() string {
	return s.String()
}

func (s *StartRayHistoryServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRayHistoryServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRayHistoryServerResponse) GetBody() *StartRayHistoryServerResponseBody {
	return s.Body
}

func (s *StartRayHistoryServerResponse) SetHeaders(v map[string]*string) *StartRayHistoryServerResponse {
	s.Headers = v
	return s
}

func (s *StartRayHistoryServerResponse) SetStatusCode(v int32) *StartRayHistoryServerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRayHistoryServerResponse) SetBody(v *StartRayHistoryServerResponseBody) *StartRayHistoryServerResponse {
	s.Body = v
	return s
}

func (s *StartRayHistoryServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
