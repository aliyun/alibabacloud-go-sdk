// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTagResponse
	GetStatusCode() *int32
	SetBody(v *GetTagResponseBody) *GetTagResponse
	GetBody() *GetTagResponseBody
}

type GetTagResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTagResponse) GoString() string {
	return s.String()
}

func (s *GetTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTagResponse) GetBody() *GetTagResponseBody {
	return s.Body
}

func (s *GetTagResponse) SetHeaders(v map[string]*string) *GetTagResponse {
	s.Headers = v
	return s
}

func (s *GetTagResponse) SetStatusCode(v int32) *GetTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagResponse) SetBody(v *GetTagResponseBody) *GetTagResponse {
	s.Body = v
	return s
}

func (s *GetTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
