// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMultiDimTableRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMultiDimTableRecordsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMultiDimTableRecordsResponseBody) *UpdateMultiDimTableRecordsResponse
	GetBody() *UpdateMultiDimTableRecordsResponseBody
}

type UpdateMultiDimTableRecordsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultiDimTableRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultiDimTableRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMultiDimTableRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMultiDimTableRecordsResponse) GetBody() *UpdateMultiDimTableRecordsResponseBody {
	return s.Body
}

func (s *UpdateMultiDimTableRecordsResponse) SetHeaders(v map[string]*string) *UpdateMultiDimTableRecordsResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultiDimTableRecordsResponse) SetStatusCode(v int32) *UpdateMultiDimTableRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultiDimTableRecordsResponse) SetBody(v *UpdateMultiDimTableRecordsResponseBody) *UpdateMultiDimTableRecordsResponse {
	s.Body = v
	return s
}

func (s *UpdateMultiDimTableRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
