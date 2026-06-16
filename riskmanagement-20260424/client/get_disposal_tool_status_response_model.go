// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDisposalToolStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDisposalToolStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDisposalToolStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetDisposalToolStatusResponseBody) *GetDisposalToolStatusResponse
	GetBody() *GetDisposalToolStatusResponseBody
}

type GetDisposalToolStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDisposalToolStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDisposalToolStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDisposalToolStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDisposalToolStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDisposalToolStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDisposalToolStatusResponse) GetBody() *GetDisposalToolStatusResponseBody {
	return s.Body
}

func (s *GetDisposalToolStatusResponse) SetHeaders(v map[string]*string) *GetDisposalToolStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDisposalToolStatusResponse) SetStatusCode(v int32) *GetDisposalToolStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDisposalToolStatusResponse) SetBody(v *GetDisposalToolStatusResponseBody) *GetDisposalToolStatusResponse {
	s.Body = v
	return s
}

func (s *GetDisposalToolStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
