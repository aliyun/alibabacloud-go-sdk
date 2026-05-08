// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIllustrationResult interface {
	dara.Model
	String() string
	GoString() string
	SetIllustration(v *Illustration) *IllustrationResult
	GetIllustration() *Illustration
	SetRequestId(v string) *IllustrationResult
	GetRequestId() *string
}

type IllustrationResult struct {
	Illustration *Illustration `json:"illustration,omitempty" xml:"illustration,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s IllustrationResult) String() string {
	return dara.Prettify(s)
}

func (s IllustrationResult) GoString() string {
	return s.String()
}

func (s *IllustrationResult) GetIllustration() *Illustration {
	return s.Illustration
}

func (s *IllustrationResult) GetRequestId() *string {
	return s.RequestId
}

func (s *IllustrationResult) SetIllustration(v *Illustration) *IllustrationResult {
	s.Illustration = v
	return s
}

func (s *IllustrationResult) SetRequestId(v string) *IllustrationResult {
	s.RequestId = &v
	return s
}

func (s *IllustrationResult) Validate() error {
	if s.Illustration != nil {
		if err := s.Illustration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
