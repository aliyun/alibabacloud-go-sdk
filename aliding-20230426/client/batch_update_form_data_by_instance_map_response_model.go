// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateFormDataByInstanceMapResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateFormDataByInstanceMapResponseBody) *BatchUpdateFormDataByInstanceMapResponse
	GetBody() *BatchUpdateFormDataByInstanceMapResponseBody
}

type BatchUpdateFormDataByInstanceMapResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateFormDataByInstanceMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateFormDataByInstanceMapResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateFormDataByInstanceMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateFormDataByInstanceMapResponse) GetBody() *BatchUpdateFormDataByInstanceMapResponseBody {
	return s.Body
}

func (s *BatchUpdateFormDataByInstanceMapResponse) SetHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceMapResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapResponse) SetStatusCode(v int32) *BatchUpdateFormDataByInstanceMapResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapResponse) SetBody(v *BatchUpdateFormDataByInstanceMapResponseBody) *BatchUpdateFormDataByInstanceMapResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
