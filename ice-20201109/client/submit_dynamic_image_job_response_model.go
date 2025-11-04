// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicImageJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDynamicImageJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDynamicImageJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDynamicImageJobResponseBody) *SubmitDynamicImageJobResponse
	GetBody() *SubmitDynamicImageJobResponseBody
}

type SubmitDynamicImageJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDynamicImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDynamicImageJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDynamicImageJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDynamicImageJobResponse) GetBody() *SubmitDynamicImageJobResponseBody {
	return s.Body
}

func (s *SubmitDynamicImageJobResponse) SetHeaders(v map[string]*string) *SubmitDynamicImageJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDynamicImageJobResponse) SetStatusCode(v int32) *SubmitDynamicImageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDynamicImageJobResponse) SetBody(v *SubmitDynamicImageJobResponseBody) *SubmitDynamicImageJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDynamicImageJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
