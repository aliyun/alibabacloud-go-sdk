// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileProtectClientEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileProtectClientEventResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileProtectClientEventResponseBody) *UpdateFileProtectClientEventResponse
	GetBody() *UpdateFileProtectClientEventResponseBody
}

type UpdateFileProtectClientEventResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileProtectClientEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileProtectClientEventResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientEventResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileProtectClientEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileProtectClientEventResponse) GetBody() *UpdateFileProtectClientEventResponseBody {
	return s.Body
}

func (s *UpdateFileProtectClientEventResponse) SetHeaders(v map[string]*string) *UpdateFileProtectClientEventResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileProtectClientEventResponse) SetStatusCode(v int32) *UpdateFileProtectClientEventResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileProtectClientEventResponse) SetBody(v *UpdateFileProtectClientEventResponseBody) *UpdateFileProtectClientEventResponse {
	s.Body = v
	return s
}

func (s *UpdateFileProtectClientEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
