// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteEventFieldRequest
	GetLang() *string
	SetEventCode(v string) *DeleteEventFieldRequest
	GetEventCode() *string
	SetFieldName(v string) *DeleteEventFieldRequest
	GetFieldName() *string
	SetRegId(v string) *DeleteEventFieldRequest
	GetRegId() *string
}

type DeleteEventFieldRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Event code
	//
	// example:
	//
	// de_awukck7117
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Field name
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteEventFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventFieldRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventFieldRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteEventFieldRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DeleteEventFieldRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *DeleteEventFieldRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteEventFieldRequest) SetLang(v string) *DeleteEventFieldRequest {
	s.Lang = &v
	return s
}

func (s *DeleteEventFieldRequest) SetEventCode(v string) *DeleteEventFieldRequest {
	s.EventCode = &v
	return s
}

func (s *DeleteEventFieldRequest) SetFieldName(v string) *DeleteEventFieldRequest {
	s.FieldName = &v
	return s
}

func (s *DeleteEventFieldRequest) SetRegId(v string) *DeleteEventFieldRequest {
	s.RegId = &v
	return s
}

func (s *DeleteEventFieldRequest) Validate() error {
	return dara.Validate(s)
}
