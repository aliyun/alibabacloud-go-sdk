// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v int64) *SelectDataSetRequest
	GetBody() *int64
}

type SelectDataSetRequest struct {
	Body *int64 `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectDataSetRequest) GoString() string {
	return s.String()
}

func (s *SelectDataSetRequest) GetBody() *int64 {
	return s.Body
}

func (s *SelectDataSetRequest) SetBody(v int64) *SelectDataSetRequest {
	s.Body = &v
	return s
}

func (s *SelectDataSetRequest) Validate() error {
	return dara.Validate(s)
}
