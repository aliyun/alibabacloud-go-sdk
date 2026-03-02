// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHistoryDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHistoryDeveloperResponse
	GetStatusCode() *int32
	SetBody(v *Developer) *GetHistoryDeveloperResponse
	GetBody() *Developer
}

type GetHistoryDeveloperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Developer         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoryDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryDeveloperResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHistoryDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHistoryDeveloperResponse) GetBody() *Developer {
	return s.Body
}

func (s *GetHistoryDeveloperResponse) SetHeaders(v map[string]*string) *GetHistoryDeveloperResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryDeveloperResponse) SetStatusCode(v int32) *GetHistoryDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryDeveloperResponse) SetBody(v *Developer) *GetHistoryDeveloperResponse {
	s.Body = v
	return s
}

func (s *GetHistoryDeveloperResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
