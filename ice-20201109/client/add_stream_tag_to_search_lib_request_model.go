// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStreamTagToSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *AddStreamTagToSearchLibRequest
	GetMediaId() *string
	SetMsgBody(v string) *AddStreamTagToSearchLibRequest
	GetMsgBody() *string
	SetNamespace(v string) *AddStreamTagToSearchLibRequest
	GetNamespace() *string
	SetSearchLibName(v string) *AddStreamTagToSearchLibRequest
	GetSearchLibName() *string
}

type AddStreamTagToSearchLibRequest struct {
	// example:
	//
	// ******b48fb04483915d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// {"startTime":1657684600793,"endTime":1657684600793,"userData":"{}"}
	MsgBody *string `json:"MsgBody,omitempty" xml:"MsgBody,omitempty"`
	// example:
	//
	// name-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// Stream_xxx
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s AddStreamTagToSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s AddStreamTagToSearchLibRequest) GoString() string {
	return s.String()
}

func (s *AddStreamTagToSearchLibRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *AddStreamTagToSearchLibRequest) GetMsgBody() *string {
	return s.MsgBody
}

func (s *AddStreamTagToSearchLibRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *AddStreamTagToSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *AddStreamTagToSearchLibRequest) SetMediaId(v string) *AddStreamTagToSearchLibRequest {
	s.MediaId = &v
	return s
}

func (s *AddStreamTagToSearchLibRequest) SetMsgBody(v string) *AddStreamTagToSearchLibRequest {
	s.MsgBody = &v
	return s
}

func (s *AddStreamTagToSearchLibRequest) SetNamespace(v string) *AddStreamTagToSearchLibRequest {
	s.Namespace = &v
	return s
}

func (s *AddStreamTagToSearchLibRequest) SetSearchLibName(v string) *AddStreamTagToSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *AddStreamTagToSearchLibRequest) Validate() error {
	return dara.Validate(s)
}
