// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactQueryDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AccountContactQueryDetailRequest
	GetAppName() *string
	SetContactId(v int64) *AccountContactQueryDetailRequest
	GetContactId() *int64
	SetOrientedEcId(v string) *AccountContactQueryDetailRequest
	GetOrientedEcId() *string
	SetOrientedLeId(v string) *AccountContactQueryDetailRequest
	GetOrientedLeId() *string
	SetOrientedNbId(v string) *AccountContactQueryDetailRequest
	GetOrientedNbId() *string
}

type AccountContactQueryDetailRequest struct {
	// example:
	//
	// yanxuan
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

func (s AccountContactQueryDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s AccountContactQueryDetailRequest) GoString() string {
	return s.String()
}

func (s *AccountContactQueryDetailRequest) GetAppName() *string {
	return s.AppName
}

func (s *AccountContactQueryDetailRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *AccountContactQueryDetailRequest) GetOrientedEcId() *string {
	return s.OrientedEcId
}

func (s *AccountContactQueryDetailRequest) GetOrientedLeId() *string {
	return s.OrientedLeId
}

func (s *AccountContactQueryDetailRequest) GetOrientedNbId() *string {
	return s.OrientedNbId
}

func (s *AccountContactQueryDetailRequest) SetAppName(v string) *AccountContactQueryDetailRequest {
	s.AppName = &v
	return s
}

func (s *AccountContactQueryDetailRequest) SetContactId(v int64) *AccountContactQueryDetailRequest {
	s.ContactId = &v
	return s
}

func (s *AccountContactQueryDetailRequest) SetOrientedEcId(v string) *AccountContactQueryDetailRequest {
	s.OrientedEcId = &v
	return s
}

func (s *AccountContactQueryDetailRequest) SetOrientedLeId(v string) *AccountContactQueryDetailRequest {
	s.OrientedLeId = &v
	return s
}

func (s *AccountContactQueryDetailRequest) SetOrientedNbId(v string) *AccountContactQueryDetailRequest {
	s.OrientedNbId = &v
	return s
}

func (s *AccountContactQueryDetailRequest) Validate() error {
	return dara.Validate(s)
}
