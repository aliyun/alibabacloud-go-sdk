// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubtaskItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubtaskItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubtaskItemResponse
	GetStatusCode() *int32
	SetBody(v *GetSubtaskItemResponseBody) *GetSubtaskItemResponse
	GetBody() *GetSubtaskItemResponseBody
}

type GetSubtaskItemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubtaskItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubtaskItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubtaskItemResponse) GoString() string {
	return s.String()
}

func (s *GetSubtaskItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubtaskItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubtaskItemResponse) GetBody() *GetSubtaskItemResponseBody {
	return s.Body
}

func (s *GetSubtaskItemResponse) SetHeaders(v map[string]*string) *GetSubtaskItemResponse {
	s.Headers = v
	return s
}

func (s *GetSubtaskItemResponse) SetStatusCode(v int32) *GetSubtaskItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubtaskItemResponse) SetBody(v *GetSubtaskItemResponseBody) *GetSubtaskItemResponse {
	s.Body = v
	return s
}

func (s *GetSubtaskItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
