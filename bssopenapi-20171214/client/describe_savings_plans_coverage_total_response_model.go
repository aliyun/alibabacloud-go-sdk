// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSavingsPlansCoverageTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSavingsPlansCoverageTotalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSavingsPlansCoverageTotalResponseBody) *DescribeSavingsPlansCoverageTotalResponse
	GetBody() *DescribeSavingsPlansCoverageTotalResponseBody
}

type DescribeSavingsPlansCoverageTotalResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSavingsPlansCoverageTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSavingsPlansCoverageTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSavingsPlansCoverageTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSavingsPlansCoverageTotalResponse) GetBody() *DescribeSavingsPlansCoverageTotalResponseBody {
	return s.Body
}

func (s *DescribeSavingsPlansCoverageTotalResponse) SetHeaders(v map[string]*string) *DescribeSavingsPlansCoverageTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponse) SetStatusCode(v int32) *DescribeSavingsPlansCoverageTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponse) SetBody(v *DescribeSavingsPlansCoverageTotalResponseBody) *DescribeSavingsPlansCoverageTotalResponse {
	s.Body = v
	return s
}

func (s *DescribeSavingsPlansCoverageTotalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
