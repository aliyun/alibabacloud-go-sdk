// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveDomainMultiStreamListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLiveDomainMultiStreamListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLiveDomainMultiStreamListResponse
	GetStatusCode() *int32
	SetBody(v *QueryLiveDomainMultiStreamListResponseBody) *QueryLiveDomainMultiStreamListResponse
	GetBody() *QueryLiveDomainMultiStreamListResponseBody
}

type QueryLiveDomainMultiStreamListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveDomainMultiStreamListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveDomainMultiStreamListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveDomainMultiStreamListResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveDomainMultiStreamListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLiveDomainMultiStreamListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLiveDomainMultiStreamListResponse) GetBody() *QueryLiveDomainMultiStreamListResponseBody {
	return s.Body
}

func (s *QueryLiveDomainMultiStreamListResponse) SetHeaders(v map[string]*string) *QueryLiveDomainMultiStreamListResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponse) SetStatusCode(v int32) *QueryLiveDomainMultiStreamListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponse) SetBody(v *QueryLiveDomainMultiStreamListResponseBody) *QueryLiveDomainMultiStreamListResponse {
	s.Body = v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponse) Validate() error {
	return dara.Validate(s)
}
