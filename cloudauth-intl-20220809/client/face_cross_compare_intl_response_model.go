// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCrossCompareIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceCrossCompareIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceCrossCompareIntlResponse
	GetStatusCode() *int32
	SetBody(v *FaceCrossCompareIntlResponseBody) *FaceCrossCompareIntlResponse
	GetBody() *FaceCrossCompareIntlResponseBody
}

type FaceCrossCompareIntlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceCrossCompareIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceCrossCompareIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s FaceCrossCompareIntlResponse) GoString() string {
	return s.String()
}

func (s *FaceCrossCompareIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceCrossCompareIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceCrossCompareIntlResponse) GetBody() *FaceCrossCompareIntlResponseBody {
	return s.Body
}

func (s *FaceCrossCompareIntlResponse) SetHeaders(v map[string]*string) *FaceCrossCompareIntlResponse {
	s.Headers = v
	return s
}

func (s *FaceCrossCompareIntlResponse) SetStatusCode(v int32) *FaceCrossCompareIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceCrossCompareIntlResponse) SetBody(v *FaceCrossCompareIntlResponseBody) *FaceCrossCompareIntlResponse {
	s.Body = v
	return s
}

func (s *FaceCrossCompareIntlResponse) Validate() error {
	return dara.Validate(s)
}
