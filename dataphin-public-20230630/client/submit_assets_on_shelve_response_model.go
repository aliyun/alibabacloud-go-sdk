// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAssetsOnShelveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAssetsOnShelveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAssetsOnShelveResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAssetsOnShelveResponseBody) *SubmitAssetsOnShelveResponse
	GetBody() *SubmitAssetsOnShelveResponseBody
}

type SubmitAssetsOnShelveResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAssetsOnShelveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAssetsOnShelveResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOnShelveResponse) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOnShelveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAssetsOnShelveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAssetsOnShelveResponse) GetBody() *SubmitAssetsOnShelveResponseBody {
	return s.Body
}

func (s *SubmitAssetsOnShelveResponse) SetHeaders(v map[string]*string) *SubmitAssetsOnShelveResponse {
	s.Headers = v
	return s
}

func (s *SubmitAssetsOnShelveResponse) SetStatusCode(v int32) *SubmitAssetsOnShelveResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAssetsOnShelveResponse) SetBody(v *SubmitAssetsOnShelveResponseBody) *SubmitAssetsOnShelveResponse {
	s.Body = v
	return s
}

func (s *SubmitAssetsOnShelveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
