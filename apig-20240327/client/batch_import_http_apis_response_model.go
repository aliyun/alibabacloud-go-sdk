// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchImportHttpApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchImportHttpApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchImportHttpApisResponse
	GetStatusCode() *int32
	SetBody(v *BatchImportHttpApisResponseBody) *BatchImportHttpApisResponse
	GetBody() *BatchImportHttpApisResponseBody
}

type BatchImportHttpApisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchImportHttpApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchImportHttpApisResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchImportHttpApisResponse) GoString() string {
	return s.String()
}

func (s *BatchImportHttpApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchImportHttpApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchImportHttpApisResponse) GetBody() *BatchImportHttpApisResponseBody {
	return s.Body
}

func (s *BatchImportHttpApisResponse) SetHeaders(v map[string]*string) *BatchImportHttpApisResponse {
	s.Headers = v
	return s
}

func (s *BatchImportHttpApisResponse) SetStatusCode(v int32) *BatchImportHttpApisResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchImportHttpApisResponse) SetBody(v *BatchImportHttpApisResponseBody) *BatchImportHttpApisResponse {
	s.Body = v
	return s
}

func (s *BatchImportHttpApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
