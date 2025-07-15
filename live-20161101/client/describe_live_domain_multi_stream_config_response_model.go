// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMultiStreamConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainMultiStreamConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainMultiStreamConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainMultiStreamConfigResponseBody) *DescribeLiveDomainMultiStreamConfigResponse
	GetBody() *DescribeLiveDomainMultiStreamConfigResponseBody
}

type DescribeLiveDomainMultiStreamConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainMultiStreamConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainMultiStreamConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMultiStreamConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMultiStreamConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainMultiStreamConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainMultiStreamConfigResponse) GetBody() *DescribeLiveDomainMultiStreamConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainMultiStreamConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainMultiStreamConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainMultiStreamConfigResponse) SetStatusCode(v int32) *DescribeLiveDomainMultiStreamConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainMultiStreamConfigResponse) SetBody(v *DescribeLiveDomainMultiStreamConfigResponseBody) *DescribeLiveDomainMultiStreamConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainMultiStreamConfigResponse) Validate() error {
	return dara.Validate(s)
}
