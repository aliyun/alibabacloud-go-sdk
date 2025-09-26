// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCodeInterpreterOut interface {
	dara.Model
	String() string
	GoString() string
	SetCodeInterpreterId(v string) *DeleteCodeInterpreterOut
	GetCodeInterpreterId() *string
	SetCodeInterpreterName(v string) *DeleteCodeInterpreterOut
	GetCodeInterpreterName() *string
	SetStatus(v string) *DeleteCodeInterpreterOut
	GetStatus() *string
}

type DeleteCodeInterpreterOut struct {
	CodeInterpreterId   *string `json:"codeInterpreterId,omitempty" xml:"codeInterpreterId,omitempty"`
	CodeInterpreterName *string `json:"codeInterpreterName,omitempty" xml:"codeInterpreterName,omitempty"`
	Status              *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteCodeInterpreterOut) String() string {
	return dara.Prettify(s)
}

func (s DeleteCodeInterpreterOut) GoString() string {
	return s.String()
}

func (s *DeleteCodeInterpreterOut) GetCodeInterpreterId() *string {
	return s.CodeInterpreterId
}

func (s *DeleteCodeInterpreterOut) GetCodeInterpreterName() *string {
	return s.CodeInterpreterName
}

func (s *DeleteCodeInterpreterOut) GetStatus() *string {
	return s.Status
}

func (s *DeleteCodeInterpreterOut) SetCodeInterpreterId(v string) *DeleteCodeInterpreterOut {
	s.CodeInterpreterId = &v
	return s
}

func (s *DeleteCodeInterpreterOut) SetCodeInterpreterName(v string) *DeleteCodeInterpreterOut {
	s.CodeInterpreterName = &v
	return s
}

func (s *DeleteCodeInterpreterOut) SetStatus(v string) *DeleteCodeInterpreterOut {
	s.Status = &v
	return s
}

func (s *DeleteCodeInterpreterOut) Validate() error {
	return dara.Validate(s)
}
