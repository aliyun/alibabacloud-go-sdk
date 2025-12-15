// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPredictionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPredictionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPredictionResponse
	GetStatusCode() *int32
	SetBody(v string) *GetPredictionResponse
	GetBody() *string
}

type GetPredictionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPredictionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPredictionResponse) GoString() string {
	return s.String()
}

func (s *GetPredictionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPredictionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPredictionResponse) GetBody() *string {
	return s.Body
}

func (s *GetPredictionResponse) SetHeaders(v map[string]*string) *GetPredictionResponse {
	s.Headers = v
	return s
}

func (s *GetPredictionResponse) SetStatusCode(v int32) *GetPredictionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPredictionResponse) SetBody(v string) *GetPredictionResponse {
	s.Body = &v
	return s
}

func (s *GetPredictionResponse) Validate() error {
	return dara.Validate(s)
}
