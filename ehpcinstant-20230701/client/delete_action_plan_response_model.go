// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteActionPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteActionPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteActionPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteActionPlanResponseBody) *DeleteActionPlanResponse
	GetBody() *DeleteActionPlanResponseBody
}

type DeleteActionPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteActionPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteActionPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteActionPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteActionPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteActionPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteActionPlanResponse) GetBody() *DeleteActionPlanResponseBody {
	return s.Body
}

func (s *DeleteActionPlanResponse) SetHeaders(v map[string]*string) *DeleteActionPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteActionPlanResponse) SetStatusCode(v int32) *DeleteActionPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteActionPlanResponse) SetBody(v *DeleteActionPlanResponseBody) *DeleteActionPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteActionPlanResponse) Validate() error {
	return dara.Validate(s)
}
