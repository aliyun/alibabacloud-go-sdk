// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBsnByResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBsnByResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBsnByResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetBsnByResourceResponseBody) *GetBsnByResourceResponse
	GetBody() *GetBsnByResourceResponseBody
}

type GetBsnByResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBsnByResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBsnByResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBsnByResourceResponse) GoString() string {
	return s.String()
}

func (s *GetBsnByResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBsnByResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBsnByResourceResponse) GetBody() *GetBsnByResourceResponseBody {
	return s.Body
}

func (s *GetBsnByResourceResponse) SetHeaders(v map[string]*string) *GetBsnByResourceResponse {
	s.Headers = v
	return s
}

func (s *GetBsnByResourceResponse) SetStatusCode(v int32) *GetBsnByResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBsnByResourceResponse) SetBody(v *GetBsnByResourceResponseBody) *GetBsnByResourceResponse {
	s.Body = v
	return s
}

func (s *GetBsnByResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
