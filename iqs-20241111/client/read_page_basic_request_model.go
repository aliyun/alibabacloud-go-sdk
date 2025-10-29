// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageBasicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ReadPageBody) *ReadPageBasicRequest
	GetBody() *ReadPageBody
}

type ReadPageBasicRequest struct {
	// post body
	Body *ReadPageBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadPageBasicRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadPageBasicRequest) GoString() string {
	return s.String()
}

func (s *ReadPageBasicRequest) GetBody() *ReadPageBody {
	return s.Body
}

func (s *ReadPageBasicRequest) SetBody(v *ReadPageBody) *ReadPageBasicRequest {
	s.Body = v
	return s
}

func (s *ReadPageBasicRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
