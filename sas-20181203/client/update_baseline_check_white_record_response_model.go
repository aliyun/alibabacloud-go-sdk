// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBaselineCheckWhiteRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBaselineCheckWhiteRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBaselineCheckWhiteRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBaselineCheckWhiteRecordResponseBody) *UpdateBaselineCheckWhiteRecordResponse
	GetBody() *UpdateBaselineCheckWhiteRecordResponseBody
}

type UpdateBaselineCheckWhiteRecordResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBaselineCheckWhiteRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBaselineCheckWhiteRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineCheckWhiteRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateBaselineCheckWhiteRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBaselineCheckWhiteRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBaselineCheckWhiteRecordResponse) GetBody() *UpdateBaselineCheckWhiteRecordResponseBody {
	return s.Body
}

func (s *UpdateBaselineCheckWhiteRecordResponse) SetHeaders(v map[string]*string) *UpdateBaselineCheckWhiteRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponse) SetStatusCode(v int32) *UpdateBaselineCheckWhiteRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponse) SetBody(v *UpdateBaselineCheckWhiteRecordResponseBody) *UpdateBaselineCheckWhiteRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateBaselineCheckWhiteRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
