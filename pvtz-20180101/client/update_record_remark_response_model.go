// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecordRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecordRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecordRemarkResponseBody) *UpdateRecordRemarkResponse
	GetBody() *UpdateRecordRemarkResponseBody
}

type UpdateRecordRemarkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecordRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecordRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecordRemarkResponse) GetBody() *UpdateRecordRemarkResponseBody {
	return s.Body
}

func (s *UpdateRecordRemarkResponse) SetHeaders(v map[string]*string) *UpdateRecordRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordRemarkResponse) SetStatusCode(v int32) *UpdateRecordRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordRemarkResponse) SetBody(v *UpdateRecordRemarkResponseBody) *UpdateRecordRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateRecordRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
