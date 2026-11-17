// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogicalModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogicalModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogicalModelResponse
	GetStatusCode() *int32
	SetBody(v *GetLogicalModelResponseBody) *GetLogicalModelResponse
	GetBody() *GetLogicalModelResponseBody
}

type GetLogicalModelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogicalModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogicalModelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogicalModelResponse) GoString() string {
	return s.String()
}

func (s *GetLogicalModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogicalModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogicalModelResponse) GetBody() *GetLogicalModelResponseBody {
	return s.Body
}

func (s *GetLogicalModelResponse) SetHeaders(v map[string]*string) *GetLogicalModelResponse {
	s.Headers = v
	return s
}

func (s *GetLogicalModelResponse) SetStatusCode(v int32) *GetLogicalModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogicalModelResponse) SetBody(v *GetLogicalModelResponseBody) *GetLogicalModelResponse {
	s.Body = v
	return s
}

func (s *GetLogicalModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
