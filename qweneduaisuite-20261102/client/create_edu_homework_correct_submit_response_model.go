// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduHomeworkCorrectSubmitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEduHomeworkCorrectSubmitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEduHomeworkCorrectSubmitResponse
	GetStatusCode() *int32
	SetBody(v *CreateEduHomeworkCorrectSubmitResponseBody) *CreateEduHomeworkCorrectSubmitResponse
	GetBody() *CreateEduHomeworkCorrectSubmitResponseBody
}

type CreateEduHomeworkCorrectSubmitResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEduHomeworkCorrectSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitResponse) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEduHomeworkCorrectSubmitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEduHomeworkCorrectSubmitResponse) GetBody() *CreateEduHomeworkCorrectSubmitResponseBody {
	return s.Body
}

func (s *CreateEduHomeworkCorrectSubmitResponse) SetHeaders(v map[string]*string) *CreateEduHomeworkCorrectSubmitResponse {
	s.Headers = v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponse) SetStatusCode(v int32) *CreateEduHomeworkCorrectSubmitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponse) SetBody(v *CreateEduHomeworkCorrectSubmitResponseBody) *CreateEduHomeworkCorrectSubmitResponse {
	s.Body = v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
