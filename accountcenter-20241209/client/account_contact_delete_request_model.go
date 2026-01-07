// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AccountContactDeleteRequest
	GetAppName() *string
	SetContactId(v int64) *AccountContactDeleteRequest
	GetContactId() *int64
	SetOrientedEcId(v string) *AccountContactDeleteRequest
	GetOrientedEcId() *string
	SetOrientedLeId(v string) *AccountContactDeleteRequest
	GetOrientedLeId() *string
	SetOrientedNbId(v string) *AccountContactDeleteRequest
	GetOrientedNbId() *string
}

type AccountContactDeleteRequest struct {
	// example:
	//
	// xxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// xxx
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// null
	OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
	// example:
	//
	// null
	OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
	// example:
	//
	// null
	OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
}

func (s AccountContactDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s AccountContactDeleteRequest) GoString() string {
	return s.String()
}

func (s *AccountContactDeleteRequest) GetAppName() *string {
	return s.AppName
}

func (s *AccountContactDeleteRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *AccountContactDeleteRequest) GetOrientedEcId() *string {
	return s.OrientedEcId
}

func (s *AccountContactDeleteRequest) GetOrientedLeId() *string {
	return s.OrientedLeId
}

func (s *AccountContactDeleteRequest) GetOrientedNbId() *string {
	return s.OrientedNbId
}

func (s *AccountContactDeleteRequest) SetAppName(v string) *AccountContactDeleteRequest {
	s.AppName = &v
	return s
}

func (s *AccountContactDeleteRequest) SetContactId(v int64) *AccountContactDeleteRequest {
	s.ContactId = &v
	return s
}

func (s *AccountContactDeleteRequest) SetOrientedEcId(v string) *AccountContactDeleteRequest {
	s.OrientedEcId = &v
	return s
}

func (s *AccountContactDeleteRequest) SetOrientedLeId(v string) *AccountContactDeleteRequest {
	s.OrientedLeId = &v
	return s
}

func (s *AccountContactDeleteRequest) SetOrientedNbId(v string) *AccountContactDeleteRequest {
	s.OrientedNbId = &v
	return s
}

func (s *AccountContactDeleteRequest) Validate() error {
	return dara.Validate(s)
}
