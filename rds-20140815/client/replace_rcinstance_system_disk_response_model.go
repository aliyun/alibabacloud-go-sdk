// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceRCInstanceSystemDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceRCInstanceSystemDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceRCInstanceSystemDiskResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceRCInstanceSystemDiskResponseBody) *ReplaceRCInstanceSystemDiskResponse
	GetBody() *ReplaceRCInstanceSystemDiskResponseBody
}

type ReplaceRCInstanceSystemDiskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceRCInstanceSystemDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceRCInstanceSystemDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceRCInstanceSystemDiskResponse) GoString() string {
	return s.String()
}

func (s *ReplaceRCInstanceSystemDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceRCInstanceSystemDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceRCInstanceSystemDiskResponse) GetBody() *ReplaceRCInstanceSystemDiskResponseBody {
	return s.Body
}

func (s *ReplaceRCInstanceSystemDiskResponse) SetHeaders(v map[string]*string) *ReplaceRCInstanceSystemDiskResponse {
	s.Headers = v
	return s
}

func (s *ReplaceRCInstanceSystemDiskResponse) SetStatusCode(v int32) *ReplaceRCInstanceSystemDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceRCInstanceSystemDiskResponse) SetBody(v *ReplaceRCInstanceSystemDiskResponseBody) *ReplaceRCInstanceSystemDiskResponse {
	s.Body = v
	return s
}

func (s *ReplaceRCInstanceSystemDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
