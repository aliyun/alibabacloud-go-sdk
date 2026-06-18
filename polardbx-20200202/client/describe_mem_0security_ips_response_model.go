// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMem0SecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMem0SecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMem0SecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMem0SecurityIpsResponseBody) *DescribeMem0SecurityIpsResponse
	GetBody() *DescribeMem0SecurityIpsResponseBody
}

type DescribeMem0SecurityIpsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMem0SecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMem0SecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0SecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMem0SecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMem0SecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMem0SecurityIpsResponse) GetBody() *DescribeMem0SecurityIpsResponseBody {
	return s.Body
}

func (s *DescribeMem0SecurityIpsResponse) SetHeaders(v map[string]*string) *DescribeMem0SecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMem0SecurityIpsResponse) SetStatusCode(v int32) *DescribeMem0SecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMem0SecurityIpsResponse) SetBody(v *DescribeMem0SecurityIpsResponseBody) *DescribeMem0SecurityIpsResponse {
	s.Body = v
	return s
}

func (s *DescribeMem0SecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
