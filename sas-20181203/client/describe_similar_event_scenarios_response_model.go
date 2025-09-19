// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimilarEventScenariosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSimilarEventScenariosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSimilarEventScenariosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSimilarEventScenariosResponseBody) *DescribeSimilarEventScenariosResponse
	GetBody() *DescribeSimilarEventScenariosResponseBody
}

type DescribeSimilarEventScenariosResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSimilarEventScenariosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSimilarEventScenariosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarEventScenariosResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSimilarEventScenariosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSimilarEventScenariosResponse) GetBody() *DescribeSimilarEventScenariosResponseBody {
	return s.Body
}

func (s *DescribeSimilarEventScenariosResponse) SetHeaders(v map[string]*string) *DescribeSimilarEventScenariosResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimilarEventScenariosResponse) SetStatusCode(v int32) *DescribeSimilarEventScenariosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSimilarEventScenariosResponse) SetBody(v *DescribeSimilarEventScenariosResponseBody) *DescribeSimilarEventScenariosResponse {
	s.Body = v
	return s
}

func (s *DescribeSimilarEventScenariosResponse) Validate() error {
	return dara.Validate(s)
}
