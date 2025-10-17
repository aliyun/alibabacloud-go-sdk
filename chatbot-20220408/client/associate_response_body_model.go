// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociate(v []*AssociateResponseBodyAssociate) *AssociateResponseBody
	GetAssociate() []*AssociateResponseBodyAssociate
	SetMessageId(v string) *AssociateResponseBody
	GetMessageId() *string
	SetRequestId(v string) *AssociateResponseBody
	GetRequestId() *string
	SetSessionId(v string) *AssociateResponseBody
	GetSessionId() *string
}

type AssociateResponseBody struct {
	Associate []*AssociateResponseBodyAssociate `json:"Associate,omitempty" xml:"Associate,omitempty" type:"Repeated"`
	// example:
	//
	// 5ADF0EBD-7C50-1922-A28B-43215B47CC1A
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 5C20F0D4-9721-178A-8236-3BF990634962
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1531ded6b3df4afca4be63943f708bb7
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AssociateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResponseBody) GetAssociate() []*AssociateResponseBodyAssociate {
	return s.Associate
}

func (s *AssociateResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *AssociateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *AssociateResponseBody) SetAssociate(v []*AssociateResponseBodyAssociate) *AssociateResponseBody {
	s.Associate = v
	return s
}

func (s *AssociateResponseBody) SetMessageId(v string) *AssociateResponseBody {
	s.MessageId = &v
	return s
}

func (s *AssociateResponseBody) SetRequestId(v string) *AssociateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateResponseBody) SetSessionId(v string) *AssociateResponseBody {
	s.SessionId = &v
	return s
}

func (s *AssociateResponseBody) Validate() error {
	if s.Associate != nil {
		for _, item := range s.Associate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AssociateResponseBodyAssociate struct {
	// example:
	//
	// {}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 公积金提取的政策
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AssociateResponseBodyAssociate) String() string {
	return dara.Prettify(s)
}

func (s AssociateResponseBodyAssociate) GoString() string {
	return s.String()
}

func (s *AssociateResponseBodyAssociate) GetMeta() *string {
	return s.Meta
}

func (s *AssociateResponseBodyAssociate) GetTitle() *string {
	return s.Title
}

func (s *AssociateResponseBodyAssociate) SetMeta(v string) *AssociateResponseBodyAssociate {
	s.Meta = &v
	return s
}

func (s *AssociateResponseBodyAssociate) SetTitle(v string) *AssociateResponseBodyAssociate {
	s.Title = &v
	return s
}

func (s *AssociateResponseBodyAssociate) Validate() error {
	return dara.Validate(s)
}
