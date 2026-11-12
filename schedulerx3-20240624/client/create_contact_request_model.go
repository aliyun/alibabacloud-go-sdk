// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateContactRequest
	GetConfig() *string
	SetContactName(v string) *CreateContactRequest
	GetContactName() *string
	SetType(v string) *CreateContactRequest
	GetType() *string
}

type CreateContactRequest struct {
	// The JSON character string for channel parameter settings. Example for the IM type: {"channels":[{"channelType":"dingtalk","clientId":"xxx","clientSecret":"xxx","targetType":"group","targetId":"xxx","robotCode":"xxx"}]}
	//
	// This parameter is required.
	//
	// example:
	//
	// {"channels":[{"channelType":"dingtalk","clientId":"xxx","clientSecret":"xxx","targetType":"group","targetId":"xxx","robotCode":"xxx"}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// Chen Rui
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The contact type.
	//
	// - IM: Configures an IM channel notification contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// IM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContactRequest) GoString() string {
	return s.String()
}

func (s *CreateContactRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateContactRequest) GetType() *string {
	return s.Type
}

func (s *CreateContactRequest) SetConfig(v string) *CreateContactRequest {
	s.Config = &v
	return s
}

func (s *CreateContactRequest) SetContactName(v string) *CreateContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateContactRequest) SetType(v string) *CreateContactRequest {
	s.Type = &v
	return s
}

func (s *CreateContactRequest) Validate() error {
	return dara.Validate(s)
}
