// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLServerUpgradeVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLServerUpgradeVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLServerUpgradeVersionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLServerUpgradeVersionsResponseBody) *DescribeSQLServerUpgradeVersionsResponse
	GetBody() *DescribeSQLServerUpgradeVersionsResponseBody
}

type DescribeSQLServerUpgradeVersionsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLServerUpgradeVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLServerUpgradeVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLServerUpgradeVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLServerUpgradeVersionsResponse) GetBody() *DescribeSQLServerUpgradeVersionsResponseBody {
	return s.Body
}

func (s *DescribeSQLServerUpgradeVersionsResponse) SetHeaders(v map[string]*string) *DescribeSQLServerUpgradeVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponse) SetStatusCode(v int32) *DescribeSQLServerUpgradeVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponse) SetBody(v *DescribeSQLServerUpgradeVersionsResponseBody) *DescribeSQLServerUpgradeVersionsResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
