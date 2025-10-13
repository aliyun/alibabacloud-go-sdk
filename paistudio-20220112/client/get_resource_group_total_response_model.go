// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupTotalResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupTotalResponseBody) *GetResourceGroupTotalResponse
	GetBody() *GetResourceGroupTotalResponseBody
}

type GetResourceGroupTotalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupTotalResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupTotalResponse) GetBody() *GetResourceGroupTotalResponseBody {
	return s.Body
}

func (s *GetResourceGroupTotalResponse) SetHeaders(v map[string]*string) *GetResourceGroupTotalResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupTotalResponse) SetStatusCode(v int32) *GetResourceGroupTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupTotalResponse) SetBody(v *GetResourceGroupTotalResponseBody) *GetResourceGroupTotalResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupTotalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
