// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStudioLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStudioLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStudioLayoutResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStudioLayoutResponseBody) *ModifyStudioLayoutResponse
	GetBody() *ModifyStudioLayoutResponseBody
}

type ModifyStudioLayoutResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStudioLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStudioLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStudioLayoutResponse) GoString() string {
	return s.String()
}

func (s *ModifyStudioLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStudioLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStudioLayoutResponse) GetBody() *ModifyStudioLayoutResponseBody {
	return s.Body
}

func (s *ModifyStudioLayoutResponse) SetHeaders(v map[string]*string) *ModifyStudioLayoutResponse {
	s.Headers = v
	return s
}

func (s *ModifyStudioLayoutResponse) SetStatusCode(v int32) *ModifyStudioLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStudioLayoutResponse) SetBody(v *ModifyStudioLayoutResponseBody) *ModifyStudioLayoutResponse {
	s.Body = v
	return s
}

func (s *ModifyStudioLayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
