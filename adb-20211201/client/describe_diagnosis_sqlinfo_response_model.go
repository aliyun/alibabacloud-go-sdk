// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSQLInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosisSQLInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosisSQLInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosisSQLInfoResponseBody) *DescribeDiagnosisSQLInfoResponse
	GetBody() *DescribeDiagnosisSQLInfoResponseBody
}

type DescribeDiagnosisSQLInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisSQLInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisSQLInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosisSQLInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosisSQLInfoResponse) GetBody() *DescribeDiagnosisSQLInfoResponseBody {
	return s.Body
}

func (s *DescribeDiagnosisSQLInfoResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisSQLInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetStatusCode(v int32) *DescribeDiagnosisSQLInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetBody(v *DescribeDiagnosisSQLInfoResponseBody) *DescribeDiagnosisSQLInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) Validate() error {
	return dara.Validate(s)
}
