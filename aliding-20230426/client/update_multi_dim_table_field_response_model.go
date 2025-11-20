// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMultiDimTableFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMultiDimTableFieldResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMultiDimTableFieldResponseBody) *UpdateMultiDimTableFieldResponse
	GetBody() *UpdateMultiDimTableFieldResponseBody
}

type UpdateMultiDimTableFieldResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMultiDimTableFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMultiDimTableFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableFieldResponse) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMultiDimTableFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMultiDimTableFieldResponse) GetBody() *UpdateMultiDimTableFieldResponseBody {
	return s.Body
}

func (s *UpdateMultiDimTableFieldResponse) SetHeaders(v map[string]*string) *UpdateMultiDimTableFieldResponse {
	s.Headers = v
	return s
}

func (s *UpdateMultiDimTableFieldResponse) SetStatusCode(v int32) *UpdateMultiDimTableFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMultiDimTableFieldResponse) SetBody(v *UpdateMultiDimTableFieldResponseBody) *UpdateMultiDimTableFieldResponse {
	s.Body = v
	return s
}

func (s *UpdateMultiDimTableFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
