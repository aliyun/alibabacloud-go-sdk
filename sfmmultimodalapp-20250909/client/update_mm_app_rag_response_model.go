// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmAppRagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmAppRagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmAppRagResponseBody) *UpdateMmAppRagResponse
	GetBody() *UpdateMmAppRagResponseBody
}

type UpdateMmAppRagResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmAppRagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmAppRagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmAppRagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmAppRagResponse) GetBody() *UpdateMmAppRagResponseBody {
	return s.Body
}

func (s *UpdateMmAppRagResponse) SetHeaders(v map[string]*string) *UpdateMmAppRagResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmAppRagResponse) SetStatusCode(v int32) *UpdateMmAppRagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmAppRagResponse) SetBody(v *UpdateMmAppRagResponseBody) *UpdateMmAppRagResponse {
	s.Body = v
	return s
}

func (s *UpdateMmAppRagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
