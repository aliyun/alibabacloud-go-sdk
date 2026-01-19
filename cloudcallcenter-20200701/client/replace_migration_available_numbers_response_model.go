// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceMigrationAvailableNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceMigrationAvailableNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceMigrationAvailableNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceMigrationAvailableNumbersResponseBody) *ReplaceMigrationAvailableNumbersResponse
	GetBody() *ReplaceMigrationAvailableNumbersResponseBody
}

type ReplaceMigrationAvailableNumbersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceMigrationAvailableNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceMigrationAvailableNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceMigrationAvailableNumbersResponse) GoString() string {
	return s.String()
}

func (s *ReplaceMigrationAvailableNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceMigrationAvailableNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceMigrationAvailableNumbersResponse) GetBody() *ReplaceMigrationAvailableNumbersResponseBody {
	return s.Body
}

func (s *ReplaceMigrationAvailableNumbersResponse) SetHeaders(v map[string]*string) *ReplaceMigrationAvailableNumbersResponse {
	s.Headers = v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponse) SetStatusCode(v int32) *ReplaceMigrationAvailableNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponse) SetBody(v *ReplaceMigrationAvailableNumbersResponseBody) *ReplaceMigrationAvailableNumbersResponse {
	s.Body = v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
