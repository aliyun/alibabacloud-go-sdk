// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResubmitConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResubmitConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResubmitConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResubmitConfigResponseBody) *ModifyResubmitConfigResponse
	GetBody() *ModifyResubmitConfigResponseBody
}

type ModifyResubmitConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResubmitConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResubmitConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResubmitConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResubmitConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResubmitConfigResponse) GetBody() *ModifyResubmitConfigResponseBody {
	return s.Body
}

func (s *ModifyResubmitConfigResponse) SetHeaders(v map[string]*string) *ModifyResubmitConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyResubmitConfigResponse) SetStatusCode(v int32) *ModifyResubmitConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResubmitConfigResponse) SetBody(v *ModifyResubmitConfigResponseBody) *ModifyResubmitConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyResubmitConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
