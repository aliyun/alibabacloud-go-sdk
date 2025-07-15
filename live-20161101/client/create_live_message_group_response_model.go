// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveMessageGroupResponseBody) *CreateLiveMessageGroupResponse
	GetBody() *CreateLiveMessageGroupResponseBody
}

type CreateLiveMessageGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveMessageGroupResponse) GetBody() *CreateLiveMessageGroupResponseBody {
	return s.Body
}

func (s *CreateLiveMessageGroupResponse) SetHeaders(v map[string]*string) *CreateLiveMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveMessageGroupResponse) SetStatusCode(v int32) *CreateLiveMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveMessageGroupResponse) SetBody(v *CreateLiveMessageGroupResponseBody) *CreateLiveMessageGroupResponse {
	s.Body = v
	return s
}

func (s *CreateLiveMessageGroupResponse) Validate() error {
	return dara.Validate(s)
}
