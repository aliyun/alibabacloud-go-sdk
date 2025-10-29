// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotIkDictsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHotIkDictsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHotIkDictsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHotIkDictsResponseBody) *UpdateHotIkDictsResponse
	GetBody() *UpdateHotIkDictsResponseBody
}

type UpdateHotIkDictsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotIkDictsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotIkDictsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotIkDictsResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotIkDictsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHotIkDictsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHotIkDictsResponse) GetBody() *UpdateHotIkDictsResponseBody {
	return s.Body
}

func (s *UpdateHotIkDictsResponse) SetHeaders(v map[string]*string) *UpdateHotIkDictsResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotIkDictsResponse) SetStatusCode(v int32) *UpdateHotIkDictsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotIkDictsResponse) SetBody(v *UpdateHotIkDictsResponseBody) *UpdateHotIkDictsResponse {
	s.Body = v
	return s
}

func (s *UpdateHotIkDictsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
