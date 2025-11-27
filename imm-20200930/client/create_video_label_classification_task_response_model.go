// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoLabelClassificationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoLabelClassificationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoLabelClassificationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoLabelClassificationTaskResponseBody) *CreateVideoLabelClassificationTaskResponse
	GetBody() *CreateVideoLabelClassificationTaskResponseBody
}

type CreateVideoLabelClassificationTaskResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoLabelClassificationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoLabelClassificationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoLabelClassificationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoLabelClassificationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoLabelClassificationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoLabelClassificationTaskResponse) GetBody() *CreateVideoLabelClassificationTaskResponseBody {
	return s.Body
}

func (s *CreateVideoLabelClassificationTaskResponse) SetHeaders(v map[string]*string) *CreateVideoLabelClassificationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponse) SetStatusCode(v int32) *CreateVideoLabelClassificationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponse) SetBody(v *CreateVideoLabelClassificationTaskResponseBody) *CreateVideoLabelClassificationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVideoLabelClassificationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
