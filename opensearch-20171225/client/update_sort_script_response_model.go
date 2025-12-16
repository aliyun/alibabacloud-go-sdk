// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSortScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSortScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSortScriptResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSortScriptResponseBody) *UpdateSortScriptResponse
	GetBody() *UpdateSortScriptResponseBody
}

type UpdateSortScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSortScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSortScriptResponse) GoString() string {
	return s.String()
}

func (s *UpdateSortScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSortScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSortScriptResponse) GetBody() *UpdateSortScriptResponseBody {
	return s.Body
}

func (s *UpdateSortScriptResponse) SetHeaders(v map[string]*string) *UpdateSortScriptResponse {
	s.Headers = v
	return s
}

func (s *UpdateSortScriptResponse) SetStatusCode(v int32) *UpdateSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSortScriptResponse) SetBody(v *UpdateSortScriptResponseBody) *UpdateSortScriptResponse {
	s.Body = v
	return s
}

func (s *UpdateSortScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
