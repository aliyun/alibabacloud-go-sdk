// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIllustration interface {
	dara.Model
	String() string
	GoString() string
	SetIllustrationId(v int64) *Illustration
	GetIllustrationId() *int64
	SetOss(v string) *Illustration
	GetOss() *string
}

type Illustration struct {
	IllustrationId *int64  `json:"illustrationId,omitempty" xml:"illustrationId,omitempty"`
	Oss            *string `json:"oss,omitempty" xml:"oss,omitempty"`
}

func (s Illustration) String() string {
	return dara.Prettify(s)
}

func (s Illustration) GoString() string {
	return s.String()
}

func (s *Illustration) GetIllustrationId() *int64 {
	return s.IllustrationId
}

func (s *Illustration) GetOss() *string {
	return s.Oss
}

func (s *Illustration) SetIllustrationId(v int64) *Illustration {
	s.IllustrationId = &v
	return s
}

func (s *Illustration) SetOss(v string) *Illustration {
	s.Oss = &v
	return s
}

func (s *Illustration) Validate() error {
	return dara.Validate(s)
}
