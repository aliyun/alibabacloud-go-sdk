// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePushReviewOnOffResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePushReviewOnOffResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePushReviewOnOffResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePushReviewOnOffResponseBody) *UpdatePushReviewOnOffResponse
	GetBody() *UpdatePushReviewOnOffResponseBody
}

type UpdatePushReviewOnOffResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePushReviewOnOffResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePushReviewOnOffResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushReviewOnOffResponse) GoString() string {
	return s.String()
}

func (s *UpdatePushReviewOnOffResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePushReviewOnOffResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePushReviewOnOffResponse) GetBody() *UpdatePushReviewOnOffResponseBody {
	return s.Body
}

func (s *UpdatePushReviewOnOffResponse) SetHeaders(v map[string]*string) *UpdatePushReviewOnOffResponse {
	s.Headers = v
	return s
}

func (s *UpdatePushReviewOnOffResponse) SetStatusCode(v int32) *UpdatePushReviewOnOffResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePushReviewOnOffResponse) SetBody(v *UpdatePushReviewOnOffResponseBody) *UpdatePushReviewOnOffResponse {
	s.Body = v
	return s
}

func (s *UpdatePushReviewOnOffResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
