// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFieldDefByUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFieldDefByUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFieldDefByUuidResponse
	GetStatusCode() *int32
	SetBody(v *GetFieldDefByUuidResponseBody) *GetFieldDefByUuidResponse
	GetBody() *GetFieldDefByUuidResponseBody
}

type GetFieldDefByUuidResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFieldDefByUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFieldDefByUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFieldDefByUuidResponse) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFieldDefByUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFieldDefByUuidResponse) GetBody() *GetFieldDefByUuidResponseBody {
	return s.Body
}

func (s *GetFieldDefByUuidResponse) SetHeaders(v map[string]*string) *GetFieldDefByUuidResponse {
	s.Headers = v
	return s
}

func (s *GetFieldDefByUuidResponse) SetStatusCode(v int32) *GetFieldDefByUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFieldDefByUuidResponse) SetBody(v *GetFieldDefByUuidResponseBody) *GetFieldDefByUuidResponse {
	s.Body = v
	return s
}

func (s *GetFieldDefByUuidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
