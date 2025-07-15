// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageGroupResponseBody) *GetMessageGroupResponse
	GetBody() *GetMessageGroupResponseBody
}

type GetMessageGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *GetMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageGroupResponse) GetBody() *GetMessageGroupResponseBody {
	return s.Body
}

func (s *GetMessageGroupResponse) SetHeaders(v map[string]*string) *GetMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *GetMessageGroupResponse) SetStatusCode(v int32) *GetMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageGroupResponse) SetBody(v *GetMessageGroupResponseBody) *GetMessageGroupResponse {
	s.Body = v
	return s
}

func (s *GetMessageGroupResponse) Validate() error {
	return dara.Validate(s)
}
