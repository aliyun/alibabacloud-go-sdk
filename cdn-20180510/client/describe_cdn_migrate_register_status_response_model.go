// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnMigrateRegisterStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnMigrateRegisterStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnMigrateRegisterStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnMigrateRegisterStatusResponseBody) *DescribeCdnMigrateRegisterStatusResponse
	GetBody() *DescribeCdnMigrateRegisterStatusResponseBody
}

type DescribeCdnMigrateRegisterStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnMigrateRegisterStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnMigrateRegisterStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnMigrateRegisterStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnMigrateRegisterStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnMigrateRegisterStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnMigrateRegisterStatusResponse) GetBody() *DescribeCdnMigrateRegisterStatusResponseBody {
	return s.Body
}

func (s *DescribeCdnMigrateRegisterStatusResponse) SetHeaders(v map[string]*string) *DescribeCdnMigrateRegisterStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnMigrateRegisterStatusResponse) SetStatusCode(v int32) *DescribeCdnMigrateRegisterStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnMigrateRegisterStatusResponse) SetBody(v *DescribeCdnMigrateRegisterStatusResponseBody) *DescribeCdnMigrateRegisterStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnMigrateRegisterStatusResponse) Validate() error {
	return dara.Validate(s)
}
