// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindAllContacterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *FindAllContacterRequest
	GetAppName() *string
	SetLocaleString(v string) *FindAllContacterRequest
	GetLocaleString() *string
	SetType(v string) *FindAllContacterRequest
	GetType() *string
	SetUserId(v int64) *FindAllContacterRequest
	GetUserId() *int64
}

type FindAllContacterRequest struct {
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	LocaleString *string `json:"LocaleString,omitempty" xml:"LocaleString,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s FindAllContacterRequest) String() string {
	return dara.Prettify(s)
}

func (s FindAllContacterRequest) GoString() string {
	return s.String()
}

func (s *FindAllContacterRequest) GetAppName() *string {
	return s.AppName
}

func (s *FindAllContacterRequest) GetLocaleString() *string {
	return s.LocaleString
}

func (s *FindAllContacterRequest) GetType() *string {
	return s.Type
}

func (s *FindAllContacterRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *FindAllContacterRequest) SetAppName(v string) *FindAllContacterRequest {
	s.AppName = &v
	return s
}

func (s *FindAllContacterRequest) SetLocaleString(v string) *FindAllContacterRequest {
	s.LocaleString = &v
	return s
}

func (s *FindAllContacterRequest) SetType(v string) *FindAllContacterRequest {
	s.Type = &v
	return s
}

func (s *FindAllContacterRequest) SetUserId(v int64) *FindAllContacterRequest {
	s.UserId = &v
	return s
}

func (s *FindAllContacterRequest) Validate() error {
	return dara.Validate(s)
}
