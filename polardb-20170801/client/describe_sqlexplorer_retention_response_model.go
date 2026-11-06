// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLExplorerRetentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLExplorerRetentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLExplorerRetentionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLExplorerRetentionResponseBody) *DescribeSQLExplorerRetentionResponse
	GetBody() *DescribeSQLExplorerRetentionResponseBody
}

type DescribeSQLExplorerRetentionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLExplorerRetentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLExplorerRetentionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLExplorerRetentionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLExplorerRetentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLExplorerRetentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLExplorerRetentionResponse) GetBody() *DescribeSQLExplorerRetentionResponseBody {
	return s.Body
}

func (s *DescribeSQLExplorerRetentionResponse) SetHeaders(v map[string]*string) *DescribeSQLExplorerRetentionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLExplorerRetentionResponse) SetStatusCode(v int32) *DescribeSQLExplorerRetentionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLExplorerRetentionResponse) SetBody(v *DescribeSQLExplorerRetentionResponseBody) *DescribeSQLExplorerRetentionResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLExplorerRetentionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
