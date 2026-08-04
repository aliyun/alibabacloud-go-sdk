// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindBizCategoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindBizCategoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindBizCategoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *FindBizCategoryConfigResponseBody) *FindBizCategoryConfigResponse
	GetBody() *FindBizCategoryConfigResponseBody
}

type FindBizCategoryConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindBizCategoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindBizCategoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s FindBizCategoryConfigResponse) GoString() string {
	return s.String()
}

func (s *FindBizCategoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindBizCategoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindBizCategoryConfigResponse) GetBody() *FindBizCategoryConfigResponseBody {
	return s.Body
}

func (s *FindBizCategoryConfigResponse) SetHeaders(v map[string]*string) *FindBizCategoryConfigResponse {
	s.Headers = v
	return s
}

func (s *FindBizCategoryConfigResponse) SetStatusCode(v int32) *FindBizCategoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *FindBizCategoryConfigResponse) SetBody(v *FindBizCategoryConfigResponseBody) *FindBizCategoryConfigResponse {
	s.Body = v
	return s
}

func (s *FindBizCategoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
