// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInput interface {
	dara.Model
	String() string
	GoString() string
	SetOSS(v *InputOSS) *Input
	GetOSS() *InputOSS
}

type Input struct {
	OSS *InputOSS `json:"OSS,omitempty" xml:"OSS,omitempty"`
}

func (s Input) String() string {
	return dara.Prettify(s)
}

func (s Input) GoString() string {
	return s.String()
}

func (s *Input) GetOSS() *InputOSS {
	return s.OSS
}

func (s *Input) SetOSS(v *InputOSS) *Input {
	s.OSS = v
	return s
}

func (s *Input) Validate() error {
	if s.OSS != nil {
		if err := s.OSS.Validate(); err != nil {
			return err
		}
	}
	return nil
}
