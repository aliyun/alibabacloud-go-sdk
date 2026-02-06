// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactBlockListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactBlockListList(v []*string) *SaveContactBlockListRequest
	GetContactBlockListList() []*string
	SetContactBlockListsJson(v string) *SaveContactBlockListRequest
	GetContactBlockListsJson() *string
	SetInstanceId(v string) *SaveContactBlockListRequest
	GetInstanceId() *string
}

type SaveContactBlockListRequest struct {
	// example:
	//
	// []
	ContactBlockListList []*string `json:"ContactBlockListList,omitempty" xml:"ContactBlockListList,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	ContactBlockListsJson *string `json:"ContactBlockListsJson,omitempty" xml:"ContactBlockListsJson,omitempty"`
	// example:
	//
	// c3c92de8-e4bd-4db4-a962-50f8acce40bc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SaveContactBlockListRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveContactBlockListRequest) GoString() string {
	return s.String()
}

func (s *SaveContactBlockListRequest) GetContactBlockListList() []*string {
	return s.ContactBlockListList
}

func (s *SaveContactBlockListRequest) GetContactBlockListsJson() *string {
	return s.ContactBlockListsJson
}

func (s *SaveContactBlockListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveContactBlockListRequest) SetContactBlockListList(v []*string) *SaveContactBlockListRequest {
	s.ContactBlockListList = v
	return s
}

func (s *SaveContactBlockListRequest) SetContactBlockListsJson(v string) *SaveContactBlockListRequest {
	s.ContactBlockListsJson = &v
	return s
}

func (s *SaveContactBlockListRequest) SetInstanceId(v string) *SaveContactBlockListRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveContactBlockListRequest) Validate() error {
	return dara.Validate(s)
}
