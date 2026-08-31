// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAssetsOffShelveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAssetsOffShelveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAssetsOffShelveResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAssetsOffShelveResponseBody) *SubmitAssetsOffShelveResponse
	GetBody() *SubmitAssetsOffShelveResponseBody
}

type SubmitAssetsOffShelveResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAssetsOffShelveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAssetsOffShelveResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOffShelveResponse) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOffShelveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAssetsOffShelveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAssetsOffShelveResponse) GetBody() *SubmitAssetsOffShelveResponseBody {
	return s.Body
}

func (s *SubmitAssetsOffShelveResponse) SetHeaders(v map[string]*string) *SubmitAssetsOffShelveResponse {
	s.Headers = v
	return s
}

func (s *SubmitAssetsOffShelveResponse) SetStatusCode(v int32) *SubmitAssetsOffShelveResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAssetsOffShelveResponse) SetBody(v *SubmitAssetsOffShelveResponseBody) *SubmitAssetsOffShelveResponse {
	s.Body = v
	return s
}

func (s *SubmitAssetsOffShelveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
