// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateDisPlayModesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOrUpdateDisPlayModesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOrUpdateDisPlayModesResponse
	GetStatusCode() *int32
	SetBody(v *AddOrUpdateDisPlayModesResponseBody) *AddOrUpdateDisPlayModesResponse
	GetBody() *AddOrUpdateDisPlayModesResponseBody
}

type AddOrUpdateDisPlayModesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateDisPlayModesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateDisPlayModesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateDisPlayModesResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOrUpdateDisPlayModesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateDisPlayModesResponse) GetBody() *AddOrUpdateDisPlayModesResponseBody {
	return s.Body
}

func (s *AddOrUpdateDisPlayModesResponse) SetHeaders(v map[string]*string) *AddOrUpdateDisPlayModesResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateDisPlayModesResponse) SetStatusCode(v int32) *AddOrUpdateDisPlayModesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponse) SetBody(v *AddOrUpdateDisPlayModesResponseBody) *AddOrUpdateDisPlayModesResponse {
	s.Body = v
	return s
}

func (s *AddOrUpdateDisPlayModesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
