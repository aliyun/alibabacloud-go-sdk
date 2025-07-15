// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSpMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSpMetadataResponse
	GetStatusCode() *int32
	SetBody(v *GetSpMetadataResponseBody) *GetSpMetadataResponse
	GetBody() *GetSpMetadataResponseBody
}

type GetSpMetadataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSpMetadataResponse) GoString() string {
	return s.String()
}

func (s *GetSpMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSpMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSpMetadataResponse) GetBody() *GetSpMetadataResponseBody {
	return s.Body
}

func (s *GetSpMetadataResponse) SetHeaders(v map[string]*string) *GetSpMetadataResponse {
	s.Headers = v
	return s
}

func (s *GetSpMetadataResponse) SetStatusCode(v int32) *GetSpMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpMetadataResponse) SetBody(v *GetSpMetadataResponseBody) *GetSpMetadataResponse {
	s.Body = v
	return s
}

func (s *GetSpMetadataResponse) Validate() error {
	return dara.Validate(s)
}
