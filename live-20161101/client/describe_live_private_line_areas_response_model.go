// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePrivateLineAreasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLivePrivateLineAreasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLivePrivateLineAreasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLivePrivateLineAreasResponseBody) *DescribeLivePrivateLineAreasResponse
	GetBody() *DescribeLivePrivateLineAreasResponseBody
}

type DescribeLivePrivateLineAreasResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLivePrivateLineAreasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLivePrivateLineAreasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAreasResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAreasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLivePrivateLineAreasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLivePrivateLineAreasResponse) GetBody() *DescribeLivePrivateLineAreasResponseBody {
	return s.Body
}

func (s *DescribeLivePrivateLineAreasResponse) SetHeaders(v map[string]*string) *DescribeLivePrivateLineAreasResponse {
	s.Headers = v
	return s
}

func (s *DescribeLivePrivateLineAreasResponse) SetStatusCode(v int32) *DescribeLivePrivateLineAreasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLivePrivateLineAreasResponse) SetBody(v *DescribeLivePrivateLineAreasResponseBody) *DescribeLivePrivateLineAreasResponse {
	s.Body = v
	return s
}

func (s *DescribeLivePrivateLineAreasResponse) Validate() error {
	return dara.Validate(s)
}
