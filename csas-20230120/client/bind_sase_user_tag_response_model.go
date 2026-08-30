// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSaseUserTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindSaseUserTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindSaseUserTagResponse
	GetStatusCode() *int32
	SetBody(v *BindSaseUserTagResponseBody) *BindSaseUserTagResponse
	GetBody() *BindSaseUserTagResponseBody
}

type BindSaseUserTagResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindSaseUserTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindSaseUserTagResponse) String() string {
	return dara.Prettify(s)
}

func (s BindSaseUserTagResponse) GoString() string {
	return s.String()
}

func (s *BindSaseUserTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindSaseUserTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindSaseUserTagResponse) GetBody() *BindSaseUserTagResponseBody {
	return s.Body
}

func (s *BindSaseUserTagResponse) SetHeaders(v map[string]*string) *BindSaseUserTagResponse {
	s.Headers = v
	return s
}

func (s *BindSaseUserTagResponse) SetStatusCode(v int32) *BindSaseUserTagResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSaseUserTagResponse) SetBody(v *BindSaseUserTagResponseBody) *BindSaseUserTagResponse {
	s.Body = v
	return s
}

func (s *BindSaseUserTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
