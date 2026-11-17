// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetModelGroupResponseBody) *GetModelGroupResponse
	GetBody() *GetModelGroupResponseBody
}

type GetModelGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelGroupResponse) GoString() string {
	return s.String()
}

func (s *GetModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelGroupResponse) GetBody() *GetModelGroupResponseBody {
	return s.Body
}

func (s *GetModelGroupResponse) SetHeaders(v map[string]*string) *GetModelGroupResponse {
	s.Headers = v
	return s
}

func (s *GetModelGroupResponse) SetStatusCode(v int32) *GetModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelGroupResponse) SetBody(v *GetModelGroupResponseBody) *GetModelGroupResponse {
	s.Body = v
	return s
}

func (s *GetModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
