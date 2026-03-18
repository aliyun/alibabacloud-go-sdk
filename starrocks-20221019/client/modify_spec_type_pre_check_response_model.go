// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySpecTypePreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySpecTypePreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySpecTypePreCheckResponse
	GetStatusCode() *int32
	SetBody(v *ModifySpecTypePreCheckResponseBody) *ModifySpecTypePreCheckResponse
	GetBody() *ModifySpecTypePreCheckResponseBody
}

type ModifySpecTypePreCheckResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySpecTypePreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySpecTypePreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySpecTypePreCheckResponse) GoString() string {
	return s.String()
}

func (s *ModifySpecTypePreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySpecTypePreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySpecTypePreCheckResponse) GetBody() *ModifySpecTypePreCheckResponseBody {
	return s.Body
}

func (s *ModifySpecTypePreCheckResponse) SetHeaders(v map[string]*string) *ModifySpecTypePreCheckResponse {
	s.Headers = v
	return s
}

func (s *ModifySpecTypePreCheckResponse) SetStatusCode(v int32) *ModifySpecTypePreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySpecTypePreCheckResponse) SetBody(v *ModifySpecTypePreCheckResponseBody) *ModifySpecTypePreCheckResponse {
	s.Body = v
	return s
}

func (s *ModifySpecTypePreCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
