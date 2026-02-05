// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MultimodalSearchBody) *MultimodalSearchRequest
	GetBody() *MultimodalSearchBody
}

type MultimodalSearchRequest struct {
	Body *MultimodalSearchBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultimodalSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s MultimodalSearchRequest) GoString() string {
	return s.String()
}

func (s *MultimodalSearchRequest) GetBody() *MultimodalSearchBody {
	return s.Body
}

func (s *MultimodalSearchRequest) SetBody(v *MultimodalSearchBody) *MultimodalSearchRequest {
	s.Body = v
	return s
}

func (s *MultimodalSearchRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
