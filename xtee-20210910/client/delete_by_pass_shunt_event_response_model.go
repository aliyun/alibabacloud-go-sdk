// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteByPassShuntEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteByPassShuntEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteByPassShuntEventResponse
	GetStatusCode() *int32
	SetBody(v *DeleteByPassShuntEventResponseBody) *DeleteByPassShuntEventResponse
	GetBody() *DeleteByPassShuntEventResponseBody
}

type DeleteByPassShuntEventResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteByPassShuntEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteByPassShuntEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteByPassShuntEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteByPassShuntEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteByPassShuntEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteByPassShuntEventResponse) GetBody() *DeleteByPassShuntEventResponseBody {
	return s.Body
}

func (s *DeleteByPassShuntEventResponse) SetHeaders(v map[string]*string) *DeleteByPassShuntEventResponse {
	s.Headers = v
	return s
}

func (s *DeleteByPassShuntEventResponse) SetStatusCode(v int32) *DeleteByPassShuntEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteByPassShuntEventResponse) SetBody(v *DeleteByPassShuntEventResponseBody) *DeleteByPassShuntEventResponse {
	s.Body = v
	return s
}

func (s *DeleteByPassShuntEventResponse) Validate() error {
	return dara.Validate(s)
}
