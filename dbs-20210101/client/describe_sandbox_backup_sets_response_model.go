// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxBackupSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSandboxBackupSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSandboxBackupSetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSandboxBackupSetsResponseBody) *DescribeSandboxBackupSetsResponse
	GetBody() *DescribeSandboxBackupSetsResponseBody
}

type DescribeSandboxBackupSetsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSandboxBackupSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSandboxBackupSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxBackupSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSandboxBackupSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSandboxBackupSetsResponse) GetBody() *DescribeSandboxBackupSetsResponseBody {
	return s.Body
}

func (s *DescribeSandboxBackupSetsResponse) SetHeaders(v map[string]*string) *DescribeSandboxBackupSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxBackupSetsResponse) SetStatusCode(v int32) *DescribeSandboxBackupSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponse) SetBody(v *DescribeSandboxBackupSetsResponseBody) *DescribeSandboxBackupSetsResponse {
	s.Body = v
	return s
}

func (s *DescribeSandboxBackupSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
