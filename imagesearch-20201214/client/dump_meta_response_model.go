// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDumpMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DumpMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DumpMetaResponse
	GetStatusCode() *int32
	SetBody(v *DumpMetaResponseBody) *DumpMetaResponse
	GetBody() *DumpMetaResponseBody
}

type DumpMetaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DumpMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DumpMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaResponse) GoString() string {
	return s.String()
}

func (s *DumpMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DumpMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DumpMetaResponse) GetBody() *DumpMetaResponseBody {
	return s.Body
}

func (s *DumpMetaResponse) SetHeaders(v map[string]*string) *DumpMetaResponse {
	s.Headers = v
	return s
}

func (s *DumpMetaResponse) SetStatusCode(v int32) *DumpMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DumpMetaResponse) SetBody(v *DumpMetaResponseBody) *DumpMetaResponse {
	s.Body = v
	return s
}

func (s *DumpMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
