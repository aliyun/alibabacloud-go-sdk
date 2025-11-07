// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInternetTupleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInternetTupleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInternetTupleResponse
	GetStatusCode() *int32
	SetBody(v *GetInternetTupleResponseBody) *GetInternetTupleResponse
	GetBody() *GetInternetTupleResponseBody
}

type GetInternetTupleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInternetTupleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInternetTupleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInternetTupleResponse) GoString() string {
	return s.String()
}

func (s *GetInternetTupleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInternetTupleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInternetTupleResponse) GetBody() *GetInternetTupleResponseBody {
	return s.Body
}

func (s *GetInternetTupleResponse) SetHeaders(v map[string]*string) *GetInternetTupleResponse {
	s.Headers = v
	return s
}

func (s *GetInternetTupleResponse) SetStatusCode(v int32) *GetInternetTupleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInternetTupleResponse) SetBody(v *GetInternetTupleResponseBody) *GetInternetTupleResponse {
	s.Body = v
	return s
}

func (s *GetInternetTupleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
