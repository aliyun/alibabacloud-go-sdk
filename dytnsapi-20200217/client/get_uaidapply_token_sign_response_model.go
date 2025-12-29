// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUAIDApplyTokenSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUAIDApplyTokenSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUAIDApplyTokenSignResponse
	GetStatusCode() *int32
	SetBody(v *GetUAIDApplyTokenSignResponseBody) *GetUAIDApplyTokenSignResponse
	GetBody() *GetUAIDApplyTokenSignResponseBody
}

type GetUAIDApplyTokenSignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUAIDApplyTokenSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUAIDApplyTokenSignResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUAIDApplyTokenSignResponse) GoString() string {
	return s.String()
}

func (s *GetUAIDApplyTokenSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUAIDApplyTokenSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUAIDApplyTokenSignResponse) GetBody() *GetUAIDApplyTokenSignResponseBody {
	return s.Body
}

func (s *GetUAIDApplyTokenSignResponse) SetHeaders(v map[string]*string) *GetUAIDApplyTokenSignResponse {
	s.Headers = v
	return s
}

func (s *GetUAIDApplyTokenSignResponse) SetStatusCode(v int32) *GetUAIDApplyTokenSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUAIDApplyTokenSignResponse) SetBody(v *GetUAIDApplyTokenSignResponseBody) *GetUAIDApplyTokenSignResponse {
	s.Body = v
	return s
}

func (s *GetUAIDApplyTokenSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
