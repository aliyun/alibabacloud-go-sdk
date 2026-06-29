// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkforce interface {
	dara.Model
	String() string
	GoString() string
	SetNodeType(v string) *Workforce
	GetNodeType() *string
	SetUsers(v []*SimpleUser) *Workforce
	GetUsers() []*SimpleUser
	SetWorkNodeId(v int32) *Workforce
	GetWorkNodeId() *int32
}

type Workforce struct {
	// Node name.
	//
	// Valid values include: SAMPLING, CHECK, MARK.
	//
	// example:
	//
	// CHECK
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// List of user information
	Users []*SimpleUser `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	// Node ID
	//
	// example:
	//
	// 2
	WorkNodeId *int32 `json:"WorkNodeId,omitempty" xml:"WorkNodeId,omitempty"`
}

func (s Workforce) String() string {
	return dara.Prettify(s)
}

func (s Workforce) GoString() string {
	return s.String()
}

func (s *Workforce) GetNodeType() *string {
	return s.NodeType
}

func (s *Workforce) GetUsers() []*SimpleUser {
	return s.Users
}

func (s *Workforce) GetWorkNodeId() *int32 {
	return s.WorkNodeId
}

func (s *Workforce) SetNodeType(v string) *Workforce {
	s.NodeType = &v
	return s
}

func (s *Workforce) SetUsers(v []*SimpleUser) *Workforce {
	s.Users = v
	return s
}

func (s *Workforce) SetWorkNodeId(v int32) *Workforce {
	s.WorkNodeId = &v
	return s
}

func (s *Workforce) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
