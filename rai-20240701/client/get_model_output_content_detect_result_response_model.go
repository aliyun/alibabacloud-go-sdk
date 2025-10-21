// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelOutputContentDetectResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelOutputContentDetectResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelOutputContentDetectResultResponse
	GetStatusCode() *int32
	SetBody(v *GetModelOutputContentDetectResultResponseBody) *GetModelOutputContentDetectResultResponse
	GetBody() *GetModelOutputContentDetectResultResponseBody
}

type GetModelOutputContentDetectResultResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelOutputContentDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelOutputContentDetectResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponse) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelOutputContentDetectResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelOutputContentDetectResultResponse) GetBody() *GetModelOutputContentDetectResultResponseBody {
	return s.Body
}

func (s *GetModelOutputContentDetectResultResponse) SetHeaders(v map[string]*string) *GetModelOutputContentDetectResultResponse {
	s.Headers = v
	return s
}

func (s *GetModelOutputContentDetectResultResponse) SetStatusCode(v int32) *GetModelOutputContentDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponse) SetBody(v *GetModelOutputContentDetectResultResponseBody) *GetModelOutputContentDetectResultResponse {
	s.Body = v
	return s
}

func (s *GetModelOutputContentDetectResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
