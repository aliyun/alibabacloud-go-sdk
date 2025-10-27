// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageEventOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageEventOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageEventOperationResponse
	GetStatusCode() *int32
	SetBody(v *GetImageEventOperationResponseBody) *GetImageEventOperationResponse
	GetBody() *GetImageEventOperationResponseBody
}

type GetImageEventOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageEventOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageEventOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageEventOperationResponse) GoString() string {
	return s.String()
}

func (s *GetImageEventOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageEventOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageEventOperationResponse) GetBody() *GetImageEventOperationResponseBody {
	return s.Body
}

func (s *GetImageEventOperationResponse) SetHeaders(v map[string]*string) *GetImageEventOperationResponse {
	s.Headers = v
	return s
}

func (s *GetImageEventOperationResponse) SetStatusCode(v int32) *GetImageEventOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageEventOperationResponse) SetBody(v *GetImageEventOperationResponseBody) *GetImageEventOperationResponse {
	s.Body = v
	return s
}

func (s *GetImageEventOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
