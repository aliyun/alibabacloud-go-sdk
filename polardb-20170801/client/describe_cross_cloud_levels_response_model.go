// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudLevelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrossCloudLevelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrossCloudLevelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrossCloudLevelsResponseBody) *DescribeCrossCloudLevelsResponse
	GetBody() *DescribeCrossCloudLevelsResponseBody
}

type DescribeCrossCloudLevelsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrossCloudLevelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrossCloudLevelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudLevelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudLevelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrossCloudLevelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrossCloudLevelsResponse) GetBody() *DescribeCrossCloudLevelsResponseBody {
	return s.Body
}

func (s *DescribeCrossCloudLevelsResponse) SetHeaders(v map[string]*string) *DescribeCrossCloudLevelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossCloudLevelsResponse) SetStatusCode(v int32) *DescribeCrossCloudLevelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrossCloudLevelsResponse) SetBody(v *DescribeCrossCloudLevelsResponseBody) *DescribeCrossCloudLevelsResponse {
	s.Body = v
	return s
}

func (s *DescribeCrossCloudLevelsResponse) Validate() error {
	return dara.Validate(s)
}
