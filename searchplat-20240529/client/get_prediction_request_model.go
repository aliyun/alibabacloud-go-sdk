// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPredictionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *GetPredictionRequest
	GetBody() *string
}

type GetPredictionRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPredictionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPredictionRequest) GoString() string {
	return s.String()
}

func (s *GetPredictionRequest) GetBody() *string {
	return s.Body
}

func (s *GetPredictionRequest) SetBody(v string) *GetPredictionRequest {
	s.Body = &v
	return s
}

func (s *GetPredictionRequest) Validate() error {
	return dara.Validate(s)
}
