// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApsJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApsJobResponseBody) *DeleteApsJobResponse
	GetBody() *DeleteApsJobResponseBody
}

type DeleteApsJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteApsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApsJobResponse) GetBody() *DeleteApsJobResponseBody {
	return s.Body
}

func (s *DeleteApsJobResponse) SetHeaders(v map[string]*string) *DeleteApsJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteApsJobResponse) SetStatusCode(v int32) *DeleteApsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApsJobResponse) SetBody(v *DeleteApsJobResponseBody) *DeleteApsJobResponse {
	s.Body = v
	return s
}

func (s *DeleteApsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
