// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAppMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAppMetadataResponse
	GetStatusCode() *int32
	SetBody(v *QueryAppMetadataResponseBody) *QueryAppMetadataResponse
	GetBody() *QueryAppMetadataResponseBody
}

type QueryAppMetadataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAppMetadataResponse) GoString() string {
	return s.String()
}

func (s *QueryAppMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAppMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAppMetadataResponse) GetBody() *QueryAppMetadataResponseBody {
	return s.Body
}

func (s *QueryAppMetadataResponse) SetHeaders(v map[string]*string) *QueryAppMetadataResponse {
	s.Headers = v
	return s
}

func (s *QueryAppMetadataResponse) SetStatusCode(v int32) *QueryAppMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppMetadataResponse) SetBody(v *QueryAppMetadataResponseBody) *QueryAppMetadataResponse {
	s.Body = v
	return s
}

func (s *QueryAppMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
