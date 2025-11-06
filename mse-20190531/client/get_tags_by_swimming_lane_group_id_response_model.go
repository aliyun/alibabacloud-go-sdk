// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagsBySwimmingLaneGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTagsBySwimmingLaneGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTagsBySwimmingLaneGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *GetTagsBySwimmingLaneGroupIdResponseBody) *GetTagsBySwimmingLaneGroupIdResponse
	GetBody() *GetTagsBySwimmingLaneGroupIdResponseBody
}

type GetTagsBySwimmingLaneGroupIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTagsBySwimmingLaneGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTagsBySwimmingLaneGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTagsBySwimmingLaneGroupIdResponse) GoString() string {
	return s.String()
}

func (s *GetTagsBySwimmingLaneGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTagsBySwimmingLaneGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTagsBySwimmingLaneGroupIdResponse) GetBody() *GetTagsBySwimmingLaneGroupIdResponseBody {
	return s.Body
}

func (s *GetTagsBySwimmingLaneGroupIdResponse) SetHeaders(v map[string]*string) *GetTagsBySwimmingLaneGroupIdResponse {
	s.Headers = v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdResponse) SetStatusCode(v int32) *GetTagsBySwimmingLaneGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdResponse) SetBody(v *GetTagsBySwimmingLaneGroupIdResponseBody) *GetTagsBySwimmingLaneGroupIdResponse {
	s.Body = v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
