// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupRequestResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupRequestResponseBody) *GetResourceGroupRequestResponse
	GetBody() *GetResourceGroupRequestResponseBody
}

type GetResourceGroupRequestResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupRequestResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupRequestResponse) GetBody() *GetResourceGroupRequestResponseBody {
	return s.Body
}

func (s *GetResourceGroupRequestResponse) SetHeaders(v map[string]*string) *GetResourceGroupRequestResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupRequestResponse) SetStatusCode(v int32) *GetResourceGroupRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupRequestResponse) SetBody(v *GetResourceGroupRequestResponseBody) *GetResourceGroupRequestResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
