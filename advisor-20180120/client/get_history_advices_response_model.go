// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryAdvicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHistoryAdvicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHistoryAdvicesResponse
	GetStatusCode() *int32
	SetBody(v *GetHistoryAdvicesResponseBody) *GetHistoryAdvicesResponse
	GetBody() *GetHistoryAdvicesResponseBody
}

type GetHistoryAdvicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoryAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoryAdvicesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryAdvicesResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryAdvicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHistoryAdvicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHistoryAdvicesResponse) GetBody() *GetHistoryAdvicesResponseBody {
	return s.Body
}

func (s *GetHistoryAdvicesResponse) SetHeaders(v map[string]*string) *GetHistoryAdvicesResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryAdvicesResponse) SetStatusCode(v int32) *GetHistoryAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryAdvicesResponse) SetBody(v *GetHistoryAdvicesResponseBody) *GetHistoryAdvicesResponse {
	s.Body = v
	return s
}

func (s *GetHistoryAdvicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
