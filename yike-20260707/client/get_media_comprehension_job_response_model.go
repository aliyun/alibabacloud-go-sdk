// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaComprehensionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaComprehensionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaComprehensionJobResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaComprehensionJobResponseBody) *GetMediaComprehensionJobResponse
	GetBody() *GetMediaComprehensionJobResponseBody
}

type GetMediaComprehensionJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaComprehensionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaComprehensionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaComprehensionJobResponse) GoString() string {
	return s.String()
}

func (s *GetMediaComprehensionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaComprehensionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaComprehensionJobResponse) GetBody() *GetMediaComprehensionJobResponseBody {
	return s.Body
}

func (s *GetMediaComprehensionJobResponse) SetHeaders(v map[string]*string) *GetMediaComprehensionJobResponse {
	s.Headers = v
	return s
}

func (s *GetMediaComprehensionJobResponse) SetStatusCode(v int32) *GetMediaComprehensionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaComprehensionJobResponse) SetBody(v *GetMediaComprehensionJobResponseBody) *GetMediaComprehensionJobResponse {
	s.Body = v
	return s
}

func (s *GetMediaComprehensionJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
