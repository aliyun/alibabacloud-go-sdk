// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportInterveneFileAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportInterveneFileAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportInterveneFileAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ImportInterveneFileAsyncResponseBody) *ImportInterveneFileAsyncResponse
	GetBody() *ImportInterveneFileAsyncResponseBody
}

type ImportInterveneFileAsyncResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportInterveneFileAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportInterveneFileAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportInterveneFileAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportInterveneFileAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportInterveneFileAsyncResponse) GetBody() *ImportInterveneFileAsyncResponseBody {
	return s.Body
}

func (s *ImportInterveneFileAsyncResponse) SetHeaders(v map[string]*string) *ImportInterveneFileAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImportInterveneFileAsyncResponse) SetStatusCode(v int32) *ImportInterveneFileAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportInterveneFileAsyncResponse) SetBody(v *ImportInterveneFileAsyncResponseBody) *ImportInterveneFileAsyncResponse {
	s.Body = v
	return s
}

func (s *ImportInterveneFileAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
