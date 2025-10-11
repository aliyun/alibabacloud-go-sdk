// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceCountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiAccountResourceCountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiAccountResourceCountsResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiAccountResourceCountsResponseBody) *GetMultiAccountResourceCountsResponse
	GetBody() *GetMultiAccountResourceCountsResponseBody
}

type GetMultiAccountResourceCountsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountResourceCountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountResourceCountsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiAccountResourceCountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiAccountResourceCountsResponse) GetBody() *GetMultiAccountResourceCountsResponseBody {
	return s.Body
}

func (s *GetMultiAccountResourceCountsResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceCountsResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceCountsResponse) SetStatusCode(v int32) *GetMultiAccountResourceCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponse) SetBody(v *GetMultiAccountResourceCountsResponseBody) *GetMultiAccountResourceCountsResponse {
	s.Body = v
	return s
}

func (s *GetMultiAccountResourceCountsResponse) Validate() error {
	return dara.Validate(s)
}
