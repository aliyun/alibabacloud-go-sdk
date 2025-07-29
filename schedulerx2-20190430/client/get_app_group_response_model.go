// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetAppGroupResponseBody) *GetAppGroupResponse
	GetBody() *GetAppGroupResponseBody
}

type GetAppGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppGroupResponse) GoString() string {
	return s.String()
}

func (s *GetAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppGroupResponse) GetBody() *GetAppGroupResponseBody {
	return s.Body
}

func (s *GetAppGroupResponse) SetHeaders(v map[string]*string) *GetAppGroupResponse {
	s.Headers = v
	return s
}

func (s *GetAppGroupResponse) SetStatusCode(v int32) *GetAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppGroupResponse) SetBody(v *GetAppGroupResponseBody) *GetAppGroupResponse {
	s.Body = v
	return s
}

func (s *GetAppGroupResponse) Validate() error {
	return dara.Validate(s)
}
