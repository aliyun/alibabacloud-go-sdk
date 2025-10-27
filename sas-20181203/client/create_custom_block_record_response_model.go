// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomBlockRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomBlockRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomBlockRecordResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomBlockRecordResponseBody) *CreateCustomBlockRecordResponse
	GetBody() *CreateCustomBlockRecordResponseBody
}

type CreateCustomBlockRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomBlockRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomBlockRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomBlockRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomBlockRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomBlockRecordResponse) GetBody() *CreateCustomBlockRecordResponseBody {
	return s.Body
}

func (s *CreateCustomBlockRecordResponse) SetHeaders(v map[string]*string) *CreateCustomBlockRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomBlockRecordResponse) SetStatusCode(v int32) *CreateCustomBlockRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomBlockRecordResponse) SetBody(v *CreateCustomBlockRecordResponseBody) *CreateCustomBlockRecordResponse {
	s.Body = v
	return s
}

func (s *CreateCustomBlockRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
