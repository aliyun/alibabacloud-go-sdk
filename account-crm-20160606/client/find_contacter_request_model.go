// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindContacterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContacterId(v int64) *FindContacterRequest
	GetContacterId() *int64
	SetLocaleString(v string) *FindContacterRequest
	GetLocaleString() *string
	SetType(v string) *FindContacterRequest
	GetType() *string
}

type FindContacterRequest struct {
	// This parameter is required.
	ContacterId  *int64  `json:"ContacterId,omitempty" xml:"ContacterId,omitempty"`
	LocaleString *string `json:"LocaleString,omitempty" xml:"LocaleString,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FindContacterRequest) String() string {
	return dara.Prettify(s)
}

func (s FindContacterRequest) GoString() string {
	return s.String()
}

func (s *FindContacterRequest) GetContacterId() *int64 {
	return s.ContacterId
}

func (s *FindContacterRequest) GetLocaleString() *string {
	return s.LocaleString
}

func (s *FindContacterRequest) GetType() *string {
	return s.Type
}

func (s *FindContacterRequest) SetContacterId(v int64) *FindContacterRequest {
	s.ContacterId = &v
	return s
}

func (s *FindContacterRequest) SetLocaleString(v string) *FindContacterRequest {
	s.LocaleString = &v
	return s
}

func (s *FindContacterRequest) SetType(v string) *FindContacterRequest {
	s.Type = &v
	return s
}

func (s *FindContacterRequest) Validate() error {
	return dara.Validate(s)
}
