// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactWhiteListList(v []*string) *SaveContactWhiteListRequest
	GetContactWhiteListList() []*string
	SetContactWhiteListsJson(v string) *SaveContactWhiteListRequest
	GetContactWhiteListsJson() *string
	SetInstanceId(v string) *SaveContactWhiteListRequest
	GetInstanceId() *string
}

type SaveContactWhiteListRequest struct {
	// example:
	//
	// {}
	ContactWhiteListList []*string `json:"ContactWhiteListList,omitempty" xml:"ContactWhiteListList,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	ContactWhiteListsJson *string `json:"ContactWhiteListsJson,omitempty" xml:"ContactWhiteListsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c415bb6c-2e6f-46aa-afd9-3b65b6dbe2bc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SaveContactWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveContactWhiteListRequest) GoString() string {
	return s.String()
}

func (s *SaveContactWhiteListRequest) GetContactWhiteListList() []*string {
	return s.ContactWhiteListList
}

func (s *SaveContactWhiteListRequest) GetContactWhiteListsJson() *string {
	return s.ContactWhiteListsJson
}

func (s *SaveContactWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveContactWhiteListRequest) SetContactWhiteListList(v []*string) *SaveContactWhiteListRequest {
	s.ContactWhiteListList = v
	return s
}

func (s *SaveContactWhiteListRequest) SetContactWhiteListsJson(v string) *SaveContactWhiteListRequest {
	s.ContactWhiteListsJson = &v
	return s
}

func (s *SaveContactWhiteListRequest) SetInstanceId(v string) *SaveContactWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveContactWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
