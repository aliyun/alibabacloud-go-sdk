// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNodeBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNodeBatchResponse
	GetStatusCode() *int32
	SetBody(v *CreateNodeBatchResponseBody) *CreateNodeBatchResponse
	GetBody() *CreateNodeBatchResponseBody
}

type CreateNodeBatchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeBatchResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNodeBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNodeBatchResponse) GetBody() *CreateNodeBatchResponseBody {
	return s.Body
}

func (s *CreateNodeBatchResponse) SetHeaders(v map[string]*string) *CreateNodeBatchResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeBatchResponse) SetStatusCode(v int32) *CreateNodeBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeBatchResponse) SetBody(v *CreateNodeBatchResponseBody) *CreateNodeBatchResponse {
	s.Body = v
	return s
}

func (s *CreateNodeBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
