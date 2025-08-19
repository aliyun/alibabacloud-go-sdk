// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeadLetterQueue interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *DeadLetterQueue
	GetArn() *string
}

type DeadLetterQueue struct {
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s DeadLetterQueue) String() string {
	return dara.Prettify(s)
}

func (s DeadLetterQueue) GoString() string {
	return s.String()
}

func (s *DeadLetterQueue) GetArn() *string {
	return s.Arn
}

func (s *DeadLetterQueue) SetArn(v string) *DeadLetterQueue {
	s.Arn = &v
	return s
}

func (s *DeadLetterQueue) Validate() error {
	return dara.Validate(s)
}
