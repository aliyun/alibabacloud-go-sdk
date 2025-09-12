// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildStsAKRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v int64) *BuildStsAKRequest
	GetBody() *int64
}

type BuildStsAKRequest struct {
	Body *int64 `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildStsAKRequest) String() string {
	return dara.Prettify(s)
}

func (s BuildStsAKRequest) GoString() string {
	return s.String()
}

func (s *BuildStsAKRequest) GetBody() *int64 {
	return s.Body
}

func (s *BuildStsAKRequest) SetBody(v int64) *BuildStsAKRequest {
	s.Body = &v
	return s
}

func (s *BuildStsAKRequest) Validate() error {
	return dara.Validate(s)
}
