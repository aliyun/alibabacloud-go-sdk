// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAegisForLingjunStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAegisForLingjunStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAegisForLingjunStatusResponse
	GetStatusCode() *int32
	SetBody(v *ListAegisForLingjunStatusResponseBody) *ListAegisForLingjunStatusResponse
	GetBody() *ListAegisForLingjunStatusResponseBody
}

type ListAegisForLingjunStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAegisForLingjunStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAegisForLingjunStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAegisForLingjunStatusResponse) GoString() string {
	return s.String()
}

func (s *ListAegisForLingjunStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAegisForLingjunStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAegisForLingjunStatusResponse) GetBody() *ListAegisForLingjunStatusResponseBody {
	return s.Body
}

func (s *ListAegisForLingjunStatusResponse) SetHeaders(v map[string]*string) *ListAegisForLingjunStatusResponse {
	s.Headers = v
	return s
}

func (s *ListAegisForLingjunStatusResponse) SetStatusCode(v int32) *ListAegisForLingjunStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAegisForLingjunStatusResponse) SetBody(v *ListAegisForLingjunStatusResponseBody) *ListAegisForLingjunStatusResponse {
	s.Body = v
	return s
}

func (s *ListAegisForLingjunStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
