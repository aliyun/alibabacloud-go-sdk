// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaToSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *UpdateMediaToSearchLibRequest
	GetMediaId() *string
	SetMsgBody(v string) *UpdateMediaToSearchLibRequest
	GetMsgBody() *string
	SetNamespace(v string) *UpdateMediaToSearchLibRequest
	GetNamespace() *string
	SetSearchLibName(v string) *UpdateMediaToSearchLibRequest
	GetSearchLibName() *string
}

type UpdateMediaToSearchLibRequest struct {
	// The ID of the media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****019b82e24b37a1c2958dec38****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The message body.
	//
	// This parameter is required.
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

func (s UpdateMediaToSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaToSearchLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaToSearchLibRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaToSearchLibRequest) GetMsgBody() *string {
	return s.MsgBody
}

func (s *UpdateMediaToSearchLibRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateMediaToSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *UpdateMediaToSearchLibRequest) SetMediaId(v string) *UpdateMediaToSearchLibRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaToSearchLibRequest) SetMsgBody(v string) *UpdateMediaToSearchLibRequest {
	s.MsgBody = &v
	return s
}

func (s *UpdateMediaToSearchLibRequest) SetNamespace(v string) *UpdateMediaToSearchLibRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateMediaToSearchLibRequest) SetSearchLibName(v string) *UpdateMediaToSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *UpdateMediaToSearchLibRequest) Validate() error {
	return dara.Validate(s)
}
