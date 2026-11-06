// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLExplorerPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLExplorerPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLExplorerPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLExplorerPolicyResponseBody) *DescribeSQLExplorerPolicyResponse
	GetBody() *DescribeSQLExplorerPolicyResponseBody
}

type DescribeSQLExplorerPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLExplorerPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLExplorerPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLExplorerPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLExplorerPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLExplorerPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLExplorerPolicyResponse) GetBody() *DescribeSQLExplorerPolicyResponseBody {
	return s.Body
}

func (s *DescribeSQLExplorerPolicyResponse) SetHeaders(v map[string]*string) *DescribeSQLExplorerPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLExplorerPolicyResponse) SetStatusCode(v int32) *DescribeSQLExplorerPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLExplorerPolicyResponse) SetBody(v *DescribeSQLExplorerPolicyResponseBody) *DescribeSQLExplorerPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLExplorerPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
