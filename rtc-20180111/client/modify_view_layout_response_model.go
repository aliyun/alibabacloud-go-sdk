// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyViewLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyViewLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyViewLayoutResponse
	GetStatusCode() *int32
	SetBody(v *ModifyViewLayoutResponseBody) *ModifyViewLayoutResponse
	GetBody() *ModifyViewLayoutResponseBody
}

type ModifyViewLayoutResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyViewLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyViewLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutResponse) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyViewLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyViewLayoutResponse) GetBody() *ModifyViewLayoutResponseBody {
	return s.Body
}

func (s *ModifyViewLayoutResponse) SetHeaders(v map[string]*string) *ModifyViewLayoutResponse {
	s.Headers = v
	return s
}

func (s *ModifyViewLayoutResponse) SetStatusCode(v int32) *ModifyViewLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyViewLayoutResponse) SetBody(v *ModifyViewLayoutResponseBody) *ModifyViewLayoutResponse {
	s.Body = v
	return s
}

func (s *ModifyViewLayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
