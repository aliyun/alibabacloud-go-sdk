// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDumpMetaListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DumpMetaListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DumpMetaListResponse
	GetStatusCode() *int32
	SetBody(v *DumpMetaListResponseBody) *DumpMetaListResponse
	GetBody() *DumpMetaListResponseBody
}

type DumpMetaListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DumpMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DumpMetaListResponse) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaListResponse) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DumpMetaListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DumpMetaListResponse) GetBody() *DumpMetaListResponseBody {
	return s.Body
}

func (s *DumpMetaListResponse) SetHeaders(v map[string]*string) *DumpMetaListResponse {
	s.Headers = v
	return s
}

func (s *DumpMetaListResponse) SetStatusCode(v int32) *DumpMetaListResponse {
	s.StatusCode = &v
	return s
}

func (s *DumpMetaListResponse) SetBody(v *DumpMetaListResponseBody) *DumpMetaListResponse {
	s.Body = v
	return s
}

func (s *DumpMetaListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
