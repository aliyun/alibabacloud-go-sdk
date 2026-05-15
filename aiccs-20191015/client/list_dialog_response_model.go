// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDialogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDialogResponse
	GetStatusCode() *int32
	SetBody(v *ListDialogResponseBody) *ListDialogResponse
	GetBody() *ListDialogResponseBody
}

type ListDialogResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDialogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDialogResponse) GoString() string {
	return s.String()
}

func (s *ListDialogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDialogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDialogResponse) GetBody() *ListDialogResponseBody {
	return s.Body
}

func (s *ListDialogResponse) SetHeaders(v map[string]*string) *ListDialogResponse {
	s.Headers = v
	return s
}

func (s *ListDialogResponse) SetStatusCode(v int32) *ListDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDialogResponse) SetBody(v *ListDialogResponseBody) *ListDialogResponse {
	s.Body = v
	return s
}

func (s *ListDialogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
