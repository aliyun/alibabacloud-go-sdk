// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLabelGeneratedResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLabelGeneratedResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLabelGeneratedResultResponse
	GetStatusCode() *int32
	SetBody(v *GetLabelGeneratedResultResponseBody) *GetLabelGeneratedResultResponse
	GetBody() *GetLabelGeneratedResultResponseBody
}

type GetLabelGeneratedResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLabelGeneratedResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLabelGeneratedResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLabelGeneratedResultResponse) GoString() string {
	return s.String()
}

func (s *GetLabelGeneratedResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLabelGeneratedResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLabelGeneratedResultResponse) GetBody() *GetLabelGeneratedResultResponseBody {
	return s.Body
}

func (s *GetLabelGeneratedResultResponse) SetHeaders(v map[string]*string) *GetLabelGeneratedResultResponse {
	s.Headers = v
	return s
}

func (s *GetLabelGeneratedResultResponse) SetStatusCode(v int32) *GetLabelGeneratedResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLabelGeneratedResultResponse) SetBody(v *GetLabelGeneratedResultResponseBody) *GetLabelGeneratedResultResponse {
	s.Body = v
	return s
}

func (s *GetLabelGeneratedResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
