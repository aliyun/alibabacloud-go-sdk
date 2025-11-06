// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllSwimmingLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAllSwimmingLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAllSwimmingLaneResponse
	GetStatusCode() *int32
	SetBody(v *QueryAllSwimmingLaneResponseBody) *QueryAllSwimmingLaneResponse
	GetBody() *QueryAllSwimmingLaneResponseBody
}

type QueryAllSwimmingLaneResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllSwimmingLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAllSwimmingLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneResponse) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAllSwimmingLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAllSwimmingLaneResponse) GetBody() *QueryAllSwimmingLaneResponseBody {
	return s.Body
}

func (s *QueryAllSwimmingLaneResponse) SetHeaders(v map[string]*string) *QueryAllSwimmingLaneResponse {
	s.Headers = v
	return s
}

func (s *QueryAllSwimmingLaneResponse) SetStatusCode(v int32) *QueryAllSwimmingLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllSwimmingLaneResponse) SetBody(v *QueryAllSwimmingLaneResponseBody) *QueryAllSwimmingLaneResponse {
	s.Body = v
	return s
}

func (s *QueryAllSwimmingLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
