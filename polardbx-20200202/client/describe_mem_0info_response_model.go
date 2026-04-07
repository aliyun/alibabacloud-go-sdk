// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMem0InfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMem0InfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMem0InfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMem0InfoResponseBody) *DescribeMem0InfoResponse
	GetBody() *DescribeMem0InfoResponseBody
}

type DescribeMem0InfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMem0InfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMem0InfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0InfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMem0InfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMem0InfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMem0InfoResponse) GetBody() *DescribeMem0InfoResponseBody {
	return s.Body
}

func (s *DescribeMem0InfoResponse) SetHeaders(v map[string]*string) *DescribeMem0InfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMem0InfoResponse) SetStatusCode(v int32) *DescribeMem0InfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMem0InfoResponse) SetBody(v *DescribeMem0InfoResponseBody) *DescribeMem0InfoResponse {
	s.Body = v
	return s
}

func (s *DescribeMem0InfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
