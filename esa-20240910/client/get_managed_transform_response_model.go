// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedTransformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetManagedTransformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetManagedTransformResponse
	GetStatusCode() *int32
	SetBody(v *GetManagedTransformResponseBody) *GetManagedTransformResponse
	GetBody() *GetManagedTransformResponseBody
}

type GetManagedTransformResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManagedTransformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManagedTransformResponse) String() string {
	return dara.Prettify(s)
}

func (s GetManagedTransformResponse) GoString() string {
	return s.String()
}

func (s *GetManagedTransformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetManagedTransformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetManagedTransformResponse) GetBody() *GetManagedTransformResponseBody {
	return s.Body
}

func (s *GetManagedTransformResponse) SetHeaders(v map[string]*string) *GetManagedTransformResponse {
	s.Headers = v
	return s
}

func (s *GetManagedTransformResponse) SetStatusCode(v int32) *GetManagedTransformResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagedTransformResponse) SetBody(v *GetManagedTransformResponseBody) *GetManagedTransformResponse {
	s.Body = v
	return s
}

func (s *GetManagedTransformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
