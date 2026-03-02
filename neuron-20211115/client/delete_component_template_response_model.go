// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComponentTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComponentTemplateResponse
	GetStatusCode() *int32
}

type DeleteComponentTemplateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteComponentTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComponentTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComponentTemplateResponse) SetHeaders(v map[string]*string) *DeleteComponentTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentTemplateResponse) SetStatusCode(v int32) *DeleteComponentTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComponentTemplateResponse) Validate() error {
	return dara.Validate(s)
}
