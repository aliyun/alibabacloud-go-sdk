// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *CopyPlaybookResponseBody) *CopyPlaybookResponse
	GetBody() *CopyPlaybookResponseBody
}

type CopyPlaybookResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *CopyPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyPlaybookResponse) GetBody() *CopyPlaybookResponseBody {
	return s.Body
}

func (s *CopyPlaybookResponse) SetHeaders(v map[string]*string) *CopyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *CopyPlaybookResponse) SetStatusCode(v int32) *CopyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyPlaybookResponse) SetBody(v *CopyPlaybookResponseBody) *CopyPlaybookResponse {
	s.Body = v
	return s
}

func (s *CopyPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
