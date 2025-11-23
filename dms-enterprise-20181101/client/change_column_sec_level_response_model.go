// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeColumnSecLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeColumnSecLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeColumnSecLevelResponse
	GetStatusCode() *int32
	SetBody(v *ChangeColumnSecLevelResponseBody) *ChangeColumnSecLevelResponse
	GetBody() *ChangeColumnSecLevelResponseBody
}

type ChangeColumnSecLevelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeColumnSecLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeColumnSecLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeColumnSecLevelResponse) GoString() string {
	return s.String()
}

func (s *ChangeColumnSecLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeColumnSecLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeColumnSecLevelResponse) GetBody() *ChangeColumnSecLevelResponseBody {
	return s.Body
}

func (s *ChangeColumnSecLevelResponse) SetHeaders(v map[string]*string) *ChangeColumnSecLevelResponse {
	s.Headers = v
	return s
}

func (s *ChangeColumnSecLevelResponse) SetStatusCode(v int32) *ChangeColumnSecLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeColumnSecLevelResponse) SetBody(v *ChangeColumnSecLevelResponseBody) *ChangeColumnSecLevelResponse {
	s.Body = v
	return s
}

func (s *ChangeColumnSecLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
