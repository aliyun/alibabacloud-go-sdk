// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateSwimmingLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateSwimmingLaneResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateSwimmingLaneResponseBody) *CreateOrUpdateSwimmingLaneResponse
	GetBody() *CreateOrUpdateSwimmingLaneResponseBody
}

type CreateOrUpdateSwimmingLaneResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateSwimmingLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateSwimmingLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateSwimmingLaneResponse) GetBody() *CreateOrUpdateSwimmingLaneResponseBody {
	return s.Body
}

func (s *CreateOrUpdateSwimmingLaneResponse) SetHeaders(v map[string]*string) *CreateOrUpdateSwimmingLaneResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponse) SetStatusCode(v int32) *CreateOrUpdateSwimmingLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponse) SetBody(v *CreateOrUpdateSwimmingLaneResponseBody) *CreateOrUpdateSwimmingLaneResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
