// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelServiceResponse
	GetStatusCode() *int32
	SetBody(v *ModelServiceResult) *GetModelServiceResponse
	GetBody() *ModelServiceResult
}

type GetModelServiceResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelServiceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelServiceResponse) GoString() string {
	return s.String()
}

func (s *GetModelServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelServiceResponse) GetBody() *ModelServiceResult {
	return s.Body
}

func (s *GetModelServiceResponse) SetHeaders(v map[string]*string) *GetModelServiceResponse {
	s.Headers = v
	return s
}

func (s *GetModelServiceResponse) SetStatusCode(v int32) *GetModelServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelServiceResponse) SetBody(v *ModelServiceResult) *GetModelServiceResponse {
	s.Body = v
	return s
}

func (s *GetModelServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
