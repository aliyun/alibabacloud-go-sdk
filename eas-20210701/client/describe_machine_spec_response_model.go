// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMachineSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMachineSpecResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMachineSpecResponseBody) *DescribeMachineSpecResponse
	GetBody() *DescribeMachineSpecResponseBody
}

type DescribeMachineSpecResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMachineSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMachineSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeMachineSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMachineSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMachineSpecResponse) GetBody() *DescribeMachineSpecResponseBody {
	return s.Body
}

func (s *DescribeMachineSpecResponse) SetHeaders(v map[string]*string) *DescribeMachineSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeMachineSpecResponse) SetStatusCode(v int32) *DescribeMachineSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMachineSpecResponse) SetBody(v *DescribeMachineSpecResponseBody) *DescribeMachineSpecResponse {
	s.Body = v
	return s
}

func (s *DescribeMachineSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
