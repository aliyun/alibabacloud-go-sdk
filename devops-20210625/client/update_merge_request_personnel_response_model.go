// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMergeRequestPersonnelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMergeRequestPersonnelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMergeRequestPersonnelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMergeRequestPersonnelResponseBody) *UpdateMergeRequestPersonnelResponse
	GetBody() *UpdateMergeRequestPersonnelResponseBody
}

type UpdateMergeRequestPersonnelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMergeRequestPersonnelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMergeRequestPersonnelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMergeRequestPersonnelResponse) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestPersonnelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMergeRequestPersonnelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMergeRequestPersonnelResponse) GetBody() *UpdateMergeRequestPersonnelResponseBody {
	return s.Body
}

func (s *UpdateMergeRequestPersonnelResponse) SetHeaders(v map[string]*string) *UpdateMergeRequestPersonnelResponse {
	s.Headers = v
	return s
}

func (s *UpdateMergeRequestPersonnelResponse) SetStatusCode(v int32) *UpdateMergeRequestPersonnelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMergeRequestPersonnelResponse) SetBody(v *UpdateMergeRequestPersonnelResponseBody) *UpdateMergeRequestPersonnelResponse {
	s.Body = v
	return s
}

func (s *UpdateMergeRequestPersonnelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
