// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaFromSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *DeleteMediaFromSearchLibRequest
	GetMediaId() *string
	SetMsgBody(v string) *DeleteMediaFromSearchLibRequest
	GetMsgBody() *string
	SetNamespace(v string) *DeleteMediaFromSearchLibRequest
	GetNamespace() *string
	SetSearchLibName(v string) *DeleteMediaFromSearchLibRequest
	GetSearchLibName() *string
}

type DeleteMediaFromSearchLibRequest struct {
	// The ID of the media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The message body.
	//
	// example:
	//
	// {}
	MsgBody   *string `json:"MsgBody,omitempty" xml:"MsgBody,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the search library. Default value: ims-default-search-lib.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s DeleteMediaFromSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaFromSearchLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaFromSearchLibRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *DeleteMediaFromSearchLibRequest) GetMsgBody() *string {
	return s.MsgBody
}

func (s *DeleteMediaFromSearchLibRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteMediaFromSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *DeleteMediaFromSearchLibRequest) SetMediaId(v string) *DeleteMediaFromSearchLibRequest {
	s.MediaId = &v
	return s
}

func (s *DeleteMediaFromSearchLibRequest) SetMsgBody(v string) *DeleteMediaFromSearchLibRequest {
	s.MsgBody = &v
	return s
}

func (s *DeleteMediaFromSearchLibRequest) SetNamespace(v string) *DeleteMediaFromSearchLibRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteMediaFromSearchLibRequest) SetSearchLibName(v string) *DeleteMediaFromSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *DeleteMediaFromSearchLibRequest) Validate() error {
	return dara.Validate(s)
}
