// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryByNodeIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyDentryByNodeIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyDentryByNodeIdResponse
	GetStatusCode() *int32
	SetBody(v *CopyDentryByNodeIdResponseBody) *CopyDentryByNodeIdResponse
	GetBody() *CopyDentryByNodeIdResponseBody
}

type CopyDentryByNodeIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDentryByNodeIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDentryByNodeIdResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryByNodeIdResponse) GoString() string {
	return s.String()
}

func (s *CopyDentryByNodeIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyDentryByNodeIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyDentryByNodeIdResponse) GetBody() *CopyDentryByNodeIdResponseBody {
	return s.Body
}

func (s *CopyDentryByNodeIdResponse) SetHeaders(v map[string]*string) *CopyDentryByNodeIdResponse {
	s.Headers = v
	return s
}

func (s *CopyDentryByNodeIdResponse) SetStatusCode(v int32) *CopyDentryByNodeIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDentryByNodeIdResponse) SetBody(v *CopyDentryByNodeIdResponseBody) *CopyDentryByNodeIdResponse {
	s.Body = v
	return s
}

func (s *CopyDentryByNodeIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
