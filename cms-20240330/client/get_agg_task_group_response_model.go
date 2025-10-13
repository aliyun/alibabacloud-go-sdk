// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggTaskGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggTaskGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggTaskGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetAggTaskGroupResponseBody) *GetAggTaskGroupResponse
	GetBody() *GetAggTaskGroupResponseBody
}

type GetAggTaskGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggTaskGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *GetAggTaskGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggTaskGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggTaskGroupResponse) GetBody() *GetAggTaskGroupResponseBody {
	return s.Body
}

func (s *GetAggTaskGroupResponse) SetHeaders(v map[string]*string) *GetAggTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *GetAggTaskGroupResponse) SetStatusCode(v int32) *GetAggTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggTaskGroupResponse) SetBody(v *GetAggTaskGroupResponseBody) *GetAggTaskGroupResponse {
	s.Body = v
	return s
}

func (s *GetAggTaskGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
