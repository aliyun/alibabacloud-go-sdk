// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelVersionLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelVersionLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelVersionLabelsResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelVersionLabelsResponseBody) *CreateModelVersionLabelsResponse
	GetBody() *CreateModelVersionLabelsResponseBody
}

type CreateModelVersionLabelsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelVersionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelVersionLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelVersionLabelsResponse) GoString() string {
	return s.String()
}

func (s *CreateModelVersionLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelVersionLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelVersionLabelsResponse) GetBody() *CreateModelVersionLabelsResponseBody {
	return s.Body
}

func (s *CreateModelVersionLabelsResponse) SetHeaders(v map[string]*string) *CreateModelVersionLabelsResponse {
	s.Headers = v
	return s
}

func (s *CreateModelVersionLabelsResponse) SetStatusCode(v int32) *CreateModelVersionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelVersionLabelsResponse) SetBody(v *CreateModelVersionLabelsResponseBody) *CreateModelVersionLabelsResponse {
	s.Body = v
	return s
}

func (s *CreateModelVersionLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
