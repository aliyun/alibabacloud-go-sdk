// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunctionLayer interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *FunctionLayer
	GetArn() *string
	SetSize(v int64) *FunctionLayer
	GetSize() *int64
}

type FunctionLayer struct {
	// example:
	//
	// acs:fc:cn-beijing:186824xxxxxx:layers/fc_layer/versions/1
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
	// example:
	//
	// 421
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s FunctionLayer) String() string {
	return dara.Prettify(s)
}

func (s FunctionLayer) GoString() string {
	return s.String()
}

func (s *FunctionLayer) GetArn() *string {
	return s.Arn
}

func (s *FunctionLayer) GetSize() *int64 {
	return s.Size
}

func (s *FunctionLayer) SetArn(v string) *FunctionLayer {
	s.Arn = &v
	return s
}

func (s *FunctionLayer) SetSize(v int64) *FunctionLayer {
	s.Size = &v
	return s
}

func (s *FunctionLayer) Validate() error {
	return dara.Validate(s)
}
