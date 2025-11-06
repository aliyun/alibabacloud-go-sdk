// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcdpZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMcdpZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMcdpZoneResponse
	GetStatusCode() *int32
	SetBody(v *QueryMcdpZoneResponseBody) *QueryMcdpZoneResponse
	GetBody() *QueryMcdpZoneResponseBody
}

type QueryMcdpZoneResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMcdpZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMcdpZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMcdpZoneResponse) GoString() string {
	return s.String()
}

func (s *QueryMcdpZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMcdpZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMcdpZoneResponse) GetBody() *QueryMcdpZoneResponseBody {
	return s.Body
}

func (s *QueryMcdpZoneResponse) SetHeaders(v map[string]*string) *QueryMcdpZoneResponse {
	s.Headers = v
	return s
}

func (s *QueryMcdpZoneResponse) SetStatusCode(v int32) *QueryMcdpZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMcdpZoneResponse) SetBody(v *QueryMcdpZoneResponseBody) *QueryMcdpZoneResponse {
	s.Body = v
	return s
}

func (s *QueryMcdpZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
