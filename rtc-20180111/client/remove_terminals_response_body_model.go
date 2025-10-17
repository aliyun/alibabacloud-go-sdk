// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTerminalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTerminalsResponseBody
	GetRequestId() *string
	SetTerminals(v *RemoveTerminalsResponseBodyTerminals) *RemoveTerminalsResponseBody
	GetTerminals() *RemoveTerminalsResponseBodyTerminals
}

type RemoveTerminalsResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Terminals *RemoveTerminalsResponseBodyTerminals `json:"Terminals,omitempty" xml:"Terminals,omitempty" type:"Struct"`
}

func (s RemoveTerminalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTerminalsResponseBody) GetTerminals() *RemoveTerminalsResponseBodyTerminals {
	return s.Terminals
}

func (s *RemoveTerminalsResponseBody) SetRequestId(v string) *RemoveTerminalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTerminalsResponseBody) SetTerminals(v *RemoveTerminalsResponseBodyTerminals) *RemoveTerminalsResponseBody {
	s.Terminals = v
	return s
}

func (s *RemoveTerminalsResponseBody) Validate() error {
	if s.Terminals != nil {
		if err := s.Terminals.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveTerminalsResponseBodyTerminals struct {
	Terminal []*RemoveTerminalsResponseBodyTerminalsTerminal `json:"Terminal,omitempty" xml:"Terminal,omitempty" type:"Repeated"`
}

func (s RemoveTerminalsResponseBodyTerminals) String() string {
	return dara.Prettify(s)
}

func (s RemoveTerminalsResponseBodyTerminals) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponseBodyTerminals) GetTerminal() []*RemoveTerminalsResponseBodyTerminalsTerminal {
	return s.Terminal
}

func (s *RemoveTerminalsResponseBodyTerminals) SetTerminal(v []*RemoveTerminalsResponseBodyTerminalsTerminal) *RemoveTerminalsResponseBodyTerminals {
	s.Terminal = v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminals) Validate() error {
	if s.Terminal != nil {
		for _, item := range s.Terminal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RemoveTerminalsResponseBodyTerminalsTerminal struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1811****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveTerminalsResponseBodyTerminalsTerminal) String() string {
	return dara.Prettify(s)
}

func (s RemoveTerminalsResponseBodyTerminalsTerminal) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) GetCode() *int32 {
	return s.Code
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) GetId() *string {
	return s.Id
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) GetMessage() *string {
	return s.Message
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) SetCode(v int32) *RemoveTerminalsResponseBodyTerminalsTerminal {
	s.Code = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) SetId(v string) *RemoveTerminalsResponseBodyTerminalsTerminal {
	s.Id = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) SetMessage(v string) *RemoveTerminalsResponseBodyTerminalsTerminal {
	s.Message = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) Validate() error {
	return dara.Validate(s)
}
