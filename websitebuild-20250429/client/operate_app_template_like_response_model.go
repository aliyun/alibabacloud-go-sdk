// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppTemplateLikeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateAppTemplateLikeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateAppTemplateLikeResponse
	GetStatusCode() *int32
	SetBody(v *OperateAppTemplateLikeResponseBody) *OperateAppTemplateLikeResponse
	GetBody() *OperateAppTemplateLikeResponseBody
}

type OperateAppTemplateLikeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAppTemplateLikeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAppTemplateLikeResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateAppTemplateLikeResponse) GoString() string {
	return s.String()
}

func (s *OperateAppTemplateLikeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateAppTemplateLikeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateAppTemplateLikeResponse) GetBody() *OperateAppTemplateLikeResponseBody {
	return s.Body
}

func (s *OperateAppTemplateLikeResponse) SetHeaders(v map[string]*string) *OperateAppTemplateLikeResponse {
	s.Headers = v
	return s
}

func (s *OperateAppTemplateLikeResponse) SetStatusCode(v int32) *OperateAppTemplateLikeResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAppTemplateLikeResponse) SetBody(v *OperateAppTemplateLikeResponseBody) *OperateAppTemplateLikeResponse {
	s.Body = v
	return s
}

func (s *OperateAppTemplateLikeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
