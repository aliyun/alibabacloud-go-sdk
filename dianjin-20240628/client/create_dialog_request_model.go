// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *CreateDialogRequest
	GetChannel() *string
	SetEnableLibrary(v bool) *CreateDialogRequest
	GetEnableLibrary() *bool
	SetMetaData(v map[string]interface{}) *CreateDialogRequest
	GetMetaData() map[string]interface{}
	SetPlayCode(v string) *CreateDialogRequest
	GetPlayCode() *string
	SetQaLibraryList(v []*string) *CreateDialogRequest
	GetQaLibraryList() []*string
	SetRequestId(v string) *CreateDialogRequest
	GetRequestId() *string
	SetSelfDirected(v bool) *CreateDialogRequest
	GetSelfDirected() *bool
}

type CreateDialogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// taobao
	Channel       *string `json:"channel,omitempty" xml:"channel,omitempty"`
	EnableLibrary *bool   `json:"enableLibrary,omitempty" xml:"enableLibrary,omitempty"`
	// example:
	//
	// null
	MetaData map[string]interface{} `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// live_broadcast_qa
	PlayCode      *string   `json:"playCode,omitempty" xml:"playCode,omitempty"`
	QaLibraryList []*string `json:"qaLibraryList,omitempty" xml:"qaLibraryList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ebf83826-dc1c-46f8-9759-0fb6da4c8xxx
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SelfDirected *bool   `json:"selfDirected,omitempty" xml:"selfDirected,omitempty"`
}

func (s CreateDialogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogRequest) GoString() string {
	return s.String()
}

func (s *CreateDialogRequest) GetChannel() *string {
	return s.Channel
}

func (s *CreateDialogRequest) GetEnableLibrary() *bool {
	return s.EnableLibrary
}

func (s *CreateDialogRequest) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *CreateDialogRequest) GetPlayCode() *string {
	return s.PlayCode
}

func (s *CreateDialogRequest) GetQaLibraryList() []*string {
	return s.QaLibraryList
}

func (s *CreateDialogRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDialogRequest) GetSelfDirected() *bool {
	return s.SelfDirected
}

func (s *CreateDialogRequest) SetChannel(v string) *CreateDialogRequest {
	s.Channel = &v
	return s
}

func (s *CreateDialogRequest) SetEnableLibrary(v bool) *CreateDialogRequest {
	s.EnableLibrary = &v
	return s
}

func (s *CreateDialogRequest) SetMetaData(v map[string]interface{}) *CreateDialogRequest {
	s.MetaData = v
	return s
}

func (s *CreateDialogRequest) SetPlayCode(v string) *CreateDialogRequest {
	s.PlayCode = &v
	return s
}

func (s *CreateDialogRequest) SetQaLibraryList(v []*string) *CreateDialogRequest {
	s.QaLibraryList = v
	return s
}

func (s *CreateDialogRequest) SetRequestId(v string) *CreateDialogRequest {
	s.RequestId = &v
	return s
}

func (s *CreateDialogRequest) SetSelfDirected(v bool) *CreateDialogRequest {
	s.SelfDirected = &v
	return s
}

func (s *CreateDialogRequest) Validate() error {
	return dara.Validate(s)
}
