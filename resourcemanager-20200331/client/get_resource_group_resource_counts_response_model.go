// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupResourceCountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupResourceCountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupResourceCountsResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupResourceCountsResponseBody) *GetResourceGroupResourceCountsResponse
	GetBody() *GetResourceGroupResourceCountsResponseBody
}

type GetResourceGroupResourceCountsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupResourceCountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupResourceCountsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResourceCountsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResourceCountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupResourceCountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupResourceCountsResponse) GetBody() *GetResourceGroupResourceCountsResponseBody {
	return s.Body
}

func (s *GetResourceGroupResourceCountsResponse) SetHeaders(v map[string]*string) *GetResourceGroupResourceCountsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupResourceCountsResponse) SetStatusCode(v int32) *GetResourceGroupResourceCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupResourceCountsResponse) SetBody(v *GetResourceGroupResourceCountsResponseBody) *GetResourceGroupResourceCountsResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupResourceCountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
