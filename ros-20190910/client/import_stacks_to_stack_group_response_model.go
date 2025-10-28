// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportStacksToStackGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportStacksToStackGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportStacksToStackGroupResponse
	GetStatusCode() *int32
	SetBody(v *ImportStacksToStackGroupResponseBody) *ImportStacksToStackGroupResponse
	GetBody() *ImportStacksToStackGroupResponseBody
}

type ImportStacksToStackGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportStacksToStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportStacksToStackGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportStacksToStackGroupResponse) GoString() string {
	return s.String()
}

func (s *ImportStacksToStackGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportStacksToStackGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportStacksToStackGroupResponse) GetBody() *ImportStacksToStackGroupResponseBody {
	return s.Body
}

func (s *ImportStacksToStackGroupResponse) SetHeaders(v map[string]*string) *ImportStacksToStackGroupResponse {
	s.Headers = v
	return s
}

func (s *ImportStacksToStackGroupResponse) SetStatusCode(v int32) *ImportStacksToStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportStacksToStackGroupResponse) SetBody(v *ImportStacksToStackGroupResponseBody) *ImportStacksToStackGroupResponse {
	s.Body = v
	return s
}

func (s *ImportStacksToStackGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
