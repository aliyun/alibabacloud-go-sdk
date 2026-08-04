// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchCreateModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterBatchCreateModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterBatchCreateModelResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterBatchCreateModelResponseBody) *ModelRouterBatchCreateModelResponse
	GetBody() *ModelRouterBatchCreateModelResponseBody
}

type ModelRouterBatchCreateModelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterBatchCreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterBatchCreateModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchCreateModelResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchCreateModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterBatchCreateModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterBatchCreateModelResponse) GetBody() *ModelRouterBatchCreateModelResponseBody {
	return s.Body
}

func (s *ModelRouterBatchCreateModelResponse) SetHeaders(v map[string]*string) *ModelRouterBatchCreateModelResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterBatchCreateModelResponse) SetStatusCode(v int32) *ModelRouterBatchCreateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterBatchCreateModelResponse) SetBody(v *ModelRouterBatchCreateModelResponseBody) *ModelRouterBatchCreateModelResponse {
	s.Body = v
	return s
}

func (s *ModelRouterBatchCreateModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
