// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetMem0AccountPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetMem0AccountPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetMem0AccountPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetMem0AccountPasswordResponseBody) *ResetMem0AccountPasswordResponse
	GetBody() *ResetMem0AccountPasswordResponseBody
}

type ResetMem0AccountPasswordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetMem0AccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetMem0AccountPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetMem0AccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetMem0AccountPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetMem0AccountPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetMem0AccountPasswordResponse) GetBody() *ResetMem0AccountPasswordResponseBody {
	return s.Body
}

func (s *ResetMem0AccountPasswordResponse) SetHeaders(v map[string]*string) *ResetMem0AccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetMem0AccountPasswordResponse) SetStatusCode(v int32) *ResetMem0AccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetMem0AccountPasswordResponse) SetBody(v *ResetMem0AccountPasswordResponseBody) *ResetMem0AccountPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetMem0AccountPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
