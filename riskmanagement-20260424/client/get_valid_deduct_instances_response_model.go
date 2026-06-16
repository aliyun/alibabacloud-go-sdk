// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidDeductInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetValidDeductInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetValidDeductInstancesResponse
	GetStatusCode() *int32
	SetBody(v *GetValidDeductInstancesResponseBody) *GetValidDeductInstancesResponse
	GetBody() *GetValidDeductInstancesResponseBody
}

type GetValidDeductInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetValidDeductInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetValidDeductInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetValidDeductInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetValidDeductInstancesResponse) GetBody() *GetValidDeductInstancesResponseBody {
	return s.Body
}

func (s *GetValidDeductInstancesResponse) SetHeaders(v map[string]*string) *GetValidDeductInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetValidDeductInstancesResponse) SetStatusCode(v int32) *GetValidDeductInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetValidDeductInstancesResponse) SetBody(v *GetValidDeductInstancesResponseBody) *GetValidDeductInstancesResponse {
	s.Body = v
	return s
}

func (s *GetValidDeductInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
