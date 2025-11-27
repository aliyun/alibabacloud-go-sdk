// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustinsResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustinsResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustinsResourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustinsResourceResponseBody) *ModifyCustinsResourceResponse
	GetBody() *ModifyCustinsResourceResponseBody
}

type ModifyCustinsResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustinsResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustinsResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustinsResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustinsResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustinsResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustinsResourceResponse) GetBody() *ModifyCustinsResourceResponseBody {
	return s.Body
}

func (s *ModifyCustinsResourceResponse) SetHeaders(v map[string]*string) *ModifyCustinsResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustinsResourceResponse) SetStatusCode(v int32) *ModifyCustinsResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustinsResourceResponse) SetBody(v *ModifyCustinsResourceResponseBody) *ModifyCustinsResourceResponse {
	s.Body = v
	return s
}

func (s *ModifyCustinsResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
