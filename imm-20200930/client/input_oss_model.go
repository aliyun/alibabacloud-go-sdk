// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInputOSS interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *InputOSS
	GetBucket() *string
	SetMatchExpressions(v []*string) *InputOSS
	GetMatchExpressions() []*string
	SetPrefix(v string) *InputOSS
	GetPrefix() *string
}

type InputOSS struct {
	// This parameter is required.
	Bucket           *string   `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	MatchExpressions []*string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty" type:"Repeated"`
	Prefix           *string   `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s InputOSS) String() string {
	return dara.Prettify(s)
}

func (s InputOSS) GoString() string {
	return s.String()
}

func (s *InputOSS) GetBucket() *string {
	return s.Bucket
}

func (s *InputOSS) GetMatchExpressions() []*string {
	return s.MatchExpressions
}

func (s *InputOSS) GetPrefix() *string {
	return s.Prefix
}

func (s *InputOSS) SetBucket(v string) *InputOSS {
	s.Bucket = &v
	return s
}

func (s *InputOSS) SetMatchExpressions(v []*string) *InputOSS {
	s.MatchExpressions = v
	return s
}

func (s *InputOSS) SetPrefix(v string) *InputOSS {
	s.Prefix = &v
	return s
}

func (s *InputOSS) Validate() error {
	return dara.Validate(s)
}
