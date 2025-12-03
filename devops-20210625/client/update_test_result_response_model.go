// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTestResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTestResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTestResultResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTestResultResponseBody) *UpdateTestResultResponse
	GetBody() *UpdateTestResultResponseBody
}

type UpdateTestResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTestResultResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestResultResponse) GoString() string {
	return s.String()
}

func (s *UpdateTestResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTestResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTestResultResponse) GetBody() *UpdateTestResultResponseBody {
	return s.Body
}

func (s *UpdateTestResultResponse) SetHeaders(v map[string]*string) *UpdateTestResultResponse {
	s.Headers = v
	return s
}

func (s *UpdateTestResultResponse) SetStatusCode(v int32) *UpdateTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTestResultResponse) SetBody(v *UpdateTestResultResponseBody) *UpdateTestResultResponse {
	s.Body = v
	return s
}

func (s *UpdateTestResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
