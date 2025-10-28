// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertSwimmingLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertSwimmingLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertSwimmingLaneResponse
	GetStatusCode() *int32
	SetBody(v *InsertSwimmingLaneResponseBody) *InsertSwimmingLaneResponse
	GetBody() *InsertSwimmingLaneResponseBody
}

type InsertSwimmingLaneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertSwimmingLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertSwimmingLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneResponse) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertSwimmingLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertSwimmingLaneResponse) GetBody() *InsertSwimmingLaneResponseBody {
	return s.Body
}

func (s *InsertSwimmingLaneResponse) SetHeaders(v map[string]*string) *InsertSwimmingLaneResponse {
	s.Headers = v
	return s
}

func (s *InsertSwimmingLaneResponse) SetStatusCode(v int32) *InsertSwimmingLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertSwimmingLaneResponse) SetBody(v *InsertSwimmingLaneResponseBody) *InsertSwimmingLaneResponse {
	s.Body = v
	return s
}

func (s *InsertSwimmingLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
