// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDialogDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDialogDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDialogDetailResponseBody) *GetDialogDetailResponse
	GetBody() *GetDialogDetailResponseBody
}

type GetDialogDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDialogDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDialogDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDialogDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDialogDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDialogDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDialogDetailResponse) GetBody() *GetDialogDetailResponseBody {
	return s.Body
}

func (s *GetDialogDetailResponse) SetHeaders(v map[string]*string) *GetDialogDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDialogDetailResponse) SetStatusCode(v int32) *GetDialogDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDialogDetailResponse) SetBody(v *GetDialogDetailResponseBody) *GetDialogDetailResponse {
	s.Body = v
	return s
}

func (s *GetDialogDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
