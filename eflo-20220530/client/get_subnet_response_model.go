// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubnetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubnetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubnetResponse
	GetStatusCode() *int32
	SetBody(v *GetSubnetResponseBody) *GetSubnetResponse
	GetBody() *GetSubnetResponseBody
}

type GetSubnetResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubnetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubnetResponse) GoString() string {
	return s.String()
}

func (s *GetSubnetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubnetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubnetResponse) GetBody() *GetSubnetResponseBody {
	return s.Body
}

func (s *GetSubnetResponse) SetHeaders(v map[string]*string) *GetSubnetResponse {
	s.Headers = v
	return s
}

func (s *GetSubnetResponse) SetStatusCode(v int32) *GetSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubnetResponse) SetBody(v *GetSubnetResponseBody) *GetSubnetResponse {
	s.Body = v
	return s
}

func (s *GetSubnetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
