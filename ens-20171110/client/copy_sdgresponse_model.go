// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopySDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopySDGResponse
	GetStatusCode() *int32
	SetBody(v *CopySDGResponseBody) *CopySDGResponse
	GetBody() *CopySDGResponseBody
}

type CopySDGResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopySDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopySDGResponse) String() string {
	return dara.Prettify(s)
}

func (s CopySDGResponse) GoString() string {
	return s.String()
}

func (s *CopySDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopySDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopySDGResponse) GetBody() *CopySDGResponseBody {
	return s.Body
}

func (s *CopySDGResponse) SetHeaders(v map[string]*string) *CopySDGResponse {
	s.Headers = v
	return s
}

func (s *CopySDGResponse) SetStatusCode(v int32) *CopySDGResponse {
	s.StatusCode = &v
	return s
}

func (s *CopySDGResponse) SetBody(v *CopySDGResponseBody) *CopySDGResponse {
	s.Body = v
	return s
}

func (s *CopySDGResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
