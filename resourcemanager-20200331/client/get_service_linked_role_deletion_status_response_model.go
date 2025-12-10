// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleDeletionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceLinkedRoleDeletionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceLinkedRoleDeletionStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceLinkedRoleDeletionStatusResponseBody) *GetServiceLinkedRoleDeletionStatusResponse
	GetBody() *GetServiceLinkedRoleDeletionStatusResponseBody
}

type GetServiceLinkedRoleDeletionStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceLinkedRoleDeletionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceLinkedRoleDeletionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleDeletionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) GetBody() *GetServiceLinkedRoleDeletionStatusResponseBody {
	return s.Body
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetHeaders(v map[string]*string) *GetServiceLinkedRoleDeletionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetStatusCode(v int32) *GetServiceLinkedRoleDeletionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) SetBody(v *GetServiceLinkedRoleDeletionStatusResponseBody) *GetServiceLinkedRoleDeletionStatusResponse {
	s.Body = v
	return s
}

func (s *GetServiceLinkedRoleDeletionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
