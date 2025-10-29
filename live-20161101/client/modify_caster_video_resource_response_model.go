// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterVideoResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCasterVideoResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCasterVideoResourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCasterVideoResourceResponseBody) *ModifyCasterVideoResourceResponse
	GetBody() *ModifyCasterVideoResourceResponseBody
}

type ModifyCasterVideoResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCasterVideoResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCasterVideoResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterVideoResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterVideoResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCasterVideoResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCasterVideoResourceResponse) GetBody() *ModifyCasterVideoResourceResponseBody {
	return s.Body
}

func (s *ModifyCasterVideoResourceResponse) SetHeaders(v map[string]*string) *ModifyCasterVideoResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyCasterVideoResourceResponse) SetStatusCode(v int32) *ModifyCasterVideoResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCasterVideoResourceResponse) SetBody(v *ModifyCasterVideoResourceResponseBody) *ModifyCasterVideoResourceResponse {
	s.Body = v
	return s
}

func (s *ModifyCasterVideoResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
